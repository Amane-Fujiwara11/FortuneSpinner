package model

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

func GetGachaItems() []GachaItem {
	return []GachaItem{
		{ID: 1, Name: "Bronze Coin", Rarity: RarityCommon, Points: 10, Probability: 0.60},
		{ID: 2, Name: "Silver Coin", Rarity: RarityRare, Points: 50, Probability: 0.30},
		{ID: 3, Name: "Gold Coin", Rarity: RarityEpic, Points: 200, Probability: 0.08},
		{ID: 4, Name: "Diamond", Rarity: RarityLegendary, Points: 1000, Probability: 0.02},
	}
}