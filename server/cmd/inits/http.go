package inits

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/bagasunix/bank-ina/pkg/envs"
	"github.com/bagasunix/bank-ina/server/endpoints"
	transportHttp "github.com/bagasunix/bank-ina/server/transports/http"
)

func InitHttpHandler(configs *envs.Configs, logger *zap.Logger, endpoints endpoints.Endpoints) http.Handler {
	return transportHttp.NewHttpHandler(configs, logger, endpoints)
}
