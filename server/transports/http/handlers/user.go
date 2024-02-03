package handlers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/bagasunix/bank-ina/pkg/errors"
	"github.com/bagasunix/bank-ina/server/endpoints"
	"github.com/bagasunix/bank-ina/server/endpoints/requests"
)

func MakeUserHandler(endpoints endpoints.UserEndpoint, logger *zap.Logger, r *gin.RouterGroup) *gin.RouterGroup {
	r.POST("", makeCreatedUser(logger, endpoints))
	r.GET("", makeListUser(logger, endpoints))
	r.GET("/:id", makeViewUser(logger, endpoints))
	r.PUT("/:id", makeUpdateUser(logger, endpoints))
	r.DELETE("/:id", makeDeleteUser(logger, endpoints))
	return r
}

func makeCreatedUser(logger *zap.Logger, endpoints endpoints.UserEndpoint) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requests.CreateUser
		if err := ctx.Bind(&req); err != nil {
			EncodeError(ctx, err, ctx.Writer)
			return
		}
		resp, err := endpoints.CreateUserEndpoint(ctx, &req)
		if err != nil {
			EncodeError(ctx, err, ctx.Writer)
			return
		}

		EncodeOk(ctx, ctx.Writer, resp)
	}
}
func makeListUser(logger *zap.Logger, endpoints endpoints.UserEndpoint) gin.HandlerFunc {
	return func(c *gin.Context) {
		resp, err := endpoints.ListUserEndpoint(c, nil)
		if err != nil {
			EncodeError(c, err, c.Writer)
			return
		}

		EncodeOk(c, c.Writer, resp)
	}
}
func makeViewUser(logger *zap.Logger, endpoints endpoints.UserEndpoint) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := decodeByEntityIdEndpoint(c)
		if err != nil {
			EncodeError(c, err, c.Writer)
			return
		}
		resp, err := endpoints.ViewUserEndpoint(c, req)
		if err != nil {
			// return ResponseError(c, logger, resp, err)
			EncodeError(c, err, c.Writer)
			return
		}

		EncodeOk(c, c.Writer, resp)
	}
}
func makeUpdateUser(logger *zap.Logger, endpoints endpoints.UserEndpoint) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			EncodeError(c, errors.ErrInvalidAttributes("id"), c.Writer)
		}
		var req requests.UpdateUser
		req.ID = id
		if err := c.Bind(&req); err != nil {
			EncodeError(c, err, c.Writer)
			return
		}

		resp, err := endpoints.UpdateUserEnpoint(c, &req)
		if err != nil {
			EncodeError(c, err, c.Writer)
			return
		}

		EncodeOk(c, c.Writer, resp)
	}
}
func makeDeleteUser(logger *zap.Logger, endpoints endpoints.UserEndpoint) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := decodeByEntityIdEndpoint(c)
		if err != nil {
			EncodeError(c, err, c.Writer)
			return
		}
		resp, err := endpoints.DeleteUserEnpoint(c, id)
		if err != nil {
			EncodeError(c, err, c.Writer)
			return
		}

		EncodeOk(c, c.Writer, resp)
	}
}
