package inits

import (
	"go.uber.org/zap"

	"github.com/bagasunix/bank-ina/server/domains"
	"github.com/bagasunix/bank-ina/server/endpoints"
	"github.com/bagasunix/bank-ina/server/endpoints/middlewares"
)

func InitEndpoint(logger *zap.Logger, s domains.Service) endpoints.Endpoints {
	a := endpoints.NewBuilder()
	a.SetMiddlewares(getEndpointMiddleware(logger))
	a.SetService(s)
	return a.Build()
}

func getEndpointMiddleware(logger *zap.Logger) (mw map[string][]endpoints.Middleware) {
	mw = map[string][]endpoints.Middleware{}
	addDefaultEndpointMiddleware(logger, mw)
	return mw
}

func middlewaresWithAuthentication(logger *zap.Logger, method string) []endpoints.Middleware {
	mw := defaultMiddlewares(logger, method)
	// return mw
	return append(mw, middlewares.Authentication(logger))
}

func defaultMiddlewares(logger *zap.Logger, method string) []endpoints.Middleware {
	return []endpoints.Middleware{
		middlewares.Logging(*logger.With(zap.String("method", method))),
	}
}

func addDefaultEndpointMiddleware(logger *zap.Logger, mw map[string][]endpoints.Middleware) {
	mw[endpoints.CREATE_TASK] = middlewaresWithAuthentication(logger, endpoints.CREATE_TASK)
	mw[endpoints.UPDATE_TASK] = middlewaresWithAuthentication(logger, endpoints.UPDATE_TASK)
	mw[endpoints.DELETE_TASK] = middlewaresWithAuthentication(logger, endpoints.DELETE_TASK)
	mw[endpoints.View_TASK] = middlewaresWithAuthentication(logger, endpoints.View_TASK)
	mw[endpoints.LIST_TASK] = middlewaresWithAuthentication(logger, endpoints.LIST_TASK)

}
