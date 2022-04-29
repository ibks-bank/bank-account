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

func currencyToProto(cur entities.Currency) bank_account.Account_Currency {
	switch cur {
	case entities.Rub:
		return bank_account.Account_CURRENCY_RUB
	case entities.Euro:
		return bank_account.Account_CURRENCY_EURO
	case entities.DollarUs:
		return bank_account.Account_CURRENCY_DOLLAR_US
	default:
		return bank_account.Account_CURRENCY_UNKNOWN
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

func accountsToProto(accounts []*entities.Account) []*bank_account.Account {
	result := make([]*bank_account.Account, 0, len(accounts))
	for _, account := range accounts {
		result = append(result, accountToProto(account))
	}
	return result
}

func accountToProto(account *entities.Account) *bank_account.Account {
	return &bank_account.Account{
		Id:       account.ID,
		Currency: currencyToProto(account.Currency),
		Limit:    account.Limit,
		UserID:   account.UserID,
		Balance:  account.Balance,
	}
}
