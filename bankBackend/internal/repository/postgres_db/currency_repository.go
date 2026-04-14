package core

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	core "github.com/Suinar/Bank-backend/bankBackend/internal/core"
	"github.com/jmoiron/sqlx"
)

type CurrencyRepository struct {
	db *sqlx.DB
}

func (r *CurrencyRepository) SetAll(ctx context.context.Context, currencies []core.Currency) error {
	//TODO implement me
	panic("implement me")
}

func (r *CurrencyRepository) Set(ctx context.context.Context, currency *core.Currency) error {
	//TODO implement me
	panic("implement me")
}

func NewCurrencyRepository(db *sqlx.DB) *CurrencyRepository {
	return &CurrencyRepository{db: db}
}

func (r *CurrencyRepository) GetAll(ctx context.Context) ([]core.Currency, error) {
	query := `
SELECT id, name, symbol, iso_code, minor_units
FROM currencies`

	var currencies []core.Currency

	if err := r.db.SelectContext(ctx, &currencies, query); err != nil {
		return nil, core.InternalServerError
	}

	return currencies, nil
}

func (r *CurrencyRepository) GetById(ctx context.Context, id uint64) (*core.Currency, error) {
	query := `
SELECT id, name, symbol, iso_code, minor_units
FROM currencies
WHERE id = $1`

	var currency core.Currency

	if err := r.db.GetContext(ctx, &currency, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	return &currency, nil
}

func (r *CurrencyRepository) GetByIso(ctx context.Context, isoCode string) (*core.Currency, error) {
	query := `
SELECT id, name, symbol, iso_code, minor_units
FROM currencies
WHERE iso_code = $1`

	var currency core.Currency

	if err := r.db.GetContext(ctx, &currency, query, isoCode); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	return &currency, nil
}

func (r *CurrencyRepository) GetBySymbol(ctx context.Context, symbol rune) (*core.Currency, error) {
	query := `
SELECT id, name, symbol, iso_code, minor_units
FROM currencies
WHERE symbol = $1`

	var currency core.Currency

	if err := r.db.GetContext(ctx, &currency, query, symbol); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	return &currency, nil
}

func (r *CurrencyRepository) Create(ctx context.Context, input *core.Currency) (*core.Currency, error) {
	query := `
INSERT INTO currencies (name, symbol, iso_code, minor_units)
VALUES (:name, :symbol, :iso_code, :minor_units)
RETURNING id, name, symbol, iso_code, minor_units`

	rows, err := r.db.NamedQueryContext(ctx, query, input)
	if err != nil {
		return nil, core.InternalServerError
	}
	defer rows.Close()
	if rows.Next() {
		var created core.Currency
		if err := rows.StructScan(&created); err != nil {
			return nil, core.InternalServerError
		}
		return &created, nil
	}

	return nil, core.InternalServerError
}

func (r *CurrencyRepository) Update(ctx context.Context, id uint64, input *core.CurrencyUpdateInput) (*core.Currency, error) {
	setParts := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setParts = append(setParts, fmt.Sprintf("name = $%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.Symbol != nil {
		setParts = append(setParts, fmt.Sprintf("symbol = $%d", argId))
		args = append(args, *input.Symbol)
		argId++
	}

	if input.IsoCode != nil {
		setParts = append(setParts, fmt.Sprintf("iso_code = $%d", argId))
		args = append(args, *input.IsoCode)
		argId++
	}

	if input.MinorUnits != nil {
		setParts = append(setParts, fmt.Sprintf("minor_units = $%d", argId))
		args = append(args, *input.MinorUnits)
		argId++
	}

	setParts = append(setParts, "updated_at = NOW()")

	if len(setParts) == 0 {
		return nil, core.BadRequest
	}

	query := fmt.Sprintf(`
UPDATE accounts
SET %s
WHERE id = $%d
RETURNING id, name, symbol, iso_code, minor_units
`, strings.Join(setParts, ", "), argId, argId+1)

	var updated core.Currency

	err := r.db.QueryRowxContext(ctx, query, args...).StructScan(&updated)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	return &updated, nil
}

func (r *CurrencyRepository) Delete(ctx context.Context, id uint64) error {
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
