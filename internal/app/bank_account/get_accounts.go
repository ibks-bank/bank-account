package bank_account

import (
	"context"

	bank_account "github.com/ibks-bank/bank-account/pkg/bank-account"
	"github.com/ibks-bank/libs/auth"
	"github.com/ibks-bank/libs/cerr"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (srv *Server) GetAccounts(ctx context.Context, _ *emptypb.Empty) (*bank_account.GetAccountsResponse, error) {
	userInfo, err := auth.GetUserInfo(ctx)
	if err != nil {
		return nil, cerr.WrapMC(err, "can't get user info from context", codes.Unauthenticated)
	}

	accounts, err := srv.accountUseCase.GetAccountsByUserID(ctx, userInfo.UserID)
	if err != nil {
		return nil, cerr.Wrap(err, "can't get accounts")
	}

	return &bank_account.GetAccountsResponse{Accounts: accountsToProto(accounts)}, nil
}
