package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/bagasunix/bank-ina/server/endpoints"
)

func Logging(logger zap.Logger) endpoints.Middleware {
	return func(e endpoints.Endpoint) endpoints.Endpoint {
		return func(ctx *gin.Context, request any) (response any, err error) {
			defer func(begin time.Time) {
				logger.Info("transport_error",
					zap.Error(err),
					zap.Duration("took", time.Since(begin)),
				)
			}(time.Now())
			return e(ctx, request)
		}
	}
}
