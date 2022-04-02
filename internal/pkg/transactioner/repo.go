package transactioner

import (
	"context"

	"github.com/ibks-bank/bank-account/internal/pkg/entities"
)

type IRepo interface {
	GetTransactionsByAccountID(ctx context.Context, accountID int64, filter *entities.TransactionFilter) (
		[]*entities.Transaction, error,
	)

	CreateTransaction(ctx context.Context, amount, accountFromID, accountToID int64) (int64, error)
}
