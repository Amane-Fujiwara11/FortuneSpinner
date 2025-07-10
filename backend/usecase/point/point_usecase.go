package point

import (
	"context"
	"errors"

	"github.com/Amane-Fujiwara11/FortuneSpinner/backend/domain/model"
	"github.com/Amane-Fujiwara11/FortuneSpinner/backend/domain/repository"
)

type PointUsecase interface {
	GetBalance(ctx context.Context, userID int) (int, error)
	GetTransactionHistory(ctx context.Context, userID int, limit int) ([]*model.PointTransaction, error)
}

type pointUsecase struct {
	pointRepo repository.PointRepository
	userRepo  repository.UserRepository
}

func NewPointUsecase(
	pointRepo repository.PointRepository,
	userRepo repository.UserRepository,
) PointUsecase {
	return &pointUsecase{
		pointRepo: pointRepo,
		userRepo:  userRepo,
	}
}

func (uc *pointUsecase) GetBalance(ctx context.Context, userID int) (int, error) {
	// ユーザーの存在確認
	user, err := uc.userRepo.FindByID(ctx, userID)
	if err != nil {
		return 0, err
	}
	if user == nil {
		return 0, errors.New("user not found")
	}

	// ポイント残高を取得
	userPoint, err := uc.pointRepo.GetUserPoint(ctx, userID)
	if err != nil {
		return 0, err
	}

	if userPoint == nil {
		// ポイントデータがない場合は0を返す
		return 0, nil
	}

	return userPoint.Balance, nil
}

func (uc *pointUsecase) GetTransactionHistory(ctx context.Context, userID int, limit int) ([]*model.PointTransaction, error) {
	// ユーザーの存在確認
	user, err := uc.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	return uc.pointRepo.FindTransactionsByUserID(ctx, userID, limit)
}