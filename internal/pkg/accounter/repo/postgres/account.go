package postgres

import (
	"context"
	"database/sql"
	"github.com/ibks-bank/bank-account/internal/pkg/accounter/repo/postgres/models"
	db "github.com/ibks-bank/bank-account/internal/pkg/db_helper"
	"github.com/ibks-bank/bank-account/internal/pkg/entities"
	"github.com/ibks-bank/libs/cerr"
)

func (st *store) CreateAccount(ctx context.Context, account *entities.Account) (int64, error) {
	accountM := &models.Account{
		Currency: account.Currency,
		Balance:  account.Balance,
		Limit:    account.Limit,
		UserID:   account.UserID,
		Name:     account.Name,
	}

	err := accountM.Insert(ctx, st.db)
	if err != nil {
		return 0, cerr.Wrap(err, "can't insert account")
	}

	return accountM.ID, nil
}

func (st *store) GetAccountByID(ctx context.Context, id int64) (*entities.Account, error) {
	account := &models.Account{ID: id}

	err := account.Select(ctx, st.db)
	if err != nil {
		return nil, cerr.Wrap(err, "can't select account")
	}

	return account.ToEntity(), nil
}

func (st *store) GetAccountsByUserID(ctx context.Context, userID int64) ([]*entities.Account, error) {
	accountsM := make([]*models.Account, 0)

	rows, err := st.db.QueryContext(ctx, "select * from accounts where user_id = $1", userID)
	if err != nil {
		return nil, cerr.Wrap(err, "can't exec query")
	}
	defer rows.Close()

	for rows.Next() {
		account := new(models.Account)

		err = rows.Scan(
			&account.ID,
			&account.CreatedAt,
			&account.Currency,
			&account.Balance,
			&account.Limit,
			&account.UserID,
			&account.Name,
		)
		if err != nil {
			return nil, cerr.Wrap(err, "can't scan row")
		}

		accountsM = append(accountsM, account)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	accounts := make([]*entities.Account, 0, len(accountsM))
	for _, accM := range accountsM {
		accounts = append(accounts, accM.ToEntity())
	}

	return accounts, nil
}

func (st *store) UpdateAccount(ctx context.Context, account *entities.Account) error {
	return updateAccount(ctx, st.db, account)
}

func (st *store) UpdateAccountTrx(ctx context.Context, trx *sql.Tx, account *entities.Account) error {
	return updateAccount(ctx, trx, account)
}

func updateAccount(ctx context.Context, exec db.IDatabase, account *entities.Account) error {
	return (&models.Account{
		ID:       account.ID,
		Currency: account.Currency,
		Balance:  account.Balance,
		Limit:    account.Limit,
		UserID:   account.UserID,
		Name:     account.Name,
	}).Update(ctx, exec)
}
