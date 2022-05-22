package bank_account

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/ibks-bank/bank-account/internal/pkg/entities"
	bank_account "github.com/ibks-bank/bank-account/pkg/bank-account"
	"github.com/ibks-bank/libs/auth"
	"github.com/ibks-bank/libs/cerr"
	"google.golang.org/grpc/codes"
)

func (srv *Server) GetAccountTransactions(ctx context.Context, req *bank_account.GetAccountTransactionsRequest) (
	*bank_account.GetAccountTransactionsResponse,
	error,
) {
	userInfo, err := auth.GetUserInfo(ctx)
	if err != nil {
		return nil, cerr.WrapMC(err, "can't get user info from context", codes.Unauthenticated)
	}

	err = validateGetAccountTransactionsRequest(req)
	if err != nil {
		return nil, cerr.WrapMC(err, "validation error", codes.InvalidArgument)
	}

	acc, err := srv.accountUseCase.GetAccountByID(ctx, req.GetAccountID())
	if err != nil {
		return nil, cerr.Wrap(err, "can't get account from by id")
	}

	if userInfo.UserID != acc.UserID {
		return nil, cerr.NewC("user ids are not equal", codes.Unauthenticated)
	}

	transactions, err := srv.trxUseCase.GetTransactionsByAccountID(ctx, acc.ID, buildFilter(req.GetFilterBy()))
	if err != nil {
		return nil, cerr.Wrap(err, "can't get transactions by id")
	}

	return &bank_account.GetAccountTransactionsResponse{Transactions: transactionsToProto(transactions)}, nil
}

func validateGetAccountTransactionsRequest(req *bank_account.GetAccountTransactionsRequest) error {
	err := validation.Validate(req, validation.NotNil)
	if err != nil {
		return err
	}

	err = validation.ValidateStruct(req,
		validation.Field(&req.AccountID, validation.Required),
	)
	if err != nil {
		return err
	}

	return nil
}

func buildFilter(filterProto *bank_account.GetAccountTransactionsRequest_FilterBy) *entities.TransactionFilter {
	if filterProto == nil {
		return nil
	}

	filter := new(entities.TransactionFilter)

	if filterProto.GetDateFrom() != nil {
		filter.DateFrom = filterProto.GetDateFrom().AsTime()
	}

	if filterProto.GetDateTo() != nil {
		filter.DateTo = filterProto.GetDateTo().AsTime()
	}

	return filter
}
