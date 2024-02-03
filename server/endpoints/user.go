package endpoints

import (
	"github.com/gin-gonic/gin"

	"github.com/bagasunix/bank-ina/server/domains"
	"github.com/bagasunix/bank-ina/server/endpoints/requests"
)

const (
	CREATE_USER = "CreateUser"
	UPDATE_USER = "UpdateUser"
	DELETE_USER = "DeleteUser"
	View_USER   = "ViewUser"
	LIST_USER   = "ListUser"
)

type UserEndpoint struct {
	CreateUserEndpoint Endpoint
	ListUserEndpoint   Endpoint
	UpdateMovieEnpoint Endpoint
	ViewUserEndpoint   Endpoint
	DeleteMovieEnpoint Endpoint
}

func NewUserEndpoint(s domains.Service, mdw map[string][]Middleware) UserEndpoint {
	eps := UserEndpoint{}
	eps.CreateUserEndpoint = makeCreateUserEndpoint(s)
	eps.ListUserEndpoint = makeListUserEndpoint(s)
	eps.UpdateMovieEnpoint = makeUpdateUserEndpoint(s)
	eps.ViewUserEndpoint = makeViewUserEndpoint(s)
	eps.DeleteMovieEnpoint = makeDeleteUserEndpoint(s)

	SetEndpoint(CREATE_USER, &eps.CreateUserEndpoint, mdw)
	SetEndpoint(LIST_USER, &eps.ListUserEndpoint, mdw)
	SetEndpoint(UPDATE_USER, &eps.UpdateMovieEnpoint, mdw)
	SetEndpoint(View_USER, &eps.ViewUserEndpoint, mdw)
	SetEndpoint(DELETE_USER, &eps.DeleteMovieEnpoint, mdw)

	return eps
}

func makeCreateUserEndpoint(s domains.Service) Endpoint {
	return func(ctx *gin.Context, request any) (response any, err error) {
		return s.CreateUser(ctx, request.(*requests.CreateUser))
	}
}

func makeUpdateUserEndpoint(s domains.Service) Endpoint {
	return func(ctx *gin.Context, request any) (response any, err error) {
		return s.UpdateUser(ctx, request.(*requests.UpdateUser))
	}
}

func makeDeleteUserEndpoint(s domains.Service) Endpoint {
	return func(ctx *gin.Context, request any) (response any, err error) {
		return s.DeleteUser(ctx, request.(*requests.EntityId))
	}
}

func makeListUserEndpoint(s domains.Service) Endpoint {
	return func(ctx *gin.Context, request any) (response any, err error) {
		return s.GetAllUser(ctx)
	}
}
func makeViewUserEndpoint(s domains.Service) Endpoint {
	return func(ctx *gin.Context, request any) (response any, err error) {
		return s.ViewUser(ctx, request.(*requests.EntityId))
	}
}
