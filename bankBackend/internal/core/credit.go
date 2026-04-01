package core

type CreditStatus int

const (
	CreditStatusClosed CreditStatus = iota
	CreditStatusActive
	CreditStatusRejected
)

type Credit struct {
	Id         uint64 `json:"id" db:"id"`
	UserId     uint64 `json:"user_id" db:"user_id"`
	CurrencyId uint64 `json:"currency_id" db:"currency_id"`

	Amount       uint64  `json:"amount" db:"amount"`
	InterestRate float32 `json:"interest_rate" db:"interest_rate"`

	TermMonths     uint8  `json:"term_months" db:"term_months"`
	MonthlyPayment uint64 `json:"monthly_payment" db:"monthly_payment"`

	Status CreditStatus `json:"status" db:"status"`
}

type CreditCreateInput struct {
	UserId     string `json:"user_id" db:"user_id"`
	CurrencyId string `json:"currency_id" db:"currency_id"`

	Amount     uint64 `json:"amount" db:"amount"`
	TermMonths uint8  `json:"term_months" db:"term_months"`
}
