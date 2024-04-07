package handler

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/illenko/analytics-service/internal/logger"
	"github.com/illenko/analytics-service/internal/service"
	"github.com/samber/lo"
	"log/slog"
	"net/http"
)

type AnalyticsHandler interface {
	Income(c *gin.Context)
	Expenses(c *gin.Context)
	IncomeDates(c *gin.Context)
	ExpensesDates(c *gin.Context)
}

type analyticsHandler struct {
	log     *slog.Logger
	service service.AnalyticsService
}

func NewAnalyticsHandler(log *slog.Logger, service service.AnalyticsService) AnalyticsHandler {
	return &analyticsHandler{
		log:     log,
		service: service,
	}
}

type Groups string

var groups = []string{"category", "merchant"}
var units = []string{"day", "month"}
var calculations = []string{"absolute", "cumulative"}

const (
	defaultGroup       = "category"
	defaultUnit        = "month"
	defaultCalculation = "absolute"
)

// Income
//
//	@Summary	    Income analytics for groups
//	@Param   		group  query     string     false  "Grouping field"       default(category) Enums(category, merchant)
//	@Tags			income
//	@Produce		json
//	@Success		200	{array}	model.AnalyticsResponse
//	@Router			/analytics/income/groups [get]
func (h *analyticsHandler) Income(c *gin.Context) {
	h.analytics(c, "income")
}

// Expenses
//
//	@Summary	    Expenses analytics for groups
//	@Param   		group  query     string     false  "Grouping field"       default(category) Enums(category, merchant)
//	@Tags			expenses
//	@Produce		json
//	@Success		200	{array}	model.AnalyticsResponse
//	@Router			/analytics/expenses/groups [get]
func (h *analyticsHandler) Expenses(c *gin.Context) {
	h.analytics(c, "expenses")
}

func (h *analyticsHandler) analytics(c *gin.Context, direction string) {
	ctx := h.buildContext()
	group, ok := c.GetQuery("group")
	if !ok || !lo.Contains(groups, group) {
		group = defaultGroup
	}

	result, err := h.service.Analytics(ctx, direction, group)

	if err != nil {
		h.error(ctx, c, err)
		return
	}

	h.success(ctx, c, result)
}

// IncomeDates
//
//	@Summary	    Income analytics for dates
//	@Param   		unit  query     string     false  "Date unit"       default(month) Enums(month, day)
//	@Tags			income
//	@Produce		json
//	@Success		200	{array}	model.AnalyticsResponse
//	@Router			/analytics/income/dates [get]
func (h *analyticsHandler) IncomeDates(c *gin.Context) {
	h.analyticsDates(c, "income")
}

// ExpensesDates
//
//	@Summary	    Expenses analytics for dates
//	@Param   		unit  query     string     false  "Date unit"       default(month) Enums(month, day)
//	@Param   		calculation  query     string     false  "Calculation type"       default(absolute) Enums(absolute, cumulative)
//	@Tags			expenses
//	@Produce		json
//	@Success		200	{array}	model.AnalyticsResponse
//	@Router			/analytics/expenses/dates [get]
func (h *analyticsHandler) ExpensesDates(c *gin.Context) {
	h.analyticsDates(c, "expenses")
}

func (h *analyticsHandler) analyticsDates(c *gin.Context, direction string) {
	ctx := h.buildContext()

	unit, ok := c.GetQuery("unit")

	if !ok || !lo.Contains(units, unit) {
		unit = defaultUnit
	}

	calculation, ok := c.GetQuery("calculation")

	if !ok || !lo.Contains(calculations, calculation) {
		calculation = defaultCalculation
	}

	result, err := h.service.AnalyticsByDates(ctx, direction, unit, calculation)

	if err != nil {
		h.error(ctx, c, err)
		return
	}

	h.success(ctx, c, result)
}

func (h *analyticsHandler) buildContext() context.Context {
	return logger.AppendCtx(context.Background(), slog.String("requestID", uuid.New().String()))
}

func (h *analyticsHandler) error(ctx context.Context, c *gin.Context, err error) {
	h.response(ctx, c, http.StatusInternalServerError, err)
}

func (h *analyticsHandler) success(ctx context.Context, c *gin.Context, res interface{}) {
	h.response(ctx, c, http.StatusOK, res)
}

func (h *analyticsHandler) response(ctx context.Context, c *gin.Context, status int, res interface{}) {
	h.log.InfoContext(ctx, fmt.Sprintf("Returned response: %v, %v", status, spew.Sdump(res)))
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(status, res)

}
