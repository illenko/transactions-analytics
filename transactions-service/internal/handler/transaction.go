package handler

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/illenko/transactions-service/internal/logger"
	"github.com/illenko/transactions-service/internal/service"
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
		t.serverError(ctx, c, err)
		return
	}

	t.successResponse(ctx, c, transactions)
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
		t.badRequest(ctx, c, "Id is not valid uuid")
		return
	}

	ctx = logger.AppendCtx(ctx, slog.String("transactionID", id.String()))
	t.log.InfoContext(ctx, "Processing find by id request")

	transaction, err := t.service.FindById(ctx, id)

	if err != nil {
		t.serverError(ctx, c, err)
		return
	}

	t.successResponse(ctx, c, transaction)
}

func (t *transactionHandler) badRequest(ctx context.Context, c *gin.Context, message string) {
	t.logAndReturn(ctx, c, http.StatusBadRequest, gin.H{"error": message})
}

func (t *transactionHandler) serverError(ctx context.Context, c *gin.Context, err error) {
	t.logAndReturn(ctx, c, http.StatusInternalServerError, err)
}

func (t *transactionHandler) successResponse(ctx context.Context, c *gin.Context, res interface{}) {
	t.logAndReturn(ctx, c, http.StatusOK, res)
}

func (t *transactionHandler) logAndReturn(ctx context.Context, c *gin.Context, status int, res interface{}) {
	t.log.InfoContext(ctx, fmt.Sprintf("Returned response: %v, %v", status, spew.Sdump(res)))
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(status, res)

}
