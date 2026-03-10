package core

import "time"

type Currency struct {
	Id string `json:"id" db:"id"`

	Name string `json:"name" db:"name"`

	Symbol     string `json:"symbol" db:"symbol"`
	IsoCode    string `json:"iso_code" db:"iso_code"`
	NumberCode string `json:"number_code" db:"number_code"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type CurrencyCreateInput struct {
	Name string `json:"name" db:"name"`

	Symbol     string `json:"symbol" db:"symbol"`
	IsoCode    string `json:"iso_code" db:"iso_code"`
	NumberCode string `json:"number_code" db:"number_code"`
}

type CurrencyUpdateInput struct {
	Name *string `json:"name" db:"name"`

	Symbol     *string `json:"symbol" db:"symbol"`
	IsoCode    *string `json:"iso_code" db:"iso_code"`
	NumberCode *string `json:"number_code" db:"number_code"`
}
