package repository

import (
	"context"
	"database/sql"

	"github.com/Amane-Fujiwara11/FortuneSpinner/backend/domain/model"
	"github.com/Amane-Fujiwara11/FortuneSpinner/backend/domain/repository"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	query := `INSERT INTO users (name, created_at, updated_at) VALUES (?, ?, ?)`
	result, err := r.db.ExecContext(ctx, query, user.Name, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = int(id)
	return nil
}

func (r *userRepository) FindByID(ctx context.Context, id int) (*model.User, error) {
	query := `SELECT id, name, created_at, updated_at FROM users WHERE id = ?`
	var user model.User
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Name,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(ctx context.Context, user *model.User) error {
	query := `UPDATE users SET name = ?, updated_at = ? WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, user.Name, user.UpdatedAt, user.ID)
	return err
}