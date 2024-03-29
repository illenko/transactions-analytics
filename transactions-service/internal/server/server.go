package server

import (
	"github.com/gin-gonic/gin"
	"github.com/illenko/transactions-service/docs"
	"github.com/illenko/transactions-service/internal/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func New(handler handler.TransactionHandler) *gin.Engine {
	e := gin.Default()
	e.GET("/transactions", handler.FindAll)
	e.GET("/transactions/:id", handler.FindById)
	e.GET("/statistics/:by", handler.Statistics)
	e.GET("/merchants/:id/expenses", handler.MerchantExpenses)
	docs.SwaggerInfo.BasePath = "/"
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return e
}
