package mgo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"

	"github.com/shawnwy/go-utils/v5/errors"
	"github.com/shawnwy/go-utils/v5/respond"
)

type iModels interface {
	Init(dbName string, db *mongo.Database)
}

type Option func(*options.ClientOptions)

func WithPoolSize(min, max uint64) Option {
	return func(opt *options.ClientOptions) {
		opt.SetMinPoolSize(min)
		if max > 0 {
			opt.SetMaxPoolSize(max)
		}
	}
}

func WithConnectTimeout(d time.Duration) Option {
	return func(opt *options.ClientOptions) {
		opt.SetConnectTimeout(d)
	}
}

func WithSocketTimeout(d time.Duration) Option {
	return func(opt *options.ClientOptions) {
		opt.SetSocketTimeout(d)
	}
}

func WithMaxConnIdleTime(d time.Duration) Option {
	return func(opt *options.ClientOptions) {
		opt.SetMaxConnIdleTime(d)
	}
}

func Connect(uri, dbname string, opts ...Option) *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cliOpts := options.Client().ApplyURI(uri)
	for _, o := range opts {
		o(cliOpts)
	}
	client, err := mongo.Connect(ctx, cliOpts)
	if err != nil {
		err = errors.With(respond.ErrMongoConnect, err)
		zap.L().Panic("failed to connect mgo",
			zap.String("uri", uri),
			zap.String("dbname", dbname),
			zap.Error(err))
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		err = errors.With(respond.ErrMongoConnect, err)
		zap.L().Panic("failed to ping mgo",
			zap.String("uri", uri),
			zap.String("dbname", dbname),
			zap.Error(err))
	}
	return client.Database(dbname)
}

func New(uri, dbname string, enableSharding bool, ms ...iModels) *mongo.Database {
	db := Connect(uri, dbname)
	if enableSharding {
		result := db.RunCommand(context.Background(), bson.D{{"enableSharding", dbname}})
		var document bson.M
		if err := result.Decode(&document); err != nil {
			zap.L().Warn("failed to enable sharding", zap.Error(err))
		}
	}
	for _, m := range ms {
		m.Init(dbname, db)
	}
	return db
}

func Close(c *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// client.Disconnect method also has deadline.
	// returns error if any,
	if err := c.Disconnect(ctx); err != nil {
		zap.L().Panic("Failed to disconnect mgo!", zap.Error(err))
	}
}

type autoIDxCounter struct {
	Collection string `bson:"collection"`
	ID         int64  `bson:"id"`
}

func AutoID(db *mongo.Database, collection string) int64 {
	return incrID(db, collection, 1)
}

func BatchAutoID(db *mongo.Database, collection string, qty int) int64 {
	return incrID(db, collection, qty)
}

func incrID(db *mongo.Database, collection string, qty int) int64 {
	filter := bson.D{{"collection", "auto_ids"}}
	update := bson.D{
		{"$inc", bson.D{
			{"id", qty},
		}},
	}
	opt := options.FindOneAndUpdate().
		SetReturnDocument(options.After).
		SetUpsert(true)

	autoID := new(autoIDxCounter)
	if err := db.Collection(fmt.Sprintf("autoids_%s", collection)).
		FindOneAndUpdate(context.TODO(), filter, update, opt).
		Decode(&autoID); err != nil {
		zap.L().Warn("AutoID Decoding Failed")
	}
	return autoID.ID
}
