package accounter

import (
	"context"
	"database/sql"

	"github.com/ibks-bank/bank-account/internal/pkg/entities"
)

type IRepo interface {
	CreateAccount(ctx context.Context, account *entities.Account) (int64, error)
	GetAccountByID(ctx context.Context, id int64) (*entities.Account, error)
	GetAccountsByUserID(ctx context.Context, userID int64) ([]*entities.Account, error)
	UpdateAccount(ctx context.Context, account *entities.Account) error
	UpdateAccountTrx(ctx context.Context, trx *sql.Tx, account *entities.Account) error
}
