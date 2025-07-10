package repository

import (
	"context"
	"database/sql"

	"github.com/Amane-Fujiwara11/FortuneSpinner/backend/domain/model"
	"github.com/Amane-Fujiwara11/FortuneSpinner/backend/domain/repository"
)

type gachaRepository struct {
	db *sql.DB
}

func NewGachaRepository(db *sql.DB) repository.GachaRepository {
	return &gachaRepository{
		db: db,
	}
}

func (r *gachaRepository) SaveResult(ctx context.Context, result *model.GachaResult) error {
	query := `INSERT INTO gacha_results (user_id, item_id, item_name, rarity, points_earned, created_at) VALUES (?, ?, ?, ?, ?, ?)`
	res, err := r.db.ExecContext(ctx, query,
		result.UserID,
		result.ItemID,
		result.ItemName,
		result.Rarity,
		result.PointsEarned,
		result.CreatedAt,
	)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	result.ID = int(id)
	return nil
}

func (r *gachaRepository) FindResultsByUserID(ctx context.Context, userID int, limit int) ([]*model.GachaResult, error) {
	query := `SELECT id, user_id, item_id, item_name, rarity, points_earned, created_at 
		FROM gacha_results 
		WHERE user_id = ? 
		ORDER BY created_at DESC 
		LIMIT ?`

	rows, err := r.db.QueryContext(ctx, query, userID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*model.GachaResult
	for rows.Next() {
		var result model.GachaResult
		err := rows.Scan(
			&result.ID,
			&result.UserID,
			&result.ItemID,
			&result.ItemName,
			&result.Rarity,
			&result.PointsEarned,
			&result.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, &result)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *gachaRepository) FindResultByID(ctx context.Context, id int) (*model.GachaResult, error) {
	query := `SELECT id, user_id, item_id, item_name, rarity, points_earned, created_at 
		FROM gacha_results 
		WHERE id = ?`

	var result model.GachaResult
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&result.ID,
		&result.UserID,
		&result.ItemID,
		&result.ItemName,
		&result.Rarity,
		&result.PointsEarned,
		&result.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}