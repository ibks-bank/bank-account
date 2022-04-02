package models

import (
	"context"
	"database/sql"
	"time"

	"github.com/ibks-bank/bank-account/internal/pkg/entities"
	"github.com/ibks-bank/libs/cerr"
)

type Transaction struct {
	ID            int64
	CreatedAt     time.Time
	AccountTo     int64
	AccountToName string
	AccountFrom   int64
	Amount        int64
	Type          entities.TransactionType
}

func (t *Transaction) Insert(ctx context.Context, db *sql.DB) error {
	query := `insert 
			  into transactions("account_to", "account_to_name", "account_from", "amount", "type") 
			  values ($1, $2, $3, $4, $5) 
			  returning "id"`

	err := db.QueryRowContext(ctx, query, t.AccountTo, t.AccountToName, t.AccountFrom, t.Amount, t.Type).Scan(&t.ID)
	if err != nil {
		return cerr.Wrap(err, "can't exec query")
	}

	return nil
}

func (t *Transaction) Select(ctx context.Context, db *sql.DB) error {
	err := db.QueryRowContext(
		ctx,
		"select * from transactions where \"id\" = $1",
		t.ID,
	).Scan(&t.ID, &t.CreatedAt, &t.AccountTo, &t.AccountToName, &t.AccountFrom, &t.Amount, &t.Type)
	if err != nil {
		return cerr.Wrap(err, "can't exec query")
	}

	return nil
}
