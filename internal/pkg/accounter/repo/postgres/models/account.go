package models

import (
	"context"
	"database/sql"
	"time"

	"github.com/ibks-bank/bank-account/internal/pkg/entities"
	"github.com/ibks-bank/libs/cerr"
)

type Account struct {
	ID        int64
	CreatedAt time.Time
	Currency  entities.Currency
	Balance   int64
	Limit     int64
	UserID    int64
	Name      string
}

func (a *Account) Insert(ctx context.Context, db *sql.DB) error {
	query := `insert 
			  into accounts("currency", "balance", "limit", "user_id", "name") 
			  values ($1, $2, $3, $4, $5) 
			  returning "id"`

	err := db.QueryRowContext(ctx, query, a.Currency, a.Balance, a.Limit, a.UserID, a.Name).Scan(&a.ID)
	if err != nil {
		return cerr.Wrap(err, "can't exec query")
	}

	return nil
}

func (a *Account) Select(ctx context.Context, db *sql.DB) error {
	err := db.QueryRowContext(
		ctx,
		"select * from accounts where \"id\" = $1",
		a.ID,
	).Scan(&a.ID, &a.CreatedAt, &a.Currency, &a.Balance, &a.Limit, &a.UserID, &a.Name)
	if err != nil {
		return cerr.Wrap(err, "can't exec query")
	}

	return nil
}

func (a *Account) Update(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(
		ctx,
		"update accounts set \"balance\" = $1, \"limit\" = $2, \"name\" = $3 where \"id\" = $4",
		a.Balance, a.Limit, a.Name, a.ID,
	)
	if err != nil {
		return cerr.Wrap(err, "can't exec query")
	}

	return nil
}

func (a *Account) ToEntity() *entities.Account {
	return &entities.Account{
		ID:       a.ID,
		UserID:   a.UserID,
		Currency: a.Currency,
		Limit:    a.Limit,
		Balance:  a.Balance,
		Name:     a.Name,
	}
}
