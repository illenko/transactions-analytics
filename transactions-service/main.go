package main

import (
	"github.com/gin-gonic/gin"
	"github.com/illenko/transactions-service/internal/config"
	"github.com/illenko/transactions-service/internal/database"
	"github.com/illenko/transactions-service/internal/handler"
	"github.com/illenko/transactions-service/internal/logger"
	"github.com/illenko/transactions-service/internal/mapper"
	"github.com/illenko/transactions-service/internal/server"
	"github.com/illenko/transactions-service/internal/service"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"log/slog"
)

// @title           Transactions Service API
// @version         1.0
func main() {
	fx.New(
		fx.Provide(
			logger.New,
			config.Get,
			database.NewConnection,
			database.NewTransactionRepository,
			database.NewAnalyticRepository,
			mapper.NewTransactionMapper,
			mapper.NewAnalyticMapper,
			service.NewTransactionService,
			service.NewAnalyticService,
			handler.NewTransactionHandler,
			handler.NewAnalyticHandler,
			server.New,
		),
		fx.Invoke(func(e *gin.Engine, config config.AppConfig) {
			err := e.Run(":" + config.Server.Port)
			if err != nil {
				return
			}
		}),
		fx.WithLogger(func(log *slog.Logger) fxevent.Logger {
			return &fxevent.SlogLogger{Logger: log}
		}),
	).Run()
}
