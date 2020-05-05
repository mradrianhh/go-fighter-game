package models

import "github.com/mradrianhh/go-fighter-game/persist"

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

// AddDamageItem adds an item to the "DamageItems"-map while encapsulating some additional logic.
func AddDamageItem(key string, item Item) {
	DamageItems[key] = item
	persist.Save("damageitems", DamageItems)
}

// ArmorItems holds a collection of all armor items in the game.
var ArmorItems Items = make(Items)

// AddArmorItem adds an item to the "ArmorItems"-map while encapsulating some additional logic.
func AddArmorItem(key string, item Item) {
	ArmorItems[key] = item
	persist.Save("armoritems", ArmorItems)
}

// HealthItems holds a collection of all health items in the game.
var HealthItems Items = make(Items)

// AddHealthItem adds an item to the "HealthItems"-map while encapsulating some additional logic.
func AddHealthItem(key string, item Item) {
	HealthItems[key] = item
	persist.Save("healthitems", HealthItems)
}

// AuxiliaryItems holds a collection of all auxiliary items in the game.
var AuxiliaryItems Items = make(Items)

// AddAuxiliaryItem adds an item to the "AuxiliaryItems"-map while encapsulating some additional logic.
func AddAuxiliaryItem(key string, item Item) {
	AuxiliaryItems[key] = item
	persist.Save("auxiliaryitems", AuxiliaryItems)
}

func init() {
	sword := Item{Name: "Sword", Cost: 10, DamageIncrease: 10, HealthIncrease: 0, ArmorIncrease: 0}
	AddDamageItem("sword", sword)
	chainvest := Item{Name: "Chainvest", Cost: 10, DamageIncrease: 0, HealthIncrease: 0, ArmorIncrease: 10}
	AddArmorItem("chainvest", chainvest)
	lifefountain := Item{Name: "Life Fountain", Cost: 10, DamageIncrease: 0, HealthIncrease: 10, ArmorIncrease: 0}
	AddHealthItem("lifefountain", lifefountain)
	healthpot := Item{Name: "Health Pot", Cost: 1, DamageIncrease: 0, HealthIncrease: 20, ArmorIncrease: 0, Description: "+20 Health"}
	AddAuxiliaryItem("healthpot", healthpot)
	wrath := Item{Name: "Wrath", Cost: 1, DamageIncrease: 20, HealthIncrease: 0, ArmorIncrease: 0, Description: "+20 Damage"}
	AddAuxiliaryItem("wrath", wrath)
	morphine := Item{Name: "Morphine", Cost: 1, DamageIncrease: 0, HealthIncrease: 0, ArmorIncrease: 20, Description: "+20 Armor"}
	AddAuxiliaryItem("morphine", morphine)
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
