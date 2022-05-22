package bank_account

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation"
	bank_account "github.com/ibks-bank/bank-account/pkg/bank-account"
	"github.com/ibks-bank/libs/auth"
	"github.com/ibks-bank/libs/cerr"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (srv *Server) FillBalance(ctx context.Context, req *bank_account.FillBalanceRequest) (*emptypb.Empty, error) {
	err := validateFillBalanceRequest(req)
	if err != nil {
		return nil, cerr.WrapMC(err, "wrong request", codes.InvalidArgument)
	}

	userInfo, err := auth.GetUserInfo(ctx)
	if err != nil {
		return nil, cerr.WrapMC(err, "can't get user info from context", codes.Unauthenticated)
	}

	acc, err := srv.accountUseCase.GetAccountByID(ctx, req.GetAccountID())
	if err != nil {
		return nil, cerr.WrapMC(err, "account not found", codes.NotFound)
	}

	if acc.UserID != userInfo.UserID {
		return nil, cerr.NewC("account doesn't belong to this user", codes.Unauthenticated)
	}

	newBalance := acc.Balance + req.GetAmount()
	if acc.Limit < newBalance {
		return nil, cerr.NewC("exceeds account limit", codes.InvalidArgument)
	}

	err = srv.accountUseCase.UpdateAccountBalance(ctx, acc, newBalance)
	if err != nil {
		return nil, cerr.Wrap(err, "can't update account balance")
	}

	return &emptypb.Empty{}, nil
}

func validateFillBalanceRequest(req *bank_account.FillBalanceRequest) error {
	err := validation.Validate(req, validation.NotNil)
	if err != nil {
		return err
	}

	err = validation.ValidateStruct(req,
		validation.Field(&req.AccountID, validation.Required),
		validation.Field(&req.Amount, validation.Required),
	)
	if err != nil {
		return err
	}

	if req.GetAmount() <= 0 {
		return cerr.New("wrong amount")
	}

	return nil
}
