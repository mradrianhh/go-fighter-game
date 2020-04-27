package models

import (
	"errors"
)

// Player will be controlled by the user.
//
// The player object will have a damage stat, health stat
// and a gold stat. Future changes may incurr.
type Player struct {
	Name      string
	Damage    int
	Health    int
	Armor     int
	Gold      int
	Inventory Inventory
}

// NewPlayer creates and initializes a new player with the correct default values and given name.
func NewPlayer(name string) Player {
	return Player{
		Name:   name,
		Damage: 10,
		Health: 100,
		Armor:  20,
		Gold:   0,
		Inventory: Inventory{
			DamageItem,
			ArmorItem,
			HealthItem,
			AuxiliaryItem,
		},
	}
}

// UpdateStats updates the players current stats based on his inventory.
func (player *Player) UpdateStats() {
	for i := 0; i < len(player.Inventory)-1; i++ {
		player.Damage += player.Inventory[i].DamageIncrease
		player.Armor += player.Inventory[i].ArmorIncrease
		player.Health += player.Inventory[i].HealthIncrease
	}
}

// WireMoney updates the players current gold based on the amount.
func (player *Player) WireMoney(amount int) error {
	if player.Gold+amount >= 0 {
		player.Gold += amount
		return nil
	}

	return errors.New("not enough money")
}

// BuyDamageItem replaces the current damage item in the inventory,
// and handles payment.
func (player *Player) BuyDamageItem(item Item) error {
	err := player.WireMoney(-item.Cost)
	if err != nil {
		return err
	}
	player.Inventory[0] = item

	return nil
}

// BuyArmorItem replaces the current armor item in the inventory,
// and handles payment.
func (player *Player) BuyArmorItem(item Item) error {
	err := player.WireMoney(-item.Cost)
	if err != nil {
		return err
	}
	player.Inventory[1] = item

	return nil
}

// BuyHealthItem replaces the current health item in the inventory,
// and handles payment.
func (player *Player) BuyHealthItem(item Item) error {
	err := player.WireMoney(-item.Cost)
	if err != nil {
		return err
	}
	player.Inventory[2] = item

	return nil
}

// BuyAuxiliaryItem replaces the current auxiliary item in the inventory,
// and handles payment.
func (player *Player) BuyAuxiliaryItem(item Item) error {
	err := player.WireMoney(-item.Cost)
	if err != nil {
		return err
	}
	player.Inventory[3] = item

	return nil
}

// SellDamageItem removes the current damage item in the inventory,
// and handles payment.
func (player *Player) SellDamageItem() {
	player.WireMoney(player.Inventory[0].Cost)
	player.Inventory[0] = DamageItem
}

// SellArmorItem removes the current armor item in the inventory,
// and handles payment.
func (player *Player) SellArmorItem() {
	player.WireMoney(player.Inventory[1].Cost)
	player.Inventory[1] = ArmorItem
}

// SellHealthItem removes the current health item in the inventory,
// and handles payment.
func (player *Player) SellHealthItem() {
	player.WireMoney(player.Inventory[2].Cost)
	player.Inventory[2] = HealthItem
}

// SellAuxiliaryItem removes the current auxiliary item in the inventory,
// and handles payment.
func (player *Player) SellAuxiliaryItem() {
	player.WireMoney(player.Inventory[3].Cost)
	player.Inventory[3] = AuxiliaryItem
}

// ConsumeAuxiliaryItem adds the auxiliary item's stats to the player, then removes it.
func (player *Player) ConsumeAuxiliaryItem() error {
	if player.Inventory[3] == AuxiliaryItem {
		return errors.New("can't consume placeholder item")
	}
	player.Damage += player.Inventory[3].DamageIncrease
	player.Armor += player.Inventory[3].ArmorIncrease
	player.Health += player.Inventory[3].HealthIncrease

	return nil
}
