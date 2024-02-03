package domains

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/bagasunix/bank-ina/server/domains/data/repositories"
	"github.com/bagasunix/bank-ina/server/domains/usecases"
)

type Service interface {
	usecases.User
	usecases.Task
	usecases.Auth
}

type service struct {
	usecases.User
	usecases.Task
	usecases.Auth
}

type ServiceBuilder struct {
	logger       *zap.Logger
	repositories repositories.Repositories
	jwtKey       string
	ctx          *gin.Context
	middlewares  []Middleware
}

func NewServiceBuilder(logger *zap.Logger, jwtKey string, repositories repositories.Repositories) *ServiceBuilder {
	a := new(ServiceBuilder)
	a.logger = logger
	a.repositories = repositories
	a.jwtKey = jwtKey
	return a
}

func buildService(logger *zap.Logger, repo repositories.Repositories, jwtKey string) Service {
	svc := new(service)
	svc.User = usecases.NewUser(logger, repo)
	svc.Task = usecases.NewTask(logger, repo)
	svc.Auth = usecases.NewAuthentication(logger, jwtKey, repo)
	return svc
}

// SetMiddlewares Setter method for the field middlewares of type []Middleware in the object ServiceBuilder
func (s *ServiceBuilder) SetMiddlewares(middlewares []Middleware) *ServiceBuilder {
	s.middlewares = middlewares
	return s
}

func (s *ServiceBuilder) Build() Service {
	svc := buildService(s.logger, s.repositories, s.jwtKey)
	for _, m := range s.middlewares {
		svc = m(s.repositories, svc)
	}
	return svc
}
