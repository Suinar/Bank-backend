package core

type AccountStatus int

const (
	AccountStatusClosed AccountStatus = iota
	AccountStatusActive
	AccountStatusBlocked
)

type AccountType int

const (
	AccountTypeClosed AccountType = iota
	AccountTypeDeposit
	AccountTypeCredit
)

type Account struct {
	Id         string `json:"id" db:"id"`
	UserId     string `json:"user_id" db:"user_id"`
	CurrencyId string `json:"currency_id" db:"currency_id"`

	Name string `json:"name" db:"name"`

	Balance uint64 `json:"balance" db:"balance"`

	Status AccountStatus `json:"status" db:"status"`
	Type   AccountType   `json:"type" db:"type"`
}

type AccountCreateInput struct {
	UserId     string `json:"user_id" db:"user_id"`
	CurrencyId string `json:"currency_id" db:"currency_id"`

	Name string `json:"name" db:"name"`

	Type AccountType `json:"type" db:"type"`
}

type AccountUpdateInput struct {
	Name *string `json:"name" db:"name"`
}
