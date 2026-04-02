package core

import (
	"context"
	"database/sql"
	"errors"

	core "github.com/Suinar/Bank-backend/bankBackend/internal/core"
	"github.com/jmoiron/sqlx"
)

type DepositRepository struct {
	db *sqlx.DB
}

func NewDepositRepository(db *sqlx.DB) *DepositRepository {
	return &DepositRepository{db: db}
}

func (r *DepositRepository) GetAll(ctx context.Context) ([]core.Deposit, error) {
	query := `
SELECT id, user_id, currency_id, amount, interest_rate, term_month, monthly_payment, status
FROM deposits`

	var deposits []core.Deposit

	if err := r.db.SelectContext(ctx, &deposits, query); err != nil {
		return nil, core.InternalServerError
	}

	return deposits, nil
}

func (r *DepositRepository) GetByUser(ctx context.Context, userId uint64) ([]core.Deposit, error) {
	query := `
SELECT id, user_id, currency_id, amount, interest_rate, term_month, monthly_payment, status
FROM deposits
WHERE user_id = $1`

	var deposits []core.Deposit

	if err := r.db.SelectContext(ctx, &deposits, query, userId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	return deposits, nil
}

func (r *DepositRepository) GetById(ctx context.Context, id uint64) (*core.Deposit, error) {
	query := `
SELECT id, user_id, currency_id, amount, interest_rate, term_month, monthly_payment, status
FROM deposits
WHERE id = $1`

	var deposit core.Deposit

	if err := r.db.GetContext(ctx, &deposit, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	return &deposit, nil
}

func (r *DepositRepository) Create(ctx context.Context, input *core.Deposit) (*core.Deposit, error) {
	query := `
INSERT INTO accounts (user_id, currency_id, amount, interest_rate, term_month, monthly_payment, status)
VALUES (:user_id, :currency_id, :amount, :interest_rate, :term_month, :monthly_payment, :status)
RETURNING user_id, currency_id, amount, interest_rate, term_month, monthly_payment, status;`

	rows, err := r.db.NamedQueryContext(ctx, query, input)
	if err != nil {
		return nil, core.InternalServerError
	}
	defer rows.Close()

	if rows.Next() {
		var created core.Deposit
		if err := rows.StructScan(&created); err != nil {
			return nil, core.InternalServerError
		}
		return &created, nil
	}

	return nil, core.InternalServerError
}

func (r *DepositRepository) Delete(ctx context.Context, id uint64) error {
	query := `
DELETE FROM deposits 
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
