package bank_account

import (
	"context"

	bank_account "github.com/ibks-bank/bank-account/pkg/bank-account"
	"github.com/ibks-bank/libs/auth"
	"github.com/ibks-bank/libs/cerr"
	"google.golang.org/grpc/codes"
)

func (srv *Server) GetAccount(ctx context.Context, req *bank_account.GetAccountRequest) (*bank_account.Account, error) {
	userInfo, err := auth.GetUserInfo(ctx)
	if err != nil {
		return nil, cerr.WrapMC(err, "can't get user info from context", codes.Unauthenticated)
	}

	account, err := srv.accountUseCase.GetAccountByID(ctx, req.GetAccountID())
	if err != nil {
		return nil, cerr.Wrap(err, "can't get account by id")
	}

	if account.UserID != userInfo.UserID {
		return nil, cerr.NewC("wrong user id", codes.InvalidArgument)
	}

	return accountToProto(account), nil
}
