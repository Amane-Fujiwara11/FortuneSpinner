package model

import "time"

type GachaResult struct {
	ID           int
	UserID       int
	ItemID       int
	ItemName     string
	Rarity       Rarity
	PointsEarned int
	CreatedAt    time.Time
}

func NewGachaResult(userID int, item GachaItem) *GachaResult {
	return &GachaResult{
		UserID:       userID,
		ItemID:       item.ID,
		ItemName:     item.Name,
		Rarity:       item.Rarity,
		PointsEarned: item.Points,
		CreatedAt:    time.Now(),
	}
}