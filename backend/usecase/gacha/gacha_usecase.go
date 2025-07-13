package gacha

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"github.com/Amane-Fujiwara11/FortuneSpinner/backend/domain/model"
	"github.com/Amane-Fujiwara11/FortuneSpinner/backend/domain/repository"
)

type GachaUsecase interface {
	ExecuteGacha(ctx context.Context, userID int) (*model.GachaResult, error)
	GetGachaHistory(ctx context.Context, userID int, limit int) ([]*model.GachaResult, error)
}

type gachaUsecase struct {
	gachaRepo repository.GachaRepository
	pointRepo repository.PointRepository
	userRepo  repository.UserRepository
}

func NewGachaUsecase(
	gachaRepo repository.GachaRepository,
	pointRepo repository.PointRepository,
	userRepo repository.UserRepository,
) GachaUsecase {
	return &gachaUsecase{
		gachaRepo: gachaRepo,
		pointRepo: pointRepo,
		userRepo:  userRepo,
	}
}

func (uc *gachaUsecase) ExecuteGacha(ctx context.Context, userID int) (*model.GachaResult, error) {
	// ユーザーの存在確認
	user, err := uc.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	// ガチャアイテムの抽選
	item, err := uc.drawGachaItem()
	if err != nil {
		return nil, err
	}

	// ガチャ結果を保存
	result := model.NewGachaResult(userID, item)
	if err := uc.gachaRepo.SaveResult(ctx, result); err != nil {
		return nil, err
	}

	// ポイント付与
	userPoint, err := uc.pointRepo.GetUserPoint(ctx, userID)
	if err != nil {
		return nil, err
	}

	if userPoint == nil {
		// 初回の場合は新規作成
		userPoint, err = model.NewUserPoint(userID)
		if err != nil {
			return nil, err
		}
		if err := userPoint.AddPoints(item.Points); err != nil {
			return nil, err
		}
		if err := uc.pointRepo.CreateUserPoint(ctx, userPoint); err != nil {
			return nil, err
		}
	} else {
		// 既存の場合は更新
		if err := userPoint.AddPoints(item.Points); err != nil {
			return nil, err
		}
		if err := uc.pointRepo.UpdateUserPoint(ctx, userPoint); err != nil {
			return nil, err
		}
	}

	// ポイント取引履歴を保存
	transaction, err := model.NewPointTransaction(userID, item.Points, model.TransactionTypeGacha, "Gacha reward: "+item.Name)
	if err != nil {
		return nil, err
	}
	if err := uc.pointRepo.SaveTransaction(ctx, transaction); err != nil {
		return nil, err
	}

	return result, nil
}

func (uc *gachaUsecase) GetGachaHistory(ctx context.Context, userID int, limit int) ([]*model.GachaResult, error) {
	return uc.gachaRepo.FindResultsByUserID(ctx, userID, limit)
}

func (uc *gachaUsecase) drawGachaItem() (model.GachaItem, error) {
	items, err := model.GetGachaItems()
	if err != nil {
		return model.GachaItem{}, err
	}
	
	// 確率に基づいてアイテムを抽選
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	roll := r.Float64()
	
	cumulative := 0.0
	for _, item := range items {
		cumulative += item.Probability
		if roll < cumulative {
			return item, nil
		}
	}
	
	// フォールバック（通常は到達しない）
	if len(items) > 0 {
		return items[0], nil
	}
	
	return model.GachaItem{}, errors.New("no gacha items available")
}