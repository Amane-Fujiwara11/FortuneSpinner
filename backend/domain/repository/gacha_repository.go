package repository

import (
	"context"
	"github.com/Amane-Fujiwara11/FortuneSpinner/backend/domain/model"
)

type GachaRepository interface {
	SaveResult(ctx context.Context, result *model.GachaResult) error
	FindResultsByUserID(ctx context.Context, userID int, limit int) ([]*model.GachaResult, error)
	FindResultByID(ctx context.Context, id int) (*model.GachaResult, error)
}