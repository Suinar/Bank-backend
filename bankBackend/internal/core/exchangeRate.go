package core

type ExchangeRate struct {
	CurrencyIdFrom string `json:"currency_id_from" db:"currency_id_from"`
	CurrencyIdTo   string `json:"currency_id_to" db:"currency_id_to"`
	Almost         int64  `json:"almost"`
}
