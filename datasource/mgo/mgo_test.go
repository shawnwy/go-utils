package mgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestOption(t *testing.T) {
	opt := options.Client().ApplyURI("uri")
	var min uint64 = 0
	var max uint64 = 0
	o := WithPoolSize(min, max)
	o(opt)
	assert.Equal(t, *opt.MinPoolSize, min)
	assert.Equal(t, opt.MaxPoolSize, max)
}
