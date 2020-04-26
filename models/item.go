package models

// Item can be bought by the player to increase his stats.
// They can be bought in exchange for gold.
type Item struct {
	Name           string
	Cost           int
	DamageIncrease int
	HealthIncrease int
	ArmorIncrease  int
	Description    string
}

// Inventory holds a collection of 4 items.
type Inventory [4]Item

// DamageItem is a placeholder for an actual damage item.
var DamageItem = Item{
	Name:           "Damage Item",
	Cost:           0,
	DamageIncrease: 0,
	HealthIncrease: 0,
	ArmorIncrease:  0,
}

// HealthItem is a placeholder for an actual health item.
var HealthItem = Item{
	Name:           "Health Item",
	Cost:           0,
	DamageIncrease: 0,
	HealthIncrease: 0,
	ArmorIncrease:  0,
}

// ArmorItem is a placeholder for an actual armor item.
var ArmorItem = Item{
	Name:           "Armor Item",
	Cost:           0,
	DamageIncrease: 0,
	HealthIncrease: 0,
	ArmorIncrease:  0,
}

// AuxiliaryItem is a placeholder for an actual auxiliary item.
var AuxiliaryItem = Item{
	Name:           "Auxiliary Item",
	Cost:           0,
	DamageIncrease: 0,
	HealthIncrease: 0,
	ArmorIncrease:  0,
}
