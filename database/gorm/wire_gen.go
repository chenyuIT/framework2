// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package gorm

import (
	"context"
	"github.com/chenyuIT/framework/contracts/config"
	"github.com/chenyuIT/framework/database/db"
)

// Injectors from wire.go:

//go:generate wire
func InitializeGorm(config2 config.Config, connection string) *GormImpl {
	configImpl := db.NewConfigImpl(config2, connection)
	dialectorImpl := NewDialectorImpl(config2, connection)
	gormImpl := NewGormImpl(config2, connection, configImpl, dialectorImpl)
	return gormImpl
}

//go:generate wire
func InitializeQuery(ctx context.Context, config2 config.Config, connection string) (*QueryImpl, error) {
	configImpl := db.NewConfigImpl(config2, connection)
	dialectorImpl := NewDialectorImpl(config2, connection)
	gormImpl := NewGormImpl(config2, connection, configImpl, dialectorImpl)
	queryImpl, err := NewQueryImpl(ctx, config2, gormImpl)
	if err != nil {
		return nil, err
	}
	return queryImpl, nil
}
