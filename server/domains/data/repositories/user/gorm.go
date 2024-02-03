package user

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/bagasunix/bank-ina/server/domains/data/models"
	"github.com/bagasunix/bank-ina/server/domains/data/repositories/base"
)

type Command interface {
	CreateUser(ctx *gin.Context, m *models.User) error
	UpdateUser(ctx *gin.Context, id string, m *models.User) error
	DeleteUser(ctx *gin.Context, id string) error
}

type Query interface {
	GetByID(ctx *gin.Context, id string) (result models.SingleResult[models.User])
	GetAllUser(ctx *gin.Context) (result models.SliceResult[models.User])
	GetByEmail(ctx context.Context, email string) (result models.SliceResult[*models.User])
}

type Repository interface {
	Command
	Query
	base.Repository
}
