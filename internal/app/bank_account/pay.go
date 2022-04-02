package bank_account

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation"
	bank_account "github.com/ibks-bank/bank-account/pkg/bank-account"
	"github.com/ibks-bank/libs/cerr"
	"google.golang.org/grpc/codes"
)

func (srv *Server) Pay(ctx context.Context, req *bank_account.CreateTransactionRequest) (
	*bank_account.CreateTransactionResponse,
	error,
) {
	err := validateCreateTransactionRequest(req)
	if err != nil {
		return nil, cerr.WrapMC(err, "validation error", codes.InvalidArgument)
	}

	trxID, err := srv.trxUseCase.CreateTransaction(ctx, req.GetAmount(), req.GetAccountID(), req.GetPayee())
	if err != nil {
		return nil, cerr.Wrap(err, "can't create transaction")
	}

	return &bank_account.CreateTransactionResponse{TransactionID: trxID}, nil
}

func validateCreateTransactionRequest(req *bank_account.CreateTransactionRequest) error {
	err := validation.Validate(req, validation.NotNil)
	if err != nil {
		return err
	}

	err = validation.ValidateStruct(req,
		validation.Field(&req.AccountID, validation.Required),
		validation.Field(&req.Payee, validation.Required),
		validation.Field(&req.Amount, validation.Required),
	)
	if err != nil {
		return err
	}

	if req.GetAccountID() == req.GetPayee() {
		return cerr.New("accountID and payee can't be equal")
	}

	return nil
}
