package user

import (
	"context"

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

// GetByEmail implements Repository.
func (g *gormProvider) GetByEmail(ctx context.Context, email string) (result models.SliceResult[*models.User]) {
	result.Error = errors.ErrRecordNotFound(g.logger, g.GetModelName(), g.db.WithContext(ctx).Where("email = ?", email).First(&result.Value).Error)
	return result
}

// CreateUser implements Repository.
func (g *gormProvider) CreateUser(ctx *gin.Context, m *models.User) error {
	return errors.ErrDuplicateValue(g.logger, g.GetModelName(), g.db.WithContext(ctx).Create(&m).Error)
}

// DeleteUser implements Repository.
func (g *gormProvider) DeleteUser(ctx *gin.Context, id string) error {
	return errors.ErrSomethingWrong(g.logger, g.db.WithContext(ctx).Delete(models.NewUserBuilder().Build(), "id = ?", id).Error)
}

// GetAllUser implements Repository.
func (g *gormProvider) GetAllUser(ctx *gin.Context) (result models.SliceResult[models.User]) {
	result.Error = g.db.WithContext(ctx).Model(models.User{}).Find(&result.Value).Error
	return result
}

// GetByID implements Repository.
func (g *gormProvider) GetByID(ctx *gin.Context, id string) (result models.SingleResult[models.User]) {
	result.Error = errors.ErrRecordNotFound(g.logger, g.GetModelName(), g.db.WithContext(ctx).Where("id = ?", id).First(&result.Value).Error)
	return result
}

// GetConnection implements Repository.
func (g *gormProvider) GetConnection() (T any) {
	return g.db
}

// GetModelName implements Repository.
func (g *gormProvider) GetModelName() string {
	return "user"
}

// UpdateUser implements Repository.
func (g *gormProvider) UpdateUser(ctx *gin.Context, id string, m *models.User) error {
	return errors.ErrSomethingWrong(g.logger, g.db.WithContext(ctx).Where("id = ?", id).Updates(&m).Error)
}

func NewGorm(logger *zap.Logger, db *gorm.DB) Repository {
	g := new(gormProvider)
	g.db = db
	g.logger = logger
	return g
}
