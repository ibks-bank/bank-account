package bank_account

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/ibks-bank/bank-account/internal/pkg/entities"
	"github.com/ibks-bank/bank-account/pkg/bank-account"
	"github.com/ibks-bank/libs/auth"
	"github.com/ibks-bank/libs/cerr"
	"google.golang.org/grpc/codes"
)

func (srv *Server) CreateAccount(ctx context.Context, req *bank_account.CreateAccountRequest) (
	*bank_account.CreateAccountResponse,
	error,
) {
	userInfo, err := auth.GetUserInfo(ctx)
	if err != nil {
		return nil, cerr.WrapMC(err, "can't get user info from context", codes.Unauthenticated)
	}

	err = validateCreateAccountRequest(req)
	if err != nil {
		return nil, cerr.WrapMC(err, "validation error", codes.InvalidArgument)
	}

	currency, err := currencyFromProto(req.GetCurrency())
	if err != nil {
		return nil, cerr.Wrap(err, "can't convert currency from proto")
	}

	limit := req.GetLimit()
	if limit == 0 {
		limit = srv.maxAccountLimit
	}

	accountID, err := srv.accountUseCase.CreateAccount(ctx, &entities.Account{
		UserID:   userInfo.UserID,
		Currency: currency,
		Limit:    limit,
		Balance:  0,
		Name:     req.GetName(),
	})
	if err != nil {
		return nil, cerr.Wrap(err, "can't create account")
	}

	return &bank_account.CreateAccountResponse{AccountID: accountID}, nil
}

func validateCreateAccountRequest(req *bank_account.CreateAccountRequest) error {
	err := validation.Validate(req, validation.NotNil)
	if err != nil {
		return err
	}

	err = validation.ValidateStruct(req,
		validation.Field(&req.Currency, validation.Required),
		validation.Field(&req.Name, validation.Required),
	)
	if err != nil {
		return err
	}

	return nil
}
