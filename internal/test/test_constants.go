package test

import "github.com/kVinsom/Bank-repository-service/pkg/core"

// Shared fixture values keep entity builders deterministic and internally consistent.
const (
	UserId           int64 = 1
	UserFirstName          = "John"
	UserMiddleName         = "Michael"
	UserLastName           = "Doe"
	UserEmail              = "john.doe@example.com"
	UserPhoneNumber        = "+380501234567"
	UserPasswordHash       = "$2a$10$test.password.hash"

	AccountId      int64 = 1
	AccountName          = "Main account"
	AccountBalance int64 = 100_000
	AccountStatus        = core.AccountStatusActive

	CardId          int64 = 1
	CardNumber            = "4242424242424242"
	CardExpiryMonth int8  = 12
	CardExpiryYear  int8  = 30
	CardStatus            = core.CardStatusActive

	CurrencyId         int64 = 1
	CurrencyName             = "US Dollar"
	CurrencySymbol     rune  = '$'
	CurrencyISOCode          = "USD"
	CurrencyMinorUnits int8  = 2

	CreditId             int64   = 1
	CreditAmount         int64   = 100_000
	CreditInterestRate   float32 = 12.5
	CreditTermMonths     int8    = 12
	CreditMonthlyPayment int64   = 8_900
	CreditStatus                 = core.CreditStatusActive

	DepositId           int64   = 1
	DepositAmount       int64   = 100_000
	DepositInterestRate float32 = 10.5
	DepositTermMonths   int8    = 12
	DepositStatus               = core.DepositStatusActive

	TransactionAmount = 1000
	ExchangeISOFrom   = 840
	ExchangeISOTo     = 978
)
