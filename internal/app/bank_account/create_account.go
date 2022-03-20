package bank_account

import (
	"context"

	"github.com/ibks-bank/bank-account/pkg/bank-account"
	"github.com/ibks-bank/libs/cerr"
)

func (srv *Server) CreateAccount(ctx context.Context, request *bank_account.CreateAccountRequest) (*bank_account.CreateAccountResponse, error) {
	//TODO implement me
	return nil, cerr.New("method not implemented")
}
