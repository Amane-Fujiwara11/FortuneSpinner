package model

import (
	"errors"
	"time"
)

const (
	// Point system business rules
	MaxPointBalance     = 1000000  // Maximum points a user can hold
	MinTransactionAmount = 1       // Minimum transaction amount
	MaxTransactionAmount = 10000   // Maximum single transaction amount
)

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

// NewUserPoint creates a new UserPoint with validation
func NewUserPoint(userID int) (*UserPoint, error) {
	if userID <= 0 {
		return nil, errors.New("user ID must be positive")
	}
	
	return &UserPoint{
		UserID:    userID,
		Balance:   0,
		UpdatedAt: time.Now(),
	}, nil
}

// AddPoints adds points with business rule validation
func (up *UserPoint) AddPoints(amount int) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}
	
	if amount < MinTransactionAmount {
		return errors.New("amount is below minimum transaction amount")
	}
	
	if amount > MaxTransactionAmount {
		return errors.New("amount exceeds maximum transaction amount")
	}
	
	newBalance := up.Balance + amount
	if newBalance > MaxPointBalance {
		return errors.New("transaction would exceed maximum point balance")
	}
	
	up.Balance = newBalance
	up.UpdatedAt = time.Now()
	return nil
}

// SpendPoints spends points with business rule validation
func (up *UserPoint) SpendPoints(amount int) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}
	
	if amount < MinTransactionAmount {
		return errors.New("amount is below minimum transaction amount")
	}
	
	if amount > MaxTransactionAmount {
		return errors.New("amount exceeds maximum transaction amount")
	}
	
	if up.Balance < amount {
		return errors.New("insufficient points")
	}
	
	up.Balance -= amount
	up.UpdatedAt = time.Now()
	return nil
}

// CanAfford checks if the user can afford a specific amount
func (up *UserPoint) CanAfford(amount int) bool {
	return up.Balance >= amount && amount > 0
}

// GetPointLevel returns the user's point level based on total balance
func (up *UserPoint) GetPointLevel() string {
	switch {
	case up.Balance >= 10000:
		return "Diamond"
	case up.Balance >= 5000:
		return "Gold"
	case up.Balance >= 1000:
		return "Silver"
	default:
		return "Bronze"
	}
}

// NewPointTransaction creates a new point transaction with validation
func NewPointTransaction(userID int, amount int, transactionType TransactionType, description string) (*PointTransaction, error) {
	if userID <= 0 {
		return nil, errors.New("user ID must be positive")
	}
	
	if amount <= 0 {
		return nil, errors.New("transaction amount must be positive")
	}
	
	if amount < MinTransactionAmount || amount > MaxTransactionAmount {
		return nil, errors.New("transaction amount is outside allowed range")
	}
	
	if transactionType != TransactionTypeGacha && transactionType != TransactionTypeSpend {
		return nil, errors.New("invalid transaction type")
	}
	
	if description == "" {
		return nil, errors.New("transaction description cannot be empty")
	}
	
	return &PointTransaction{
		UserID:      userID,
		Amount:      amount,
		Type:        transactionType,
		Description: description,
		CreatedAt:   time.Now(),
	}, nil
}