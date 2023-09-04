package sql

import (
	"database/sql"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Option func(*sql.DB)

func WithConnMaxIdleTime(d time.Duration) Option {
	return func(db *sql.DB) {
		db.SetConnMaxIdleTime(d)
	}
}

func WithConnMaxLifetime(d time.Duration) Option {
	return func(db *sql.DB) {
		db.SetConnMaxLifetime(d)
	}
}

func WithMaxIdleConns(n int) Option {
	return func(db *sql.DB) {
		db.SetMaxIdleConns(n)
	}
}

func WithMaxOpenConns(n int) Option {
	return func(db *sql.DB) {
		db.SetMaxOpenConns(n)
	}
}

type GormOption func(*gorm.Config)

func WithPrepareStmt() GormOption {
	return func(config *gorm.Config) {
		config.PrepareStmt = true
	}
}

func WithDryRun() GormOption {
	return func(config *gorm.Config) {
		config.DryRun = true
	}
}

func WithLogger(logger logger.Interface) GormOption {
	return func(config *gorm.Config) {
		config.Logger = logger
	}
}

func WithNamingStrategy(namer schema.Namer) GormOption {
	return func(config *gorm.Config) {
		config.NamingStrategy = namer
	}
}

func NewGormCFG(opts ...GormOption) (cfg *gorm.Config) {
	cfg = &gorm.Config{}
	for _, o := range opts {
		o(cfg)
	}
	return
}

func NewMySQL(uri string, config *gorm.Config, opts ...Option) *gorm.DB {
	db, err := gorm.Open(mysql.Open(uri), config)
	if err != nil {
		zap.L().Panic("failed to open db", zap.Error(err))
	}
	sqlDB, err := db.DB()
	if err != nil {
		zap.L().Panic("failed to fetch sql db", zap.Error(err))
	}
	for _, o := range opts {
		o(sqlDB)
	}
	return db
}
