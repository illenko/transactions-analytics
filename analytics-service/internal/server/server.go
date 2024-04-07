package server

import (
	"github.com/gin-gonic/gin"
	"github.com/illenko/analytics-service/docs"
	"github.com/illenko/analytics-service/internal/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func New(transactionHandler handler.TransactionHandler, analyticsHandler handler.AnalyticsHandler) *gin.Engine {
	e := gin.Default()
	e.GET("/transactions", transactionHandler.FindAll)
	e.GET("/transactions/:id", transactionHandler.FindById)
	e.GET("/analytics/income/groups", analyticsHandler.Income)
	e.GET("/analytics/expenses/groups", analyticsHandler.Expenses)
	e.GET("/analytics/income/dates", analyticsHandler.IncomeDates)
	e.GET("/analytics/expenses/dates", analyticsHandler.ExpensesDates)
	docs.SwaggerInfo.BasePath = "/"
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return e
}
