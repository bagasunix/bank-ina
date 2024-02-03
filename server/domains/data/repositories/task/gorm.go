package task

import (
	"github.com/gin-gonic/gin"

	"github.com/bagasunix/bank-ina/server/domains/data/models"
	"github.com/bagasunix/bank-ina/server/domains/data/repositories/base"
)

type Command interface {
	CreateTask(ctx *gin.Context, m *models.Task) error
	UpdateTask(ctx *gin.Context, id string, m *models.Task) error
	DeleteTask(ctx *gin.Context, id string) error
}

type Query interface {
	GetByID(ctx *gin.Context, id string) (result models.SingleResult[models.Task])
	GetAllTask(ctx *gin.Context) (result models.SliceResult[models.Task])
	GetByName(ctx *gin.Context, title string) (result models.SliceResult[models.Task])
}

type Repository interface {
	Command
	Query
	base.Repository
}
