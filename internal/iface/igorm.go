package iface

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
)

type (
	IGorm interface {
		Create(value interface{}) (tx *gorm.DB)
		First(dest interface{}, conditions ...interface{}) (tx *gorm.DB)
		Find(dest interface{}, conditions ...interface{}) (tx *gorm.DB)
		Save(value interface{}) (tx *gorm.DB)
		Updates(value interface{}) (tx *gorm.DB)
		Delete(value interface{}, conditions ...interface{}) (tx *gorm.DB)
		Limit(value int) (tx *gorm.DB)
		Offset(value int) (tx *gorm.DB)
		Select(query interface{}, args ...interface{}) (tx *gorm.DB)
		Where(query interface{}, args ...interface{}) (tx *gorm.DB)
		Joins(query string, args ...interface{}) (tx *gorm.DB)
		Model(value interface{}) (tx *gorm.DB)
		Order(value interface{}) (tx *gorm.DB)
		WithContext(ctx context.Context) (tx *gorm.DB)
		Begin(opts ...*sql.TxOptions) (tx *gorm.DB)
	}
)
