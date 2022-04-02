package usecase

import (
	"context"

	"github.com/ibks-bank/bank-account/internal/pkg/entities"
	"github.com/ibks-bank/bank-account/internal/pkg/transactioner"
)

type trxUseCase struct {
	repo transactioner.IRepo
}

func NewTransactionUseCase(repo transactioner.IRepo) *trxUseCase {
	return &trxUseCase{repo: repo}
}

func (trx *trxUseCase) GetTransactionsByAccountID(
	ctx context.Context,
	accountID int64,
	filter *entities.TransactionFilter,
) ([]*entities.Transaction, error) {

	return trx.repo.GetTransactionsByAccountID(ctx, accountID, filter)
}

func (trx *trxUseCase) CreateTransaction(ctx context.Context, amount, accountFrom, accountTo int64) (int64, error) {
	return trx.repo.CreateTransaction(ctx, amount, accountFrom, accountTo)
}
