package inits

import (
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/bagasunix/bank-ina/pkg/errors"
	"github.com/bagasunix/bank-ina/server/domains/data/models"
)

func GetTables() (tables []interface{}) {
	tables = append(tables, models.NewUserBuilder().Build())
	tables = append(tables, models.NewTaskBuilder().Build())

	return tables
}

func Migrate(logger *zap.Logger, db *gorm.DB) {
	errors.HandlerWithOSExit(logger, db.AutoMigrate(GetTables()...), "AutoMigrate")
}
