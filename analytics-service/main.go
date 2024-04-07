package main

import (
	"github.com/gin-gonic/gin"
	"github.com/illenko/analytics-service/internal/config"
	"github.com/illenko/analytics-service/internal/database"
	"github.com/illenko/analytics-service/internal/handler"
	"github.com/illenko/analytics-service/internal/logger"
	"github.com/illenko/analytics-service/internal/mapper"
	"github.com/illenko/analytics-service/internal/server"
	"github.com/illenko/analytics-service/internal/service"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"log/slog"
)

// @title           Analytics Service API
// @version         1.0
func main() {
	fx.New(
		fx.Provide(
			logger.New,
			config.Get,
			database.NewConnection,
			database.NewMigration,
			database.NewTransactionRepository,
			database.NewAnalyticsRepository,
			mapper.NewTransactionMapper,
			mapper.NewAnalyticsMapper,
			service.NewTransactionService,
			service.NewAnalyticsService,
			handler.NewTransactionHandler,
			handler.NewAnalyticsHandler,
			server.New,
		),
		fx.Invoke(func(e *gin.Engine, migration database.Migration, config config.AppConfig) {
			err := migration.Execute("migrations")
			if err != nil {
				return
			}
			err = e.Run(":" + config.Server.Port)
			if err != nil {
				return
			}
		}),
		fx.WithLogger(func(log *slog.Logger) fxevent.Logger {
			return &fxevent.SlogLogger{Logger: log}
		}),
	).Run()
}
