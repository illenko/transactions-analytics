package handler

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/illenko/transactions-service/internal/logger"
	"github.com/illenko/transactions-service/internal/service"
	"github.com/samber/lo"
	"log/slog"
	"net/http"
	"strconv"
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

var groups = []string{"category", "merchant"}
var units = []string{"day", "month"}

const defaultGroupBy = "category"
const defaultUnit = "month"
const defaultDayPeriod = 7
const defaultMonthPeriod = 6

func (h *analyticHandler) Income(c *gin.Context) {
	h.analytic(c, "income")
}

func (h *analyticHandler) Expenses(c *gin.Context) {
	h.analytic(c, "expenses")
}

func (h *analyticHandler) IncomeDates(c *gin.Context) {
	h.analyticDates(c, "income")
}

func (h *analyticHandler) ExpensesDates(c *gin.Context) {
	h.analyticDates(c, "expenses")
}

func (h *analyticHandler) analytic(c *gin.Context, analyticType string) {
	ctx := h.buildContext()
	groupBy, ok := c.GetQuery("groupBy")
	if !ok || !lo.Contains(groups, groupBy) {
		groupBy = defaultGroupBy
	}

	result, err := h.service.Analytic(ctx, analyticType, groupBy)

	if err != nil {
		h.serverError(ctx, c, err)
		return
	}

	h.successResponse(ctx, c, result)
}

func (h *analyticHandler) analyticDates(c *gin.Context, analyticType string) {
	ctx := h.buildContext()

	unit, ok := c.GetQuery("unit")

	if !ok || !lo.Contains(units, unit) {
		unit = defaultUnit
	}

	periodString, ok := c.GetQuery("period")
	period, err := strconv.Atoi(periodString)

	if !ok || err != nil || period < 0 {
		if unit == "day" {
			period = defaultDayPeriod
		} else {
			period = defaultMonthPeriod
		}
	}

	result, err := h.service.AnalyticByDates(ctx, analyticType, unit, period, c.Query("category"), c.Query("merchant"))

	if err != nil {
		h.serverError(ctx, c, err)
		return
	}

	h.successResponse(ctx, c, result)
}

func (h *analyticHandler) buildContext() context.Context {
	return logger.AppendCtx(context.Background(), slog.String("requestID", uuid.New().String()))
}

func (h *analyticHandler) badRequest(ctx context.Context, c *gin.Context, message string) {
	h.logAndReturn(ctx, c, http.StatusBadRequest, gin.H{"error": message})
}

func (h *analyticHandler) serverError(ctx context.Context, c *gin.Context, err error) {
	h.logAndReturn(ctx, c, http.StatusInternalServerError, err)
}

func (h *analyticHandler) successResponse(ctx context.Context, c *gin.Context, res interface{}) {
	h.logAndReturn(ctx, c, http.StatusOK, res)
}

func (h *analyticHandler) logAndReturn(ctx context.Context, c *gin.Context, status int, res interface{}) {
	h.log.InfoContext(ctx, fmt.Sprintf("Returned response: %v, %v", status, spew.Sdump(res)))
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(status, res)

}
