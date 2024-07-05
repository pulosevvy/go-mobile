package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-mobile/config"
	"go-mobile/internal/handler/http"
	taskRepository "go-mobile/internal/repository/postgres/task"
	userRepository "go-mobile/internal/repository/postgres/user"
	"go-mobile/internal/service/task"
	"go-mobile/internal/service/user"
	"go-mobile/package/database/postgres"
	"go-mobile/package/httpserver"
	sl "go-mobile/package/logger/slog"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {

	//Log init
	log := sl.SetupLogger(cfg.Env)
	log.Info("starting app", slog.String("env", cfg.Env))

	//PG Connection
	pg, err := postgres.NewPostgres(&cfg.PG)
	if err != nil {
		log.Error("failed to connect to postgres", sl.Err(err))
		os.Exit(1)
	}
	log.Info("database is connecting", slog.String("db:", cfg.PG.Database))
	defer func(pg *postgres.Postgres, ctx context.Context) {
		err := pg.Close(ctx)
		if err != nil {
			log.Error("failed to close postgres connection", sl.Err(err))
			os.Exit(1)
		}
	}(pg, context.Background())

	//Services init
	userService := service.NewUserService(userRepository.NewUserRepository(pg), "")
	taskService := task.NewTaskService(taskRepository.NewTaskRepository(pg))

	//Routes init
	app := gin.New()
	http.NewControllers(app, log, userService, taskService)

	//Http Server
	log.Info("starting server", slog.String("env", cfg.Env))

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	httpServer := httpserver.NewHttpServer(app, &cfg.HttpService)
	go func() {
		if err := httpServer.Start(); err != nil {
			log.Error("failed to start http server", sl.Err(err))
			os.Exit(1)
		}
	}()
	log.Info("starting app", slog.String("address", cfg.HttpService.Address))

	<-done
	log.Info("stopping server")
	err = httpServer.Shutdown()
	if err != nil {
		log.Error("failed to stop http server", sl.Err(err))
		os.Exit(1)
	}

	log.Info("server stopped")
}
