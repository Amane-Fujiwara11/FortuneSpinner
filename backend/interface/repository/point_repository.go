package repository

import (
	"context"
	"database/sql"

	"github.com/Amane-Fujiwara11/FortuneSpinner/backend/domain/model"
	"github.com/Amane-Fujiwara11/FortuneSpinner/backend/domain/repository"
)

type pointRepository struct {
	db *sql.DB
}

func NewPointRepository(db *sql.DB) repository.PointRepository {
	return &pointRepository{
		db: db,
	}
}

func (r *pointRepository) GetUserPoint(ctx context.Context, userID int) (*model.UserPoint, error) {
	query := `SELECT id, user_id, balance, updated_at FROM user_points WHERE user_id = ?`
	var userPoint model.UserPoint
	err := r.db.QueryRowContext(ctx, query, userID).Scan(
		&userPoint.ID,
		&userPoint.UserID,
		&userPoint.Balance,
		&userPoint.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &userPoint, nil
}

func (r *pointRepository) CreateUserPoint(ctx context.Context, userPoint *model.UserPoint) error {
	query := `INSERT INTO user_points (user_id, balance, updated_at) VALUES (?, ?, ?)`
	result, err := r.db.ExecContext(ctx, query, userPoint.UserID, userPoint.Balance, userPoint.UpdatedAt)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	userPoint.ID = int(id)
	return nil
}

func (r *pointRepository) UpdateUserPoint(ctx context.Context, userPoint *model.UserPoint) error {
	query := `UPDATE user_points SET balance = ?, updated_at = ? WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, userPoint.Balance, userPoint.UpdatedAt, userPoint.ID)
	return err
}

func (r *pointRepository) SaveTransaction(ctx context.Context, transaction *model.PointTransaction) error {
	query := `INSERT INTO point_transactions (user_id, amount, type, description, created_at) VALUES (?, ?, ?, ?, ?)`
	result, err := r.db.ExecContext(ctx, query,
		transaction.UserID,
		transaction.Amount,
		transaction.Type,
		transaction.Description,
		transaction.CreatedAt,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	transaction.ID = int(id)
	return nil
}

func (r *pointRepository) FindTransactionsByUserID(ctx context.Context, userID int, limit int) ([]*model.PointTransaction, error) {
	query := `SELECT id, user_id, amount, type, description, created_at 
		FROM point_transactions 
		WHERE user_id = ? 
		ORDER BY created_at DESC 
		LIMIT ?`

	rows, err := r.db.QueryContext(ctx, query, userID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []*model.PointTransaction
	for rows.Next() {
		var tx model.PointTransaction
		err := rows.Scan(
			&tx.ID,
			&tx.UserID,
			&tx.Amount,
			&tx.Type,
			&tx.Description,
			&tx.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, &tx)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}