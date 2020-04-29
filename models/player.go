package models

import (
	"errors"
	"fmt"
	"time"
)

// ResetType is the type that can be tested against in the "Reset()"-method.
type ResetType = string

// List of reset types that can be passed to "Reset()".
const (
	ToTrainingOpponent = ResetType("ToTrainingOpponent")
	ToTrainingPlayer   = ResetType("ToTrainingPlayer")
	ToDefault          = ResetType("ToDefault")
)

// Player will be controlled by the user.
//
// The player object will have a damage stat, health stat
// and a gold stat. Future changes may incurr.
type Player struct {
	Name      string
	Number    int
	Stats     Stats
	Gold      int
	Inventory Inventory
}

// PlayerList is the list of all players in the game.
var PlayerList []Player

func init() {
	PlayerList = append(PlayerList, NewPlayer("Adrian", 1), NewPlayer("Jack", 2))
}

// NewPlayer creates and initializes a new player with the correct default values and given name.
func NewPlayer(name string, number int) Player {
	return Player{
		Name:   name,
		Number: number,
		Stats:  NewBaseStats(),
		Gold:   0,
		Inventory: Inventory{
			DamageItem,
			ArmorItem,
			HealthItem,
			AuxiliaryItem,
		},
	}
}

// NewPlayerTrainer returns a player that is initialized to be used in training by the user.
func NewPlayerTrainer() Player {
	return Player{
		Name:   "Player",
		Number: 1,
		Stats:  NewBaseStats(),
		Gold:   0,
		Inventory: Inventory{
			DamageItems["sword"],
			ArmorItems["chainvest"],
			HealthItems["lifefountain"],
			AuxiliaryItems["healthpot"],
		},
	}
}

// NewEnemyTrainer returns a player that is initialized to be faced in training by the user.
func NewEnemyTrainer() Player {
	return Player{
		Name:   "Enemy",
		Number: 2,
		Stats:  NewBaseStats(),
		Gold:   0,
		Inventory: Inventory{
			DamageItem,
			ArmorItem,
			HealthItem,
			AuxiliaryItem,
		},
	}
}

// Reset changes everything but the name back to training or default specifications based on the reset type provided.
func (player *Player) Reset(resetType ResetType) error {
	switch resetType {
	case ToTrainingPlayer:
		player = &Player{
			Name:   player.Name,
			Number: player.Number,
			Stats:  NewBaseStats(),
			Gold:   0,
			Inventory: Inventory{
				DamageItems["sword"],
				ArmorItems["chainvest"],
				HealthItems["lifefountain"],
				AuxiliaryItems["healthpot"],
			},
		}
		return nil
	case ToTrainingOpponent:
		player = &Player{
			Name:   player.Name,
			Number: player.Number,
			Stats:  NewBaseStats(),
			Gold:   0,
			Inventory: Inventory{
				DamageItem,
				ArmorItem,
				HealthItem,
				AuxiliaryItem,
			},
		}
		return nil
	case ToDefault:
		player = &Player{
			Name:   player.Name,
			Number: player.Number,
			Stats:  NewBaseStats(),
			Gold:   0,
			Inventory: Inventory{
				DamageItem,
				ArmorItem,
				HealthItem,
				AuxiliaryItem,
			},
		}
		return nil
	default:
		return errors.New("not a ResetType")
	}
}

// UpdateStats updates the players current stats based on his inventory.
func (player *Player) UpdateStats() {
	for i := 0; i < len(player.Inventory)-1; i++ {
		player.Stats.Damage += player.Inventory[i].DamageIncrease
		player.Stats.Armor += player.Inventory[i].ArmorIncrease
		player.Stats.Health += player.Inventory[i].HealthIncrease
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
	fmt.Print("\nConsuming...\n")
	time.Sleep(1 * time.Second)
	if player.Inventory[3] == AuxiliaryItem {
		return errors.New("can't consume placeholder item")
	}

	player.Stats.Damage += player.Inventory[3].DamageIncrease
	player.Stats.Armor += player.Inventory[3].ArmorIncrease
	player.Stats.Health += player.Inventory[3].HealthIncrease

	return nil
}

// Attack damages the opponent.
func (player *Player) Attack(target *Player) {
	fmt.Print("\nAttacking...\n")
	time.Sleep(1 * time.Second)
	target.Stats.Health -= player.Stats.Damage
}

// Defend increases the health of the player by his armor.
func (player *Player) Defend() {
	fmt.Print("\nDefending...\n")
	time.Sleep(1 * time.Second)
	player.Stats.Health += player.Stats.Armor
}
