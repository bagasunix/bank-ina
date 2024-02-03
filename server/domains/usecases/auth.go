package usecases

import (
	"context"
	"strconv"
	"time"

	"go.uber.org/zap"

	"github.com/bagasunix/bank-ina/pkg/errors"
	"github.com/bagasunix/bank-ina/pkg/hash"
	"github.com/bagasunix/bank-ina/pkg/jwt"
	"github.com/bagasunix/bank-ina/server/domains/data/repositories"
	"github.com/bagasunix/bank-ina/server/domains/entities"
	"github.com/bagasunix/bank-ina/server/endpoints/requests"
	"github.com/bagasunix/bank-ina/server/endpoints/responses"
)

type Auth interface {
	SignIn(ctx context.Context, request *requests.Signin) (response *responses.ViewEntity[*responses.SignIn], err error)
}

type authentication struct {
	repositories repositories.Repositories
	jwtKey       string
	logger       *zap.Logger
}

// SignIn implements Auth.
func (a *authentication) SignIn(ctx context.Context, request *requests.Signin) (response *responses.ViewEntity[*responses.SignIn], err error) {
	responseBuild := responses.NewViewEntityBuilder[*responses.SignIn]()
	if err = request.Validate(); err != nil {
		return responseBuild.Build(), err
	}

	checkMail := a.repositories.GetUser().GetByEmail(ctx, request.Email)
	if len(checkMail.Value) != 1 {
		return responseBuild.Build(), errors.CustomError(errors.ERR_NOT_FOUND)

	}

	if !hash.ComparePasswords(checkMail.Value[0].Password, []byte(request.Password)) {
		return responseBuild.Build(), errors.ErrInvalidAttributes("username and password")
	}

	userBuild := entities.NewUserBuilder()
	userBuild.SetID(strconv.Itoa(checkMail.Value[0].ID))
	userBuild.SetEmail(checkMail.Value[0].Email)
	userBuild.SetName(checkMail.Value[0].Name)

	clm := jwt.NewClaimsBuilder()
	clm.User(userBuild.Build())
	clm.ExpiresAt(time.Now().Add(1 * time.Hour))

	token, err := jwt.GenerateToken(a.jwtKey, *clm.Build())
	if err != nil {
		return responseBuild.Build(), err
	}

	resBuild := responses.NewSignInBuilder()
	resBuild.SetToken(token)

	responseBuild.SetData(resBuild.Build())
	return responseBuild.Build(), nil
}

func NewAuthentication(logger *zap.Logger, jwtKey string, repositories repositories.Repositories) Auth {
	a := new(authentication)
	a.logger = logger
	a.jwtKey = jwtKey
	a.repositories = repositories
	return a
}
