package accounter

import (
	"context"

	"github.com/ibks-bank/bank-account/internal/pkg/entities"
)

type IUseCase interface {
	CreateAccount(ctx context.Context, account *entities.Account) (int64, error)
	GetAccountByID(ctx context.Context, id int64) (*entities.Account, error)
	GetAccountsByUserID(ctx context.Context, userID int64) ([]*entities.Account, error)
	TransferMoney(ctx context.Context, amount int64, accountFrom, accountTo *entities.Account) (int64, error)
	UpdateAccountBalance(ctx context.Context, accID, balance int64) error
	UpdateAccountLimit(ctx context.Context, accID, limit int64) error
}
