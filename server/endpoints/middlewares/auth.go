package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/bagasunix/bank-ina/pkg/errors"
	"github.com/bagasunix/bank-ina/pkg/jwt"
	"github.com/bagasunix/bank-ina/server/endpoints"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func Authentication(logs *zap.Logger) endpoints.Middleware {
	return func(e endpoints.Endpoint) endpoints.Endpoint {
		return func(ctx *gin.Context, request any) (response any, err error) {
			tokenString := ctx.GetHeader(authorizationHeaderKey)
			if tokenString == "" {
				return nil, errors.ErrUnAuthorized()
			}

			fields := strings.Fields(tokenString)
			if len(fields) < 2 {
				return nil, errors.ErrUnAuthorized()
			}

			authorizationType := strings.ToLower(fields[0])
			if authorizationType != authorizationTypeBearer {
				return nil, errors.ErrUnAuthorized()
			}

			claims, err := jwt.ValidateToken(fields[1])
			if err != nil {
				return nil, err
			}

			ctx.Set(authorizationPayloadKey, claims.User)
			return e(ctx, request)
		}
	}
}
