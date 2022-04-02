package bank_account

import (
	"context"
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/ibks-bank/bank-account/internal/pkg/entities"
	bank_account "github.com/ibks-bank/bank-account/pkg/bank-account"
	"github.com/ibks-bank/libs/auth"
	"github.com/ibks-bank/libs/cerr"
	"google.golang.org/grpc/codes"
)

func (srv *Server) Pay(ctx context.Context, req *bank_account.CreateTransactionRequest) (
	*bank_account.CreateTransactionResponse,
	error,
) {
	userInfo, err := auth.GetUserInfo(ctx)
	if err != nil {
		return nil, cerr.WrapMC(err, "can't get user info from context", codes.Unauthenticated)
	}

	err = validateCreateTransactionRequest(req)
	if err != nil {
		return nil, cerr.WrapMC(err, "validation error", codes.InvalidArgument)
	}

	trxID, err := srv.accountUseCase.TransferMoney(ctx, req.GetAmount(), userInfo.UserID, req.GetPayee())
	if err != nil {
		if errors.Is(err, entities.ErrNotEnoughMoney) || errors.Is(err, entities.ErrMoreThanLimit) {
			return nil, cerr.WrapMC(err, "", codes.InvalidArgument)
		}

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
		validation.Field(&req.Payee, validation.Required),
		validation.Field(&req.Amount, validation.Required),
	)
	if err != nil {
		return err
	}

	return nil
}
