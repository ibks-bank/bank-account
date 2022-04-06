package usecase

import (
	"context"

	"github.com/ibks-bank/bank-account/internal/pkg/accounter"
	"github.com/ibks-bank/bank-account/internal/pkg/entities"
	"github.com/ibks-bank/bank-account/internal/pkg/transactioner"
	"github.com/ibks-bank/libs/cerr"
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

func (a *account) UpdateAccountBalance(ctx context.Context, accID, balance int64) error {
	acc, err := a.GetAccountByID(ctx, accID)
	if err != nil {
		return cerr.Wrap(err, "can't get account by id")
	}
	acc.Balance = balance

	return a.accountRepo.UpdateAccount(ctx, acc)
}

func (a *account) UpdateAccountLimit(ctx context.Context, accID, limit int64) error {
	acc, err := a.GetAccountByID(ctx, accID)
	if err != nil {
		return cerr.Wrap(err, "can't get account by id")
	}
	acc.Limit = limit

	return a.accountRepo.UpdateAccount(ctx, acc)
}
