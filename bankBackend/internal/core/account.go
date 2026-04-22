package core

type AccountStatus int

const (
	AccountStatusClosed AccountStatus = iota
	AccountStatusActive
	AccountStatusBlocked
)

type Account struct {
	Id         int64 `json:"id" db:"id"`
	UserId     int64 `json:"user_id" db:"user_id"`
	CurrencyId int64 `json:"currency_id" db:"currency_id"`

	Name string `json:"name" db:"name"`

	Balance int64 `json:"balance" db:"balance"`

	Status AccountStatus `json:"status" db:"status"`
}

type AccountCreateInput struct {
	UserId     string `json:"user_id" db:"user_id"`
	CurrencyId string `json:"currency_id" db:"currency_id"`

	Name string `json:"name" db:"name"`
}

type AccountUpdateInput struct {
	Name *string `json:"name" db:"name"`
}
