package postgres

import (
	"context"
	"database/sql"

	"github.com/ibks-bank/bank-account/internal/pkg/accounter"
	"github.com/ibks-bank/libs/cerr"
)

var (
	ErrNotFound      = cerr.New("not found")
	ErrAlreadyExists = cerr.New("already exists")
)

const errViolatesUnique = "violates unique"

type store struct {
	db *sql.DB

	accountRepo accounter.IRepo
}

func NewTransactionRepo(db *sql.DB, accountUseCase accounter.IRepo) *store {
	return &store{db: db, accountRepo: accountUseCase}
}

type TxFunc func(ctx context.Context, tx *sql.Tx) error

func (st *store) WithTransaction(ctx context.Context, fn TxFunc) (err error) {
	tx, err := st.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	err = fn(ctx, tx)
	return err
}
