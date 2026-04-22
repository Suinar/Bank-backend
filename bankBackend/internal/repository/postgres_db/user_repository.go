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

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAll(ctx context.Context) ([]core.User, error) {
	query := `
SELECT id, user_id, middle_name, last_name, email, phone_number, password_hash
FROM users`

	var users []core.User

	if err := r.db.SelectContext(ctx, &users, query); err != nil {
		return nil, core.InternalServerError
	}

	return users, nil
}

func (r *UserRepository) GetById(ctx context.Context, id int64) (*core.User, error) {
	query := `
SELECT id, user_id, middle_name, last_name, email, phone_number, password_hash
FROM users
WHERE id = $1`

	var user core.User

	if err := r.db.GetContext(ctx, &user, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	return &user, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*core.User, error) {
	query := `
SELECT id, user_id, middle_name, last_name, email, phone_number, password_hash
FROM users
WHERE email = $1`

	var user core.User

	if err := r.db.GetContext(ctx, &user, query, email); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	return &user, nil
}

func (r *UserRepository) GetByPhoneNumber(ctx context.Context, phoneNumber string) (*core.User, error) {
	query := `
SELECT id, user_id, middle_name, last_name, email, phone_number, password_hash
FROM users
WHERE phone_number = $1`

	var user core.User

	if err := r.db.GetContext(ctx, &user, query, phoneNumber); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, core.NotFound
		}

		return nil, core.InternalServerError
	}

	return &user, nil
}

func (r *UserRepository) Create(ctx context.Context, input *core.User) (*core.User, error) {
	query := `
INSERT INTO users (user_id, middle_name, last_name, email, phone_number, password_hash)
VALUES (:user_id, :middle_name, :last_name, :email, :phone_number, :password_hash)
RETURNING id, user_id, middle_name, last_name, email, phone_number, password_hash;`

	rows, err := r.db.NamedQueryContext(ctx, query, input)
	if err != nil {
		return nil, core.InternalServerError
	}
	defer rows.Close()

	if rows.Next() {
		var created core.User
		if err := rows.StructScan(&created); err != nil {
			return nil, core.InternalServerError
		}
		return &created, nil
	}

	return nil, core.InternalServerError
}

func (r *UserRepository) ChangePassword(ctx context.Context, id int64, newPassword string) error {
	query := `
UPDATE users
SET password_hash = $1, updated_at = $2
where id = $3`

	result, err := r.db.ExecContext(ctx, query, newPassword, time.Now(), id)
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

func (r *UserRepository) Update(ctx context.Context, id int64, input *core.UserUpdateInput) (*core.User, error) {
	setParts := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	// Динамічно формуємо SET для кожного поля
	if input.FirstName != nil {
		setParts = append(setParts, fmt.Sprintf("first_name = $%d", argId))
		args = append(args, *input.FirstName)
		argId++
	}
	if input.MiddleName != nil {
		setParts = append(setParts, fmt.Sprintf("middle_name = $%d", argId))
		args = append(args, *input.MiddleName)
		argId++
	}
	if input.LastName != nil {
		setParts = append(setParts, fmt.Sprintf("last_name = $%d", argId))
		args = append(args, *input.LastName)
		argId++
	}

	if len(setParts) == 0 {
		return nil, core.BadRequest
	}

	setParts = append(setParts, "updated_at = NOW()")

	args = append(args, id)

	query := fmt.Sprintf(`
		UPDATE users
		SET %s
		WHERE id = $%d
		RETURNING id, first_name, middle_name, last_name, email, created_at, updated_at
	`, strings.Join(setParts, ", "), argId)

	var updated core.User

	err := r.db.QueryRowxContext(ctx, query, args...).StructScan(&updated)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, core.NotFound
		}
		return nil, core.InternalServerError
	}

	return &updated, nil
}

func (r *UserRepository) Delete(ctx context.Context, id int64) error {
	query := `
DELETE FROM users 
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
