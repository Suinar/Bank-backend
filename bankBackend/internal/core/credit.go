package core

type CreditStatus int

const (
	CreditStatusRejected CreditStatus = iota
	CreditStatusActive
	CreditStatusClosed
)

type Credit struct {
	Id         uint64 `json:"id" db:"id"`
	UserId     uint64 `json:"user_id" db:"user_id"`
	CurrencyId uint64 `json:"currency_id" db:"currency_id"`

	Amount       uint64 `json:"amount" db:"amount"`
	InterestRate uint8  `json:"interest_rate" db:"interest_rate"`
	TermMonths   uint8  `json:"term_months" db:"term_months"`

	Status CreditStatus `json:"status" db:"status"`
}

type CreditCreateInput struct {
	UserId     uint64 `json:"user_id" db:"user_id"`
	CurrencyId uint64 `json:"currency_id" db:"currency_id"`

	Amount       uint64 `json:"amount" db:"amount"`
	InterestRate uint8  `json:"interest_rate" db:"interest_rate"`
	TermMonths   uint8  `json:"term_months" db:"term_months"`
}
