package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_logger "github.com/DSTR19/ToDo/internal/core/logger"
	core_postgres_pool "github.com/DSTR19/ToDo/internal/core/repository/postgres/pool"
	core_http_middleware "github.com/DSTR19/ToDo/internal/core/transport/http/middleware"
	core_http_server "github.com/DSTR19/ToDo/internal/core/transport/http/server"
	users_postgres_repository "github.com/DSTR19/ToDo/internal/features/users/repository/postgres"
	users_service "github.com/DSTR19/ToDo/internal/features/users/service"
	users_transport_http "github.com/DSTR19/ToDo/internal/features/users/transport/http"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer cancel()
	fmt.Println("Hello, ToDo App!")

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		fmt.Println("Failed to init application logger:", err)
		os.Exit(1)
	}
	defer logger.Close()
	logger.Debug("Initialiaing postgres connection pool")
	pool, err := core_postgres_pool.NewConnectionPool(
		ctx,
		core_postgres_pool.NewConfigMust(),
	)
	if err != nil {
		logger.Fatal("Failed to init postgres connection pool", zap.Error(err))

	}
	defer pool.Close()

	logger.Debug("Initializing feature", zap.String("Feature", "users"))
	usersRepository := users_postgres_repository.NewUsersPostgresRepository(pool)
	usersService := users_service.NewUsersService(usersRepository)

	user_TransportHTTP := users_transport_http.NewUsersHTTPHandler(usersService)
	logger.Debug("Initializing HTTP Server")
	userRoutes := user_TransportHTTP.Routes()

	apiVersionRouter := core_http_server.NewApiVersionRouter(core_http_server.ApiVersionV1)
	apiVersionRouter.RegisterRoute(userRoutes)

	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Panic())
	httpServer.RegisterAPIRouters([]core_http_server.ApiVersionRouter{*apiVersionRouter})
	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server stopped with error:", zap.Error(err))
	}
}
