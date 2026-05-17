package core

import (
	"context"
	"database/sql"
	"errors"
	"time"

	core "github.com/Suinar/Bank-backend/bankBackend/internal/core"
	"github.com/jmoiron/sqlx"
)

type CardRepository struct {
	db *sqlx.DB
}

func NewCardRepository(db *sqlx.DB) *CardRepository {
	return &CardRepository{db: db}
}

func (r *CardRepository) GetAll(ctx context.Context) ([]core.Card, error) {
	query := `
SELECT id, user_id, account_id, number, expiry_month, expiry_year, status
FROM cards`

	var cards []core.Card

	if err := r.db.SelectContext(ctx, &cards, query); err != nil {
		return nil, core.InternalServerError
	}

	return cards, nil
}

func (r *CardRepository) GetByUser(ctx context.Context, userId int64) ([]core.Card, error) {
	query := `
SELECT id, user_id, account_id, number, expiry_month, expiry_year, status
FROM cards
WHERE user_id = $1`

	var cards []core.Card

	if err := r.db.SelectContext(ctx, &cards, query, userId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	return cards, nil
}

func (r *CardRepository) GetById(ctx context.Context, id int64) (*core.Card, error) {
	query := `
SELECT id, user_id, account_id, number, expiry_month, expiry_year, status
FROM cards
WHERE id = $1`

	var card core.Card

	if err := r.db.GetContext(ctx, &card, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	return &card, nil
}

func (r *CardRepository) GetByNumber(ctx context.Context, number string) (*core.Card, error) {
	query := `
SELECT id, user_id, account_id, number, expiry_month, expiry_year, status
FROM cards
WHERE number = $1`

	var card core.Card

	if err := r.db.GetContext(ctx, &card, query, number); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	return &card, nil
}

func (r *CardRepository) Blocking(ctx context.Context, id int64) (core.Card, error) {
	query := `
	UPDATE cards
	SET status = $1, updated_at = $2
	WHERE id = $3
	RETURNING id, user_id, account_id, number, expiry_month, expiry_year, status
	`

	var card core.Card

	err := r.db.QueryRowContext(ctx, query, core.CardStatusBlocked, time.Now(), id).
		Scan(&card.Id, &card.UserId, &card.AccountId, &card.Number, &card.ExpiryMonth, &card.ExpiryYear, &card.Status)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return core.Card{}, core.NotFound
		}

		return core.Card{}, core.InternalServerError
	}

	return card, nil
}

func (r *CardRepository) Create(ctx context.Context, input *core.Card) (*core.Card, error) {
	query := `
INSERT INTO accounts (user_id, account_id, number, expiry_month, expiry_year, status)
VALUES (:user_id, :account_id, :number, :expiry_month, :expiry_year, :status)
RETURNING id, user_id, account_id, number, expiry_month, expiry_year, status;`

	rows, err := r.db.NamedQueryContext(ctx, query, input)
	if err != nil {
		return nil, core.InternalServerError
	}
	defer rows.Close()

	if rows.Next() {
		var created core.Card
		if err := rows.StructScan(&created); err != nil {
			return nil, core.InternalServerError
		}
		return &created, nil
	}

	return nil, core.InternalServerError
}

func (r *CardRepository) Delete(ctx context.Context, id int64) (int64, error) {
	query := `
	DELETE FROM cards
	WHERE id = $1
	RETURNING user_id
	`

	var userId int64

	err := r.db.QueryRowContext(ctx, query, id).Scan(&userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, core.NotFound
		}

		return 0, core.InternalServerError
	}

	return userId, nil
}
