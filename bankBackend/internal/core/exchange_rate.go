package core

type ExchangeRate struct {
	CurrencyIdFrom int64   `json:"currency_id_from" db:"currency_id_from"`
	CurrencyIdTo   int64   `json:"currency_id_to" db:"currency_id_to"`
	RateSell       float32 `json:"rate_sell"`
	RateBuy        float32 `json:"rate_buy"`
	RateCross      float32 `json:"rate_cross"`
}
