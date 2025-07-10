package repository

import (
	"context"
	"github.com/Amane-Fujiwara11/FortuneSpinner/backend/domain/model"
)

type PointRepository interface {
	GetUserPoint(ctx context.Context, userID int) (*model.UserPoint, error)
	CreateUserPoint(ctx context.Context, userPoint *model.UserPoint) error
	UpdateUserPoint(ctx context.Context, userPoint *model.UserPoint) error
	SaveTransaction(ctx context.Context, transaction *model.PointTransaction) error
	FindTransactionsByUserID(ctx context.Context, userID int, limit int) ([]*model.PointTransaction, error)
}