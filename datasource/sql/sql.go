package sql

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQL(uri string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(uri), &gorm.Config{})
	if err != nil {
		zap.L().Panic("RDB> Failed to open DB", zap.Error(err))
	}
	return db
}
