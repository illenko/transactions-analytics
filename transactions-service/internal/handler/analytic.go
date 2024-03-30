package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/illenko/transactions-service/internal/service"
	"log/slog"
)

type AnalyticHandler interface {
	Income(c *gin.Context)
	Expenses(c *gin.Context)
	IncomeDates(c *gin.Context)
	ExpensesDates(c *gin.Context)
}

type analyticHandler struct {
	log     *slog.Logger
	service service.AnalyticService
}

func NewAnalyticHandler(log *slog.Logger, service service.AnalyticService) AnalyticHandler {
	return &analyticHandler{
		log:     log,
		service: service,
	}
}

func (a analyticHandler) Income(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (a analyticHandler) Expenses(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (a analyticHandler) IncomeDates(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (a analyticHandler) ExpensesDates(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
