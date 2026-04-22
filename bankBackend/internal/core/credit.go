package core

type CreditStatus int

const (
	CreditStatusClosed CreditStatus = iota
	CreditStatusActive
	CreditStatusRejected
)

type Credit struct {
	Id         int64 `json:"id" db:"id"`
	UserId     int64 `json:"user_id" db:"user_id"`
	CurrencyId int64 `json:"currency_id" db:"currency_id"`

	Amount       int64   `json:"amount" db:"amount"`
	InterestRate float32 `json:"interest_rate" db:"interest_rate"`

	TermMonths     int8  `json:"term_months" db:"term_months"`
	MonthlyPayment int64 `json:"monthly_payment" db:"monthly_payment"`

	Status CreditStatus `json:"status" db:"status"`
}

type CreditCreateInput struct {
	UserId     string `json:"user_id" db:"user_id"`
	CurrencyId string `json:"currency_id" db:"currency_id"`

	Amount     int64 `json:"amount" db:"amount"`
	TermMonths int8  `json:"term_months" db:"term_months"`
}
