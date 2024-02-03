package endpoints

import "github.com/bagasunix/bank-ina/server/domains"

type Endpoints struct {
	UserEndpoint UserEndpoint
	TaskEndpoint TaskEndpoint
}

// Builder EndpointsBuilder Builder Object for Endpoints
type Builder struct {
	service     domains.Service
	middlewares map[string][]Middleware
}

// NewBuilder Constructor for EndpointsBuilder
func NewBuilder() *Builder {
	o := new(Builder)
	return o
}

// Build Method which creates Endpoints
func (b *Builder) Build() Endpoints {
	eps := new(Endpoints)
	eps.UserEndpoint = NewUserEndpoint(b.service, b.middlewares)
	eps.TaskEndpoint = NewTaskEndpoint(b.service, b.middlewares)
	return *eps
}

// SetService Setter method for the field service of type domains.Service in the object EndpointsBuilder
func (b *Builder) SetService(service domains.Service) *Builder {
	b.service = service
	return b
}

// SetMiddlewares Setter method for the field middlewares of type map[string][]endpoint.Middleware in the object EndpointsBuilder
func (b *Builder) SetMiddlewares(middlewares map[string][]Middleware) *Builder {
	b.middlewares = middlewares
	return b
}
