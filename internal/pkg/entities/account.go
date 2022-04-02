package entities

type Account struct {
	ID       int64
	UserID   int64
	Currency Currency
	Limit    int64
	Balance  int64
	Name     string
}

type Currency string

const (
	Rub      Currency = "rub"
	Euro     Currency = "eur"
	DollarUs Currency = "usd"
)

func (c Currency) String() string {
	return string(c)
}
