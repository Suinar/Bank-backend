package core

type DepositStatus int

const (
	DepositStatusClosed DepositStatus = iota
	DepositStatusActive
	DepositStatusRejected
)

type Deposit struct {
	Id         string `json:"id" db:"id"`
	UserId     string `json:"user_id" db:"user_id"`
	CurrencyId string `json:"currency_id" db:"currency_id"`

	Principal    uint64 `json:"principal" db:"principal"`
	InterestRate uint8  `json:"interest_rate" db:"interest_rate"`

	TermMonths     uint8  `json:"term_months" db:"term_months"`
	MonthlyPayment uint32 `json:"monthly_payment" db:"monthly_payment"`

	Status DepositStatus `json:"status" db:"status"`
}

type DepositCreateInput struct {
	UserId     string `json:"user_id" db:"user_id"`
	CurrencyId string `json:"currency_id" db:"currency_id"`

	Principal    uint64 `json:"principal" db:"principal"`
	InterestRate uint8  `json:"interest_rate" db:"interest_rate"`

	TermMonths     uint8  `json:"term_months" db:"term_months"`
	MonthlyPayment uint32 `json:"monthly_payment" db:"monthly_payment"`
}
