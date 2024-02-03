package handlers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/bagasunix/bank-ina/server/endpoints"
	"github.com/bagasunix/bank-ina/server/endpoints/requests"
)

func MakeAuthHandler(endpoints endpoints.AuthEndpoint, logger *zap.Logger, r *gin.RouterGroup) *gin.RouterGroup {
	r.POST("login", makeSignIn(logger, endpoints))
	return r
}

func makeSignIn(logger *zap.Logger, endpoints endpoints.AuthEndpoint) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requests.Signin
		if err := ctx.Bind(&req); err != nil {
			EncodeError(ctx, err, ctx.Writer)
			return
		}
		resp, err := endpoints.SignInEndpoint(ctx, &req)
		if err != nil {
			EncodeError(ctx, err, ctx.Writer)
			return
		}

		EncodeOk(ctx, ctx.Writer, resp)
	}
}
