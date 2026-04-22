package core

type DepositStatus int

const (
	DepositStatusClosed DepositStatus = iota
	DepositStatusActive
	DepositStatusRejected
)

type Deposit struct {
	Id         int64 `json:"id" db:"id"`
	UserId     int64 `json:"user_id" db:"user_id"`
	CurrencyId int64 `json:"currency_id" db:"currency_id"`

	Amount       int64   `json:"amount" db:"amount"`
	InterestRate float32 `json:"interest_rate" db:"interest_rate"`

	TermMonths int8 `json:"term_months" db:"term_months"`

	Status DepositStatus `json:"status" db:"status"`
}

type DepositCreateInput struct {
	UserId     string `json:"user_id" db:"user_id"`
	CurrencyId string `json:"currency_id" db:"currency_id"`

	Amount int64 `json:"amount" db:"amount"`

	TermMonths int8 `json:"term_months" db:"term_months"`
}
