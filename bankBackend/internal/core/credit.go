package core

type CreditStatus int

const (
	CreditStatusRejected CreditStatus = iota
	CreditStatusActive
	CreditStatusClosed
)

type Credit struct {
	Id         string `json:"id" db:"id"`
	UserId     string `json:"user_id" db:"user_id"`
	CurrencyId string `json:"currency_id" db:"currency_id"`

	Amount       uint64 `json:"amount" db:"amount"`
	InterestRate uint8  `json:"interest_rate" db:"interest_rate"`
	TermMonths   uint8  `json:"term_months" db:"term_months"`

	Status CreditStatus `json:"status" db:"status"`
}

type CreditCreateInput struct {
	UserId     string `json:"user_id" db:"user_id"`
	CurrencyId string `json:"currency_id" db:"currency_id"`

	Amount       uint64 `json:"amount" db:"amount"`
	InterestRate uint8  `json:"interest_rate" db:"interest_rate"`
	TermMonths   uint8  `json:"term_months" db:"term_months"`
}
