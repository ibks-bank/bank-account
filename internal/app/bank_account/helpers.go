package bank_account

import (
	"github.com/ibks-bank/bank-account/internal/pkg/entities"
	bank_account "github.com/ibks-bank/bank-account/pkg/bank-account"
	"github.com/ibks-bank/libs/cerr"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func currencyFromProto(cur bank_account.Account_Currency) (entities.Currency, error) {
	switch cur {
	case bank_account.Account_CURRENCY_RUB:
		return entities.Rub, nil
	case bank_account.Account_CURRENCY_EURO:
		return entities.Euro, nil
	case bank_account.Account_CURRENCY_DOLLAR_US:
		return entities.DollarUs, nil
	default:
		return "", cerr.New("wrong currency")
	}
}

func transactionsToProto(trxs []*entities.Transaction) []*bank_account.Transaction {
	trxsProto := make([]*bank_account.Transaction, 0, len(trxs))

	for _, trx := range trxs {
		trxsProto = append(trxsProto, &bank_account.Transaction{
			Id:          trx.ID,
			AccountFrom: trx.AccountFrom.ID,
			AccountTo:   trx.AccountTo.ID,
			Amount:      trx.Amount,
			Type:        transactionTypeToProto(trx.Type),
			Time:        timestamppb.New(trx.Time),
		})
	}

	return trxsProto
}

func transactionTypeToProto(t entities.TransactionType) bank_account.Transaction_Type {
	switch t {
	case entities.Transfer:
		return bank_account.Transaction_TYPE_TRANSFER
	case entities.Payment:
		return bank_account.Transaction_TYPE_PAYMENT
	default:
		return bank_account.Transaction_TYPE_UNKNOWN
	}
}
