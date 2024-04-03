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
//	@Summary	    Income analytic for groups
//	@Param   		group  query     string     false  "Grouping field"       default(category) Enums(category, merchant)
//	@Tags			income
//	@Produce		json
//	@Success		200	{array}	model.AnalyticResponse
//	@Router			/analytic/income/groups [get]
func (h *analyticHandler) Income(c *gin.Context) {
	h.analytic(c, "income")
}

// Expenses
//
//	@Summary	    Expenses analytic for groups
//	@Param   		group  query     string     false  "Grouping field"       default(category) Enums(category, merchant)
//	@Tags			expenses
//	@Produce		json
//	@Success		200	{array}	model.AnalyticResponse
//	@Router			/analytic/expenses/groups [get]
func (h *analyticHandler) Expenses(c *gin.Context) {
	h.analytic(c, "expenses")
}

func (h *analyticHandler) analytic(c *gin.Context, direction string) {
	ctx := h.buildContext()
	group, ok := c.GetQuery("group")
	if !ok || !lo.Contains(groups, group) {
		group = defaultGroup
	}

	result, err := h.service.Analytic(ctx, direction, group)

	if err != nil {
		h.error(ctx, c, err)
		return
	}

	h.success(ctx, c, result)
}

// IncomeDates
//
//	@Summary	    Income analytic for dates
//	@Param   		unit  query     string     false  "Date unit"       default(month) Enums(month, day)
//	@Tags			income
//	@Produce		json
//	@Success		200	{array}	model.AnalyticResponse
//	@Router			/analytic/income/dates [get]
func (h *analyticHandler) IncomeDates(c *gin.Context) {
	h.analyticDates(c, "income")
}

// ExpensesDates
//
//	@Summary	    Expenses analytic for dates
//	@Param   		unit  query     string     false  "Date unit"       default(month) Enums(month, day)
//	@Param   		calculation  query     string     false  "Calculation type"       default(absolute) Enums(absolute, cumulative)
//	@Tags			expenses
//	@Produce		json
//	@Success		200	{array}	model.AnalyticResponse
//	@Router			/analytic/expenses/dates [get]
func (h *analyticHandler) ExpensesDates(c *gin.Context) {
	h.analyticDates(c, "expenses")
}

func (h *analyticHandler) analyticDates(c *gin.Context, direction string) {
	ctx := h.buildContext()

	unit, ok := c.GetQuery("unit")

	if !ok || !lo.Contains(units, unit) {
		unit = defaultUnit
	}

	calculation, ok := c.GetQuery("calculation")

	if !ok || !lo.Contains(calculations, calculation) {
		calculation = defaultCalculation
	}

	result, err := h.service.AnalyticByDates(ctx, direction, unit, calculation)

	if err != nil {
		h.error(ctx, c, err)
		return
	}

	h.success(ctx, c, result)
}

func (h *analyticHandler) buildContext() context.Context {
	return logger.AppendCtx(context.Background(), slog.String("requestID", uuid.New().String()))
}

func (h *analyticHandler) error(ctx context.Context, c *gin.Context, err error) {
	h.response(ctx, c, http.StatusInternalServerError, err)
}

func (h *analyticHandler) success(ctx context.Context, c *gin.Context, res interface{}) {
	h.response(ctx, c, http.StatusOK, res)
}

func (h *analyticHandler) response(ctx context.Context, c *gin.Context, status int, res interface{}) {
	h.log.InfoContext(ctx, fmt.Sprintf("Returned response: %v, %v", status, spew.Sdump(res)))
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(status, res)

}
