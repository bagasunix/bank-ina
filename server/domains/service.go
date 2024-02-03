package domains

import (
	"go.uber.org/zap"

	"github.com/bagasunix/bank-ina/server/domains/data/repositories"
	"github.com/bagasunix/bank-ina/server/domains/usecases"
)

type Service interface {
	usecases.User
	usecases.Task
}

type service struct {
	usecases.User
	usecases.Task
}

type ServiceBuilder struct {
	logger       *zap.Logger
	repositories repositories.Repositories
	middlewares  []Middleware
}

func NewServiceBuilder(logger *zap.Logger, repositories repositories.Repositories) *ServiceBuilder {
	a := new(ServiceBuilder)
	a.logger = logger
	a.repositories = repositories
	return a
}

func buildService(logger *zap.Logger, repo repositories.Repositories) Service {
	svc := new(service)
	svc.User = usecases.NewUser(logger, repo)
	svc.Task = usecases.NewTask(logger, repo)
	return svc
}

// SetMiddlewares Setter method for the field middlewares of type []Middleware in the object ServiceBuilder
func (s *ServiceBuilder) SetMiddlewares(middlewares []Middleware) *ServiceBuilder {
	s.middlewares = middlewares
	return s
}

func (s *ServiceBuilder) Build() Service {
	svc := buildService(s.logger, s.repositories)
	for _, m := range s.middlewares {
		svc = m(s.repositories, svc)
	}
	return svc
}
