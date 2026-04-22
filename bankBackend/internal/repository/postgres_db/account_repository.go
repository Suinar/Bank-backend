package core

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	core "github.com/Suinar/Bank-backend/bankBackend/internal/core"
	"github.com/jmoiron/sqlx"
)

type AccountRepository struct {
	db *sqlx.DB
}

func NewAccountRepository(db *sqlx.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (r *AccountRepository) GetAll(ctx context.Context) ([]core.Account, error) {
	query := `
SELECT id, user_id, currency_id, name, balance, status
FROM accounts
    `

	var accounts []core.Account

	if err := r.db.SelectContext(ctx, &accounts, query); err != nil {
		return nil, core.InternalServerError
	}

	return accounts, nil
}

func (r *AccountRepository) GetByUser(ctx context.Context, userId int64) ([]core.Account, error) {
	query := `
SELECT id, user_id, currency_id, name, balance, status
FROM accounts
WHERE user_id = $1`

	var accounts []core.Account

	if err := r.db.SelectContext(ctx, &accounts, query, userId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	return accounts, nil
}

func (r *AccountRepository) GetById(ctx context.Context, id int64) (*core.Account, error) {
	query := `
SELECT id, user_id, currency_id, name, balance, status
FROM accounts
WHERE id = $1`

	var account core.Account

	if err := r.db.GetContext(ctx, &account, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	return &account, nil
}

func (r *AccountRepository) Create(ctx context.Context, input *core.Account) (*core.Account, error) {
	query := `
INSERT INTO accounts (user_id, currency_id, name, balance, status)
VALUES (:user_id, :currency_id, :name, :balance, :status)
RETURNING id, user_id, currency_id, name, balance, status;`

	rows, err := r.db.NamedQueryContext(ctx, query, input)
	if err != nil {
		return nil, core.InternalServerError
	}
	defer rows.Close()

	if rows.Next() {
		var created core.Account
		if err := rows.StructScan(&created); err != nil {
			return nil, core.InternalServerError
		}
		return &created, nil
	}

	return nil, core.InternalServerError
}

func (r *AccountRepository) Blocking(ctx context.Context, id int64) error {
	query := `
UPDATE accounts
SET status = $1, updated_at = $2
where id = $3`

	result, err := r.db.ExecContext(ctx, query, core.AccountStatusBlocked, time.Now(), id)
	if err != nil {
		return core.InternalServerError
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return core.InternalServerError
	}

	if rows == 0 {
		return core.NotFound
	}

	return nil
}

func (r *AccountRepository) Close(ctx context.Context, id int64) error {
	query := `
UPDATE accounts
SET status = $1, updated_at = $2
where id = $3`

	result, err := r.db.ExecContext(ctx, query, core.AccountStatusClosed, time.Now(), id)
	if err != nil {
		return core.InternalServerError
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return core.InternalServerError
	}

	if rows == 0 {
		return core.NotFound
	}

	return nil
}

func (r *AccountRepository) Update(ctx context.Context, id int64, input *core.AccountUpdateInput) (*core.Account, error) {
	setParts := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setParts = append(setParts, fmt.Sprintf("name = $%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	setParts = append(setParts, "updated_at = NOW()")

	args = append(args, id)
	args = append(args, core.AccountStatusClosed)

	if len(setParts) == 0 {
		return nil, core.BadRequest
	}

	query := fmt.Sprintf(`
UPDATE accounts
SET %s
WHERE id = $%d AND status != $%d
RETURNING id, user_id, currency_id, name, balance, status
`, strings.Join(setParts, ", "), argId, argId+1)

	var updated core.Account

	err := r.db.QueryRowxContext(ctx, query, args...).StructScan(&updated)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	return &updated, nil
}

func (r *AccountRepository) Delete(ctx context.Context, id int64) error {
	query := `
DELETE FROM accounts 
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
