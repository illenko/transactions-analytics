package database

import (
	"github.com/illenko/transactions-service/internal/transaction"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"testing"
)

func TestAnalytic(t *testing.T) {
	db := NewTestDatabase(t)
	defer testDB.Close(t)

	BeforeEach(t)
	r := transaction.NewRepository(slog.Default(), GetConnection(t))
	transactions, err := r.FindAll()

	assert.Nil(t, err)
	assert.NotNil(t, transactions)
	// println(testDB.ConnectionString(t))
}
