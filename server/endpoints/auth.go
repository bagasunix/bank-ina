package endpoints

import (
	"github.com/gin-gonic/gin"

	"github.com/bagasunix/bank-ina/server/domains"
	"github.com/bagasunix/bank-ina/server/endpoints/requests"
)

const (
	SIGNIN = "Signin"
)

type AuthEndpoint struct {
	SignInEndpoint Endpoint
}

func NewAuthEndpoint(s domains.Service, mdw map[string][]Middleware) AuthEndpoint {
	eps := AuthEndpoint{}
	eps.SignInEndpoint = makeSigninEndpoint(s)

	SetEndpoint(SIGNIN, &eps.SignInEndpoint, mdw)

	return eps
}

func makeSigninEndpoint(s domains.Service) Endpoint {
	return func(ctx *gin.Context, request any) (response any, err error) {
		return s.SignIn(ctx, request.(*requests.Signin))
	}
}
