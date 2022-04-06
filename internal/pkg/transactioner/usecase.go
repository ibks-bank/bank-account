package transactioner

import (
	"context"

	"github.com/ibks-bank/bank-account/internal/pkg/entities"
)

type IUseCase interface {
	GetTransactionsByAccountID(ctx context.Context, accountID int64, filter *entities.TransactionFilter) (
		[]*entities.Transaction, error,
	)

	CreateTransaction(ctx context.Context, amount int64, accountFrom, accountTo *entities.Account) (int64, error)
}
