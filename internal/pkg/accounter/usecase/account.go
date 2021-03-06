package usecase

import (
	"context"

	"github.com/ibks-bank/bank-account/internal/pkg/accounter"
	"github.com/ibks-bank/bank-account/internal/pkg/entities"
	"github.com/ibks-bank/bank-account/internal/pkg/transactioner"
)

type account struct {
	accountRepo accounter.IRepo
	trxUseCase  transactioner.IUseCase
}

func NewAccountUseCase(accountRepo accounter.IRepo, trxUseCase transactioner.IUseCase) *account {
	return &account{accountRepo: accountRepo, trxUseCase: trxUseCase}
}

func (a *account) CreateAccount(ctx context.Context, account *entities.Account) (int64, error) {
	return a.accountRepo.CreateAccount(ctx, account)
}

func (a *account) GetAccountByID(ctx context.Context, id int64) (*entities.Account, error) {
	return a.accountRepo.GetAccountByID(ctx, id)
}

func (a *account) GetAccountsByUserID(ctx context.Context, userID int64) ([]*entities.Account, error) {
	return a.accountRepo.GetAccountsByUserID(ctx, userID)
}

func (a *account) TransferMoney(ctx context.Context, amount int64, accountFrom, accountTo *entities.Account) (int64, error) {
	return a.trxUseCase.CreateTransaction(ctx, amount, accountFrom, accountTo)
}

func (a *account) UpdateAccountBalance(ctx context.Context, acc *entities.Account, newBalance int64) error {
	acc.Balance = newBalance

	return a.accountRepo.UpdateAccount(ctx, acc)
}

func (a *account) UpdateAccountLimit(ctx context.Context, acc *entities.Account, newLimit int64) error {
	acc.Limit = newLimit

	return a.accountRepo.UpdateAccount(ctx, acc)
}
