package server

import (
	"github.com/gin-gonic/gin"
	"github.com/illenko/transactions-service/docs"
	"github.com/illenko/transactions-service/internal/analytic"
	"github.com/illenko/transactions-service/internal/transaction"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func New(transactionHandler transaction.Handler, analyticHandler analytic.Handler) *gin.Engine {
	e := gin.Default()
	e.GET("/transactions", transactionHandler.FindAll)
	e.GET("/transactions/:id", transactionHandler.FindById)
	e.GET("/analytic/income/groups", analyticHandler.Income)
	e.GET("/analytic/expenses/groups", analyticHandler.Expenses)
	e.GET("/analytic/income/dates", analyticHandler.IncomeDates)
	e.GET("/analytic/expenses/dates", analyticHandler.ExpensesDates)
	docs.SwaggerInfo.BasePath = "/"
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return e
}
