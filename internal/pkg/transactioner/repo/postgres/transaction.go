package postgres

import (
	"context"
	"database/sql"
	"github.com/ibks-bank/bank-account/internal/pkg/transactioner/repo/postgres/filter"

	"github.com/ibks-bank/bank-account/internal/pkg/entities"
	"github.com/ibks-bank/bank-account/internal/pkg/transactioner/repo/postgres/models"
	"github.com/ibks-bank/libs/cerr"
)

func (st *store) GetTransactionsByAccountID(
	ctx context.Context,
	accountID int64,
	trxFilter *entities.TransactionFilter,
) ([]*entities.Transaction, error) {

	return st.getTransactions(ctx, accountID, st.buildFilter(trxFilter)...)
}

func (st *store) getTransactions(
	ctx context.Context,
	accountID int64,
	filters ...filter.Filter,
) ([]*entities.Transaction, error) {

	trxsM := make([]*models.Transaction, 0)

	query := "select * from transactions where (\"account_from\" = $1 or \"account_to\" = $1)"
	args := []interface{}{accountID}

	for _, f := range filters {
		query, args = f(query, args)
	}

	rows, err := st.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, cerr.Wrap(err, "can't exec query")
	}
	defer rows.Close()

	for rows.Next() {
		trx := new(models.Transaction)

		err = rows.Scan(
			&trx.ID,
			&trx.CreatedAt,
			&trx.AccountTo,
			&trx.AccountToName,
			&trx.AccountFrom,
			&trx.Amount,
			&trx.Type,
		)
		if err != nil {
			return nil, cerr.Wrap(err, "can't scan row")
		}

		trxsM = append(trxsM, trx)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	trxs := make([]*entities.Transaction, 0, len(trxsM))
	for _, trxM := range trxsM {
		accountTo, err := st.accountRepo.GetAccountByID(ctx, trxM.AccountTo)
		if err != nil {
			return nil, cerr.Wrap(err, "can't get account_to by id")
		}

		accountFrom, err := st.accountRepo.GetAccountByID(ctx, trxM.AccountFrom)
		if err != nil {
			return nil, cerr.Wrap(err, "can't get account_from by id")
		}

		trxs = append(trxs, &entities.Transaction{
			ID:          trxM.ID,
			AccountTo:   accountTo,
			AccountFrom: accountFrom,
			Amount:      trxM.Amount,
			Type:        trxM.Type,
		})
	}

	return trxs, nil
}

func (st *store) buildFilter(trxFilter *entities.TransactionFilter) []filter.Filter {
	if trxFilter == nil {
		return nil
	}

	filters := make([]filter.Filter, 0)

	if !trxFilter.DateFrom.IsZero() {
		filters = append(filters, filter.ByDateFrom(trxFilter.DateFrom))
	}

	if !trxFilter.DateTo.IsZero() {
		filters = append(filters, filter.ByDateTo(trxFilter.DateFrom))
	}

	return filters
}

func (st *store) CreateTransaction(ctx context.Context, amount, accountFromID, accountToID int64) (int64, error) {
	accountFrom, err := st.accountRepo.GetAccountByID(ctx, accountFromID)
	if err != nil {
		return 0, cerr.Wrap(err, "can't get account_from by id")
	}

	accountTo, err := st.accountRepo.GetAccountByID(ctx, accountToID)
	if err != nil {
		return 0, cerr.Wrap(err, "can't get account_to by id")
	}

	trxType := entities.Payment
	if accountTo.UserID == 0 {
		trxType = entities.Transfer
	}

	trxM := &models.Transaction{
		AccountTo:     accountTo.ID,
		AccountToName: accountTo.Name,
		AccountFrom:   accountFrom.ID,
		Amount:        amount,
		Type:          trxType,
	}

	trxErr := st.WithTransaction(ctx, func(ctx context.Context, tx *sql.Tx) error {
		accountFrom.Balance -= amount
		if accountFrom.Balance < 0 {
			return entities.ErrNotEnoughMoney
		}

		err = st.accountRepo.UpdateAccount(ctx, accountFrom)
		if err != nil {
			return cerr.Wrap(err, "can't update account_from balance")
		}

		accountTo.Balance += amount
		if accountTo.Balance > accountTo.Limit {
			return entities.ErrMoreThanLimit
		}

		err = st.accountRepo.UpdateAccount(ctx, accountTo)
		if err != nil {
			return cerr.Wrap(err, "can't update account_to balance")
		}

		err = trxM.Insert(ctx, st.db)
		if err != nil {
			return cerr.Wrap(err, "can't insert transaction")
		}

		return nil
	})
	if trxErr != nil {
		return 0, cerr.Wrap(trxErr, "can't perform transaction")
	}

	return trxM.ID, nil
}
