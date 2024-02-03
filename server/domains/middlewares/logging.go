package middlewares

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/bagasunix/bank-ina/server/domains"
	"github.com/bagasunix/bank-ina/server/domains/data/repositories"
	"github.com/bagasunix/bank-ina/server/domains/entities"
	"github.com/bagasunix/bank-ina/server/endpoints/requests"
	"github.com/bagasunix/bank-ina/server/endpoints/responses"
)

type loggingMiddleware struct {
	logger *zap.Logger
	next   domains.Service
}

// SignIn implements domains.Service.
func (l *loggingMiddleware) SignIn(ctx context.Context, request *requests.Signin) (response *responses.ViewEntity[*responses.SignIn], err error) {
	defer func(begin time.Time) {
		l.logger.Info("BANK-INA", zap.String("request", string(request.ToJSON())), zap.String("response", string(response.ToJSON())), zap.Error(err), zap.Duration("took", time.Since(begin)))
	}(time.Now())
	return l.next.SignIn(ctx, request)
}

// CreateTask implements domains.Service.
func (l *loggingMiddleware) CreateTask(ctx *gin.Context, req *requests.CreateTask) (response *responses.EntityId, err error) {
	defer func(begin time.Time) {
		l.logger.Info("BANK-INA", zap.String("request", string(req.ToJSON())), zap.String("response", string(response.ToJSON())), zap.Error(err), zap.Duration("took", time.Since(begin)))
	}(time.Now())
	return l.next.CreateTask(ctx, req)
}

// CreateUser implements domains.Service.
func (l *loggingMiddleware) CreateUser(ctx *gin.Context, req *requests.CreateUser) (response *responses.EntityId, err error) {
	defer func(begin time.Time) {
		l.logger.Info("BANK-INA", zap.String("request", string(req.ToJSON())), zap.String("response", string(response.ToJSON())), zap.Error(err), zap.Duration("took", time.Since(begin)))
	}(time.Now())
	return l.next.CreateUser(ctx, req)
}

// DeleteTask implements domains.Service.
func (l *loggingMiddleware) DeleteTask(ctx *gin.Context, req *requests.EntityId) (response *responses.Empty, err error) {
	defer func(begin time.Time) {
		l.logger.Info("BANK-INA", zap.String("request", string(req.ToJSON())), zap.String("response", string(response.ToJSON())), zap.Error(err), zap.Duration("took", time.Since(begin)))
	}(time.Now())
	return l.next.DeleteTask(ctx, req)
}

// DeleteUser implements domains.Service.
func (l *loggingMiddleware) DeleteUser(ctx *gin.Context, req *requests.EntityId) (response *responses.Empty, err error) {
	defer func(begin time.Time) {
		l.logger.Info("BANK-INA", zap.String("request", string(req.ToJSON())), zap.String("response", string(response.ToJSON())), zap.Error(err), zap.Duration("took", time.Since(begin)))
	}(time.Now())
	return l.next.DeleteUser(ctx, req)
}

// GetAllTask implements domains.Service.
func (l *loggingMiddleware) GetAllTask(ctx *gin.Context) (response *responses.ListEntity[entities.Task], err error) {
	defer func(begin time.Time) {
		l.logger.Info("BANK-INA", zap.String("request", ""), zap.String("response", string(response.ToJSON())), zap.Error(err), zap.Duration("took", time.Since(begin)))
	}(time.Now())
	return l.next.GetAllTask(ctx)
}

// GetAllUser implements domains.Service.
func (l *loggingMiddleware) GetAllUser(ctx *gin.Context) (response *responses.ListEntity[entities.User], err error) {
	defer func(begin time.Time) {
		l.logger.Info("BANK-INA", zap.String("request", ""), zap.String("response", string(response.ToJSON())), zap.Error(err), zap.Duration("took", time.Since(begin)))
	}(time.Now())
	return l.next.GetAllUser(ctx)
}

// UpdateTask implements domains.Service.
func (l *loggingMiddleware) UpdateTask(ctx *gin.Context, req *requests.UpdateTask) (response *responses.ViewEntity[*entities.Task], err error) {
	defer func(begin time.Time) {
		l.logger.Info("BANK-INA", zap.String("request", string(req.ToJSON())), zap.String("response", string(response.ToJSON())), zap.Error(err), zap.Duration("took", time.Since(begin)))
	}(time.Now())
	return l.next.UpdateTask(ctx, req)
}

// UpdateUser implements domains.Service.
func (l *loggingMiddleware) UpdateUser(ctx *gin.Context, req *requests.UpdateUser) (response *responses.ViewEntity[*entities.User], err error) {
	defer func(begin time.Time) {
		l.logger.Info("BANK-INA", zap.String("request", ""), zap.String("response", string(response.ToJSON())), zap.Error(err), zap.Duration("took", time.Since(begin)))
	}(time.Now())
	return l.next.UpdateUser(ctx, req)
}

// ViewTask implements domains.Service.
func (l *loggingMiddleware) ViewTask(ctx *gin.Context, req *requests.EntityId) (response *responses.ViewEntity[*entities.Task], err error) {
	defer func(begin time.Time) {
		l.logger.Info("BANK-INA", zap.String("request", string(req.ToJSON())), zap.String("response", string(response.ToJSON())), zap.Error(err), zap.Duration("took", time.Since(begin)))
	}(time.Now())
	return l.next.ViewTask(ctx, req)
}

// ViewUser implements domains.Service.
func (l *loggingMiddleware) ViewUser(ctx *gin.Context, req *requests.EntityId) (response *responses.ViewEntity[*entities.User], err error) {
	defer func(begin time.Time) {
		l.logger.Info("BANK-INA", zap.String("request", string(req.ToJSON())), zap.String("response", string(response.ToJSON())), zap.Error(err), zap.Duration("took", time.Since(begin)))
	}(time.Now())
	return l.next.ViewUser(ctx, req)
}

func LoggingMiddleware(logger *zap.Logger, repo repositories.Repositories) domains.Middleware {
	return func(repo repositories.Repositories, next domains.Service) domains.Service {
		return &loggingMiddleware{logger: logger, next: next}
	}
}
