package models

// Shop is a place where the player can buy items.
type Shop struct {
	Items []*Item
}

// NewShop initializes and returns a new shop.
func NewShop() Shop {
	var items []*Item
	sword := Item{Name: "Sword", Cost: 10, DamageIncrease: 10, HealthIncrease: 0}
	healthPot := Item{Name: "Health Pot", Cost: 1, DamageIncrease: 0, HealthIncrease: 20}

	items = append(items, &sword, &healthPot)

	shop := Shop{Items: items}

	return shop
}
