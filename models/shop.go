package models

// Shop is a place where the player can buy items.
type Shop struct {
	Inventory *Inventory
}

// NewShop initializes and returns a new shop.
func NewShop() Shop {
	var inventory Inventory
	sword := Item{Name: "Sword", Cost: 10, DamageIncrease: 10, HealthIncrease: 0, ArmorIncrease: 0}
	chainVest := Item{Name: "Chainvest", Cost: 10, DamageIncrease: 0, HealthIncrease: 0, ArmorIncrease: 10}
	lifeFountain := Item{Name: "Life Fountain", Cost: 10, DamageIncrease: 0, HealthIncrease: 10, ArmorIncrease: 10}
	healthPot := Item{Name: "Health Pot", Cost: 1, DamageIncrease: 0, HealthIncrease: 20, ArmorIncrease: 0, Description: "+20 Health"}

	inventory = Inventory{sword, chainVest, lifeFountain, healthPot}

	shop := Shop{Inventory: &inventory}

	return shop
}
