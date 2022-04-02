package bank_account

import (
	"github.com/ibks-bank/bank-account/internal/pkg/accounter"
	"github.com/ibks-bank/bank-account/internal/pkg/transactioner"
	bank_account "github.com/ibks-bank/bank-account/pkg/bank-account"
)

type Server struct {
	bank_account.UnimplementedBankAccountServer

	accountUseCase accounter.IUseCase
	trxUseCase     transactioner.IUseCase

	maxAccountLimit int64
}

func NewServer(
	accountUseCase accounter.IUseCase,
	trxUseCase transactioner.IUseCase,
	maxAccountLimit int64,
) *Server {

	return &Server{
		accountUseCase:  accountUseCase,
		trxUseCase:      trxUseCase,
		maxAccountLimit: maxAccountLimit,
	}
}
