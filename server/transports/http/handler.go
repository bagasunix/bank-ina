package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/bagasunix/bank-ina/pkg/envs"
	"github.com/bagasunix/bank-ina/server/endpoints"
	"github.com/bagasunix/bank-ina/server/transports/http/handlers"
)

func NewHttpHandler(configs *envs.Configs, loggers *zap.Logger, endpoints endpoints.Endpoints) http.Handler {
	gin.SetMode(configs.Env)
	r := gin.New()
	r.Use(gin.Recovery())

	handlers.MakeUserHandler(endpoints.UserEndpoint, loggers, r.Group(configs.ApiVersion+"/user"))

	return r
}
