package domains

import "github.com/bagasunix/bank-ina/server/domains/data/repositories"

type Middleware func(repo repositories.Repositories, contract Service) Service
