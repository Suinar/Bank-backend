package core

type Currency struct {
	Id int64 `json:"id" db:"id"`

	Name string `json:"name" db:"name"`

	Symbol  rune   `json:"symbol" db:"symbol"`
	IsoCode string `json:"iso_code" db:"iso_code"`

	MinorUnits int8 `json:"minor_units" db:"minor_units"`
}

type CurrencyCreateInput struct {
	Name string `json:"name" db:"name"`

	Symbol  rune   `json:"symbol" db:"symbol"`
	IsoCode string `json:"iso_code" db:"iso_code"`

	MinorUnits int8 `json:"minor_units" db:"minor_units"`
}

type CurrencyUpdateInput struct {
	Name *string `json:"name" db:"name"`

	Symbol     *rune   `json:"symbol" db:"symbol"`
	IsoCode    *string `json:"iso_code" db:"iso_code"`
	NumberCode *string `json:"number_code" db:"number_code"`

	MinorUnits *int `json:"minor_units" db:"minor_units"`
}
