package entities

import (
	"time"

	"github.com/ibks-bank/libs/cerr"
)

var (
	ErrNotEnoughMoney = cerr.New("account_from doesn't have enough money")
	ErrMoreThanLimit  = cerr.New("exceeds account_to limit ")
)

type Transaction struct {
	ID          int64
	AccountTo   *Account
	AccountFrom *Account
	Amount      int64
	Type        TransactionType
	Time        time.Time
}

type TransactionType string

const (
	Transfer TransactionType = "transfer"
	Payment  TransactionType = "payment"
)

func (t TransactionType) String() string {
	return string(t)
}

type TransactionFilter struct {
	DateFrom time.Time
	DateTo   time.Time

	WithIncomes  bool
	WithExpenses bool
}
