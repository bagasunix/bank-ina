package inits

import (
	"context"
	"strconv"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/bagasunix/bank-ina/pkg/db"
	"github.com/bagasunix/bank-ina/pkg/envs"
	"github.com/bagasunix/bank-ina/pkg/errors"
)

func InitDb(ctx context.Context, logger *zap.Logger, configs *envs.Configs) *gorm.DB {
	port, _ := strconv.Atoi(configs.DbPort)

	configBuilder := db.NewDbConfigBuilder()
	configBuilder.SetHost(configs.DbHost)
	configBuilder.SetPort(int64(port))
	configBuilder.SetUser(configs.DbUsername)
	configBuilder.SetDatabaseName(configs.DbName)
	configBuilder.SetPassword(configs.DbPassword)

	return NewDB(ctx, logger, configBuilder.Build())
}

func NewDB(ctx context.Context, logger *zap.Logger, dbConfig *db.DbConfig) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dbConfig.GetDSN()), &gorm.Config{SkipDefaultTransaction: true, PrepareStmt: true})
	errors.HandlerWithOSExit(logger, err, "init", "database", "config", dbConfig.GetDSN())
	return db
}
