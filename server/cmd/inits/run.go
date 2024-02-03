package inits

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"github.com/bagasunix/bank-ina/pkg/envs"
	"github.com/bagasunix/bank-ina/server/domains/data/repositories"
)

var (
	configs, _ = envs.LoadEnv()
	httpAddr   = flag.String("http.addr", ":"+configs.Port, "HTTP listen address")
)

func Run() {
	ctx := context.Background()
	logger := InitLogger()
	db := InitDb(ctx, logger, configs)
	Migrate(logger, db)
	repositories := repositories.New(logger, db)
	svc := InitService(logger, repositories)
	eps := InitEndpoint(logger, svc)

	httpHandler := InitHttpHandler(configs, logger, eps)
	errs := make(chan error)
	defer close(errs)
	go initCancel(errs)
	go initHttp(httpHandler, errs)

	logger.Error("exit", zap.Error(<-errs))
}

func initCancel(errs chan error) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	errs <- fmt.Errorf("%s", <-c)
}

func initHttp(httpHandler http.Handler, errs chan error) {
	server := &http.Server{
		Addr:    *httpAddr,
		Handler: httpHandler,
	}
	errs <- server.ListenAndServe()
}
