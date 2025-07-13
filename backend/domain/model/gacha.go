package model

import (
	"errors"
	"math"
)

type Rarity int

const (
	RarityCommon Rarity = iota + 1
	RarityRare
	RarityEpic
	RarityLegendary
)

type GachaItem struct {
	ID          int
	Name        string
	Rarity      Rarity
	Points      int
	Probability float64
}

func (r Rarity) String() string {
	switch r {
	case RarityCommon:
		return "Common"
	case RarityRare:
		return "Rare"
	case RarityEpic:
		return "Epic"
	case RarityLegendary:
		return "Legendary"
	default:
		return "Unknown"
	}
}

// IsValid checks if the rarity is valid
func (r Rarity) IsValid() bool {
	return r >= RarityCommon && r <= RarityLegendary
}

// GetMinPoints returns the minimum points for this rarity level
func (r Rarity) GetMinPoints() int {
	switch r {
	case RarityCommon:
		return 1
	case RarityRare:
		return 10
	case RarityEpic:
		return 100
	case RarityLegendary:
		return 500
	default:
		return 0
	}
}

// GetMaxPoints returns the maximum points for this rarity level
func (r Rarity) GetMaxPoints() int {
	switch r {
	case RarityCommon:
		return 50
	case RarityRare:
		return 200
	case RarityEpic:
		return 1000
	case RarityLegendary:
		return 5000
	default:
		return 0
	}
}

// NewGachaItem creates a new gacha item with validation
func NewGachaItem(id int, name string, rarity Rarity, points int, probability float64) (*GachaItem, error) {
	item := &GachaItem{
		ID:          id,
		Name:        name,
		Rarity:      rarity,
		Points:      points,
		Probability: probability,
	}
	
	if err := item.Validate(); err != nil {
		return nil, err
	}
	
	return item, nil
}

// Validate validates the gacha item according to business rules
func (gi *GachaItem) Validate() error {
	if gi.ID <= 0 {
		return errors.New("gacha item ID must be positive")
	}
	
	if gi.Name == "" {
		return errors.New("gacha item name cannot be empty")
	}
	
	if !gi.Rarity.IsValid() {
		return errors.New("invalid rarity level")
	}
	
	minPoints := gi.Rarity.GetMinPoints()
	maxPoints := gi.Rarity.GetMaxPoints()
	
	if gi.Points < minPoints || gi.Points > maxPoints {
		return errors.New("points value doesn't match rarity constraints")
	}
	
	if gi.Probability < 0 || gi.Probability > 1 {
		return errors.New("probability must be between 0 and 1")
	}
	
	return nil
}

// GetGachaItems returns the configured gacha items with business rule validation
func GetGachaItems() ([]GachaItem, error) {
	items := []GachaItem{
		{ID: 1, Name: "Bronze Coin", Rarity: RarityCommon, Points: 10, Probability: 0.60},
		{ID: 2, Name: "Silver Coin", Rarity: RarityRare, Points: 50, Probability: 0.30},
		{ID: 3, Name: "Gold Coin", Rarity: RarityEpic, Points: 200, Probability: 0.08},
		{ID: 4, Name: "Diamond", Rarity: RarityLegendary, Points: 1000, Probability: 0.02},
	}
	
	// Validate all items
	for _, item := range items {
		if err := item.Validate(); err != nil {
			return nil, err
		}
	}
	
	// Validate that probabilities sum to approximately 1.0
	totalProbability := 0.0
	for _, item := range items {
		totalProbability += item.Probability
	}
	
	if math.Abs(totalProbability-1.0) > 0.001 {
		return nil, errors.New("gacha item probabilities must sum to 1.0")
	}
	
	return items, nil
}

// GetItemByRarity returns all items of a specific rarity
func GetItemsByRarity(rarity Rarity) ([]GachaItem, error) {
	allItems, err := GetGachaItems()
	if err != nil {
		return nil, err
	}
	
	var filteredItems []GachaItem
	for _, item := range allItems {
		if item.Rarity == rarity {
			filteredItems = append(filteredItems, item)
		}
	}
	
	return filteredItems, nil
}