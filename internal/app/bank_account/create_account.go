package bank_account

import (
	"context"

	"github.com/ibks-bank/bank-account/internal/pkg/cerr"
	"github.com/ibks-bank/bank-account/pkg/bank_account"
)

func (srv *Server) CreateAccount(ctx context.Context, request *bank_account.CreateAccountRequest) (*bank_account.CreateAccountResponse, error) {
	//TODO implement me
	return nil, cerr.New("method not implemented")
}
