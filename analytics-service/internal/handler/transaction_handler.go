package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/illenko/analytics-service/internal/logger"
	"github.com/illenko/analytics-service/internal/service"
	"log/slog"
	"net/http"
)

type TransactionHandler interface {
	FindAll(c *gin.Context)
	FindById(c *gin.Context)
}

type transactionHandler struct {
	log     *slog.Logger
	service service.TransactionService
}

func NewTransactionHandler(log *slog.Logger, service service.TransactionService) TransactionHandler {
	return &transactionHandler{
		log:     log,
		service: service,
	}
}

// FindAll
//
//	@Summary	    Retrieve transactions
//	@Schemes
//	@Tags			transaction
//	@Produce		json
//	@Success		200	{array}	model.TransactionResponse
//	@Router			/transactions [get]
func (t *transactionHandler) FindAll(c *gin.Context) {
	requestID := uuid.New()
	ctx := logger.AppendCtx(context.Background(), slog.String("requestID", requestID.String()))
	t.log.InfoContext(ctx, "Processing find all request")

	transactions, err := t.service.FindAll(ctx)

	if err != nil {
		t.error(c, err)
		return
	}

	t.success(c, transactions)
}

// FindById
//
//	@Summary	Retrieve transaction details
//	@Schemes
//	@Tags			transaction
//	@Produce		json
//	@Param          id   path      string  true  "Transaction ID"
//	@Success		200	{object}	model.TransactionResponse
//	@Router			/transactions/{id} [get]
func (t *transactionHandler) FindById(c *gin.Context) {
	requestID := uuid.New()
	ctx := logger.AppendCtx(context.Background(), slog.String("requestID", requestID.String()))
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		t.badRequest(c, "Id is not valid uuid")
		return
	}

	ctx = logger.AppendCtx(ctx, slog.String("transactionID", id.String()))
	t.log.InfoContext(ctx, "Processing find by id request")

	transaction, err := t.service.FindById(ctx, id)

	if err != nil {
		t.error(c, err)
		return
	}

	t.success(c, transaction)
}

func (t *transactionHandler) badRequest(c *gin.Context, message string) {
	t.response(c, http.StatusBadRequest, gin.H{"error": message})
}

func (t *transactionHandler) error(c *gin.Context, err error) {
	t.response(c, http.StatusInternalServerError, err)
}

func (t *transactionHandler) success(c *gin.Context, res interface{}) {
	t.response(c, http.StatusOK, res)
}

func (t *transactionHandler) response(c *gin.Context, status int, res interface{}) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(status, res)

}
