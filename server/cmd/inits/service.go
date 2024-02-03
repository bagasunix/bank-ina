package inits

import (
	"go.uber.org/zap"

	"github.com/bagasunix/bank-ina/pkg/envs"
	"github.com/bagasunix/bank-ina/server/domains"
	"github.com/bagasunix/bank-ina/server/domains/data/repositories"
	"github.com/bagasunix/bank-ina/server/domains/middlewares"
)

func InitService(logger *zap.Logger, conf *envs.Configs, repo repositories.Repositories) domains.Service {
	serviceBuilder := domains.NewServiceBuilder(logger, conf.JWTKey, repo)
	serviceBuilder.SetMiddlewares(getServiceMiddleware(logger))
	return serviceBuilder.Build()
}

func getServiceMiddleware(logger *zap.Logger) []domains.Middleware {
	var mw []domains.Middleware
	var repo repositories.Repositories
	mw = addDefaultServiceMiddleware(logger, mw, repo)
	return mw
}

func addDefaultServiceMiddleware(logger *zap.Logger, mw []domains.Middleware, repo repositories.Repositories) []domains.Middleware {
	return append(mw, middlewares.LoggingMiddleware(logger, repo))
}
