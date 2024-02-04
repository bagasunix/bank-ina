package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/bagasunix/bank-ina/pkg/envs"
	"github.com/bagasunix/bank-ina/server/endpoints"
	"github.com/bagasunix/bank-ina/server/endpoints/middlewares"
	"github.com/bagasunix/bank-ina/server/transports/http/handlers"
)

func NewHttpHandler(configs *envs.Configs, loggers *zap.Logger, endpoints endpoints.Endpoints) http.Handler {
	gin.SetMode(configs.Env)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middlewares.Secure())
	r.Use(middlewares.CORSMiddleware())
	r.Use(middlewares.RemoveTrailingSlash())

	handlers.MakeUserHandler(endpoints.UserEndpoint, loggers, r.Group("/user"))
	handlers.MakeTaskHandler(endpoints.TaskEndpoint, loggers, r.Group("/task"))
	handlers.MakeAuthHandler(endpoints.AuthEndpoint, loggers, r.Group("/"))

	return r
}
