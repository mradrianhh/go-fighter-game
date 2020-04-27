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

// Items holds a collection of items.
type Items map[string]Item

// DamageItems holds a collection of all damage items in the game.
var DamageItems Items = make(Items)

// ArmorItems holds a collection of all armor items in the game.
var ArmorItems Items = make(Items)

// HealthItems holds a collection of all health items in the game.
var HealthItems Items = make(Items)

// AuxiliaryItems holds a collection of all auxiliary items in the game.
var AuxiliaryItems Items = make(Items)

func init() {
	DamageItems["sword"] = Item{Name: "Sword", Cost: 10, DamageIncrease: 10, HealthIncrease: 0, ArmorIncrease: 0}
	ArmorItems["chainvest"] = Item{Name: "Chainvest", Cost: 10, DamageIncrease: 0, HealthIncrease: 0, ArmorIncrease: 10}
	HealthItems["lifefountain"] = Item{Name: "Life Fountain", Cost: 10, DamageIncrease: 0, HealthIncrease: 10, ArmorIncrease: 10}
	AuxiliaryItems["healthpot"] = Item{Name: "Health Pot", Cost: 1, DamageIncrease: 0, HealthIncrease: 20, ArmorIncrease: 0, Description: "+20 Health"}
}

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
