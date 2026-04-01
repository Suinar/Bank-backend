package core

type CardStatus int

const (
	CardStatusClosed CardStatus = iota
	CardStatusActive
	CardStatusBlocked
	StatusExpired
)

type Card struct {
	Id        uint64 `json:"id" db:"id"`
	UserId    uint64 `json:"user_id" db:"user_id"`
	AccountId uint64 `json:"account_id" db:"account_id"`

	Number string `json:"number" db:"number"`

	ExpiryMonth uint8 `json:"expiry_month" db:"expiry_month"`
	ExpiryYear  uint8 `json:"expiry_year" db:"expiry_year"`

	Status CardStatus `json:"status" db:"status"`
}

type CardCreateInput struct {
	UserId    string `json:"user_id" db:"user_id"`
	AccountId string `json:"account_id" db:"account_id"`
}
