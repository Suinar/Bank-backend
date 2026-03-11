package core

import "time"

type Currency struct {
	Id uint64 `json:"id" db:"id"`

	Name string `json:"name" db:"name"`

	Symbol     string `json:"symbol" db:"symbol"`
	IsoCode    string `json:"iso_code" db:"iso_code"`
	NumberCode string `json:"number_code" db:"number_code"`

	MinorUnits int `json:"minor_units" db:"minor_units"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type CurrencyCreateInput struct {
	Name string `json:"name" db:"name"`

	Symbol     string `json:"symbol" db:"symbol"`
	IsoCode    string `json:"iso_code" db:"iso_code"`
	NumberCode string `json:"number_code" db:"number_code"`

	MinorUnits int `json:"minor_units" db:"minor_units"`
}

type CurrencyUpdateInput struct {
	Name *string `json:"name" db:"name"`

	Symbol     *string `json:"symbol" db:"symbol"`
	IsoCode    *string `json:"iso_code" db:"iso_code"`
	NumberCode *string `json:"number_code" db:"number_code"`

	MinorUnits *int `json:"minor_units" db:"minor_units"`
}
