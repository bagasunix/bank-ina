package task

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/bagasunix/bank-ina/pkg/errors"
	"github.com/bagasunix/bank-ina/server/domains/data/models"
)

type gormProvider struct {
	db     *gorm.DB
	logger *zap.Logger
}

// GetByName implements Repository.
func (g *gormProvider) GetByName(ctx *gin.Context, title string) (result models.SliceResult[models.Task]) {
	result.Error = errors.ErrRecordNotFound(g.logger, g.GetModelName(), g.db.WithContext(ctx).Where("title = ?", title).First(&result.Value).Error)
	return result
}

// CreateTask implements Repository.
func (g *gormProvider) CreateTask(ctx *gin.Context, m *models.Task) error {
	return errors.ErrDuplicateValue(g.logger, g.GetModelName(), g.db.WithContext(ctx).Create(&m).Error)
}

// DeleteTask implements Repository.
func (g *gormProvider) DeleteTask(ctx *gin.Context, id string) error {
	return errors.ErrSomethingWrong(g.logger, g.db.WithContext(ctx).Delete(models.NewTaskBuilder().Build(), "id = ?", id).Error)
}

// GetAllTask implements Repository.
func (g *gormProvider) GetAllTask(ctx *gin.Context) (result models.SliceResult[models.Task]) {
	result.Error = g.db.WithContext(ctx).Model(models.Task{}).Find(&result.Value).Error
	return result
}

// GetByID implements Repository.
func (g *gormProvider) GetByID(ctx *gin.Context, id string) (result models.SingleResult[models.Task]) {
	result.Error = errors.ErrRecordNotFound(g.logger, g.GetModelName(), g.db.WithContext(ctx).Where("id = ?", id).First(&result.Value).Error)
	return result
}

// GetConnection implements Repository.
func (g *gormProvider) GetConnection() (T any) {
	return g.db
}

// GetModelName implements Repository.
func (g *gormProvider) GetModelName() string {
	return "task"
}

// UpdateTask implements Repository.
func (g *gormProvider) UpdateTask(ctx *gin.Context, id string, m *models.Task) error {
	return errors.ErrSomethingWrong(g.logger, g.db.WithContext(ctx).Where("id = ?", id).Updates(&m).Error)
}

func NewGorm(logger *zap.Logger, db *gorm.DB) Repository {
	g := new(gormProvider)
	g.db = db
	g.logger = logger
	return g
}
