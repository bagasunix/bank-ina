package repositories

import (
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/bagasunix/bank-ina/server/domains/data/repositories/task"
	"github.com/bagasunix/bank-ina/server/domains/data/repositories/user"
)

type Repositories interface {
	GetTask() task.Repository
	GetUser() user.Repository
}

type repo struct {
	task task.Repository
	user user.Repository
}

// GetTask implements Repositories.
func (r *repo) GetTask() task.Repository {
	return r.task
}

// GetUser implements Repositories.
func (r *repo) GetUser() user.Repository {
	return r.user
}

func New(logger *zap.Logger, db *gorm.DB) Repositories {
	rp := new(repo)
	rp.user = user.NewGorm(logger, db)
	rp.task = task.NewGorm(logger, db)
	return rp
}
