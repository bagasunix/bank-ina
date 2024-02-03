package handlers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/bagasunix/bank-ina/pkg/errors"
	"github.com/bagasunix/bank-ina/server/endpoints"
	"github.com/bagasunix/bank-ina/server/endpoints/requests"
)

func MakeTaskHandler(endpoints endpoints.TaskEndpoint, logger *zap.Logger, r *gin.RouterGroup) *gin.RouterGroup {
	r.POST("", makeCreatedTask(logger, endpoints))
	r.GET("", makeListTask(logger, endpoints))
	r.GET("/:id", makeViewTask(logger, endpoints))
	r.PUT("/:id", makeUpdateTask(logger, endpoints))
	r.DELETE("/:id", makeDeleteTask(logger, endpoints))
	return r
}

func makeCreatedTask(logger *zap.Logger, endpoints endpoints.TaskEndpoint) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requests.CreateTask
		if err := ctx.Bind(&req); err != nil {
			EncodeError(ctx, err, ctx.Writer)
			return
		}
		resp, err := endpoints.CreateTaskEndpoint(ctx, &req)
		if err != nil {
			EncodeError(ctx, err, ctx.Writer)
			return
		}

		EncodeOk(ctx, ctx.Writer, resp)
	}
}
func makeListTask(logger *zap.Logger, endpoints endpoints.TaskEndpoint) gin.HandlerFunc {
	return func(c *gin.Context) {
		resp, err := endpoints.ListTaskEndpoint(c, nil)
		if err != nil {
			EncodeError(c, err, c.Writer)
			return
		}

		EncodeOk(c, c.Writer, resp)
	}
}
func makeViewTask(logger *zap.Logger, endpoints endpoints.TaskEndpoint) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := decodeByEntityIdEndpoint(c)
		if err != nil {
			EncodeError(c, err, c.Writer)
			return
		}
		resp, err := endpoints.ViewTaskEndpoint(c, req)
		if err != nil {
			// return ResponseError(c, logger, resp, err)
			EncodeError(c, err, c.Writer)
			return
		}

		EncodeOk(c, c.Writer, resp)
	}
}
func makeUpdateTask(logger *zap.Logger, endpoints endpoints.TaskEndpoint) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			EncodeError(c, errors.ErrInvalidAttributes("id"), c.Writer)
		}
		var req requests.UpdateTask
		req.ID = id
		if err := c.Bind(&req); err != nil {
			EncodeError(c, err, c.Writer)
			return
		}

		resp, err := endpoints.UpdateTaskEnpoint(c, &req)
		if err != nil {
			EncodeError(c, err, c.Writer)
			return
		}

		EncodeOk(c, c.Writer, resp)
	}
}
func makeDeleteTask(logger *zap.Logger, endpoints endpoints.TaskEndpoint) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := decodeByEntityIdEndpoint(c)
		if err != nil {
			EncodeError(c, err, c.Writer)
			return
		}
		resp, err := endpoints.DeleteTaskEnpoint(c, id)
		if err != nil {
			EncodeError(c, err, c.Writer)
			return
		}

		EncodeOk(c, c.Writer, resp)
	}
}
