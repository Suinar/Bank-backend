package core

import (
	"context"
	"database/sql"
	"errors"

	core "github.com/Suinar/Bank-backend/bankBackend/internal/core"
	"github.com/jmoiron/sqlx"
)

type CreditRepository struct {
	db *sqlx.DB
}

func NewCreditRepository(db *sqlx.DB) *CreditRepository {
	return &CreditRepository{db: db}
}

func (r *CreditRepository) GetAll(ctx context.Context) ([]core.Credit, error) {
	query := `
SELECT id, user_id, currency_id, amount, interest_rate, term_month, monthly_payment, status
FROM credits`

	var credits []core.Credit

	if err := r.db.SelectContext(ctx, &credits, query); err != nil {
		return nil, core.InternalServerError
	}

	return credits, nil
}

func (r *CreditRepository) GetByUser(ctx context.Context, userId int64) ([]core.Credit, error) {
	query := `
SELECT id, user_id, currency_id, amount, interest_rate, term_month, monthly_payment, status
FROM credits
WHERE user_id = $1`

	var credits []core.Credit

	if err := r.db.SelectContext(ctx, &credits, query, userId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	return credits, nil
}

func (r *CreditRepository) GetById(ctx context.Context, id int64) (*core.Credit, error) {
	query := `
SELECT id, user_id, currency_id, amount, interest_rate, term_month, monthly_payment, status
FROM credits
WHERE id = $1`

	var credits core.Credit

	if err := r.db.GetContext(ctx, &credits, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	return &credits, nil
}

func (r *CreditRepository) Create(ctx context.Context, input *core.Credit) (*core.Credit, error) {
	query := `
INSERT INTO credits (user_id, currency_id, amount, interest_rate, term_month, monthly_payment, status)
VALUES (:user_id, :currency_id, :amount, :interest_rate, :term_month, :monthly_payment, :status)
Returning id, user_id, currency_id, amount, interest_rate, term_month, monthly_payment, status;`

	rows, err := r.db.NamedQueryContext(ctx, query, input)
	if err != nil {
		return nil, core.InternalServerError
	}
	defer rows.Close()

	if rows.Next() {
		var created core.Credit
		if err := rows.StructScan(&created); err != nil {
			return nil, core.InternalServerError
		}
		return &created, nil
	}

	return nil, core.InternalServerError
}

func (r *CreditRepository) Delete(ctx context.Context, id int64) error {
	query := `
DELETE FROM credits 
WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return core.InternalServerError
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return core.InternalServerError
	}

	if rowsAffected == 0 {
		return core.NotFound
	}

	return nil
}
