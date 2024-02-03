package endpoints

import (
	"github.com/gin-gonic/gin"

	"github.com/bagasunix/bank-ina/server/domains"
	"github.com/bagasunix/bank-ina/server/endpoints/requests"
)

const (
	CREATE_TASK = "CreateTask"
	UPDATE_TASK = "UpdateTask"
	DELETE_TASK = "DeleteTask"
	View_TASK   = "ViewTask"
	LIST_TASK   = "ListTask"
)

type TaskEndpoint struct {
	CreateTaskEndpoint Endpoint
	ListTaskEndpoint   Endpoint
	UpdateTaskEnpoint  Endpoint
	ViewTaskEndpoint   Endpoint
	DeleteTaskEnpoint  Endpoint
}

func NewTaskEndpoint(s domains.Service, mdw map[string][]Middleware) TaskEndpoint {
	eps := TaskEndpoint{}
	eps.CreateTaskEndpoint = makeCreateTaskEndpoint(s)
	eps.ListTaskEndpoint = makeListTaskEndpoint(s)
	eps.UpdateTaskEnpoint = makeUpdateTaskEndpoint(s)
	eps.ViewTaskEndpoint = makeViewTaskEndpoint(s)
	eps.DeleteTaskEnpoint = makeDeleteTaskEndpoint(s)

	SetEndpoint(CREATE_TASK, &eps.CreateTaskEndpoint, mdw)
	SetEndpoint(LIST_TASK, &eps.ListTaskEndpoint, mdw)
	SetEndpoint(UPDATE_TASK, &eps.UpdateTaskEnpoint, mdw)
	SetEndpoint(View_TASK, &eps.ViewTaskEndpoint, mdw)
	SetEndpoint(DELETE_TASK, &eps.DeleteTaskEnpoint, mdw)

	return eps
}

func makeCreateTaskEndpoint(s domains.Service) Endpoint {
	return func(ctx *gin.Context, request any) (response any, err error) {
		return s.CreateTask(ctx, request.(*requests.CreateTask))
	}
}

func makeUpdateTaskEndpoint(s domains.Service) Endpoint {
	return func(ctx *gin.Context, request any) (response any, err error) {
		return s.UpdateTask(ctx, request.(*requests.UpdateTask))
	}
}

func makeDeleteTaskEndpoint(s domains.Service) Endpoint {
	return func(ctx *gin.Context, request any) (response any, err error) {
		return s.DeleteTask(ctx, request.(*requests.EntityId))
	}
}

func makeListTaskEndpoint(s domains.Service) Endpoint {
	return func(ctx *gin.Context, request any) (response any, err error) {
		return s.GetAllTask(ctx)
	}
}
func makeViewTaskEndpoint(s domains.Service) Endpoint {
	return func(ctx *gin.Context, request any) (response any, err error) {
		return s.ViewTask(ctx, request.(*requests.EntityId))
	}
}
