package model

import "time"

type UserPoint struct {
	ID        int
	UserID    int
	Balance   int
	UpdatedAt time.Time
}

type PointTransaction struct {
	ID          int
	UserID      int
	Amount      int
	Type        TransactionType
	Description string
	CreatedAt   time.Time
}

type TransactionType string

const (
	TransactionTypeGacha TransactionType = "gacha"
	TransactionTypeSpend TransactionType = "spend"
)

func NewUserPoint(userID int) *UserPoint {
	return &UserPoint{
		UserID:    userID,
		Balance:   0,
		UpdatedAt: time.Now(),
	}
}

func (up *UserPoint) AddPoints(amount int) {
	up.Balance += amount
	up.UpdatedAt = time.Now()
}

func (up *UserPoint) SpendPoints(amount int) bool {
	if up.Balance >= amount {
		up.Balance -= amount
		up.UpdatedAt = time.Now()
		return true
	}
	return false
}