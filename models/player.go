package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/mradrianhh/go-fighter-game/events"
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
	Name                string
	Number              int
	Stats               Stats
	Gold                int
	Inventory           Inventory
	baseStats           Stats
	auxiliaryItemEffect Stats
}

// PlayerList is the list of all players in the game.
var PlayerList []Player

func init() {
	PlayerList = append(PlayerList, NewPlayer("Adrian", 1), NewPlayer("Jack", 2))
}

// NewPlayer creates and initializes a new player with the correct default values and given name.
func NewPlayer(name string, number int) Player {
	player := Player{
		Name:      name,
		Number:    number,
		baseStats: NewBaseStats(),
		Gold:      0,
		Inventory: Inventory{
			DamageItem,
			ArmorItem,
			HealthItem,
			AuxiliaryItem,
		},
	}
	player.UpdateStats()
	return player
}

// NewPlayerTrainer returns a player that is initialized to be used in training by the user.
func NewPlayerTrainer() Player {
	player := Player{
		Name:      "Player",
		Number:    1,
		baseStats: NewBaseStats(),
		Gold:      0,
		Inventory: Inventory{
			DamageItems["sword"],
			ArmorItems["chainvest"],
			HealthItems["lifefountain"],
			AuxiliaryItems["healthpot"],
		},
	}
	player.UpdateStats()

	return player
}

// NewEnemyTrainer returns a player that is initialized to be faced in training by the user.
func NewEnemyTrainer() Player {
	player := Player{
		Name:      "Enemy",
		Number:    2,
		baseStats: NewBaseStats(),
		Gold:      0,
		Inventory: Inventory{
			DamageItem,
			ArmorItem,
			HealthItem,
			AuxiliaryItem,
		},
	}
	player.UpdateStats()

	return player
}

// Reset changes everything but the name back to training or default specifications based on the reset type provided.
func (player *Player) Reset(resetType ResetType) error {
	switch resetType {
	case ToTrainingPlayer:
		player.baseStats = NewBaseStats()
		player.Gold = 0
		player.Inventory = Inventory{
			DamageItems["sword"],
			ArmorItems["chainvest"],
			HealthItems["lifefountain"],
			AuxiliaryItems["healthpot"],
		}
		player.UpdateStats()
		return nil
	case ToTrainingOpponent:
		player.baseStats = NewBaseStats()
		player.Gold = 0
		player.Inventory = Inventory{
			DamageItem,
			ArmorItem,
			HealthItem,
			AuxiliaryItem,
		}
		player.UpdateStats()
		return nil
	case ToDefault:
		player.baseStats = NewBaseStats()
		player.Gold = 0
		player.Inventory = Inventory{
			DamageItem,
			ArmorItem,
			HealthItem,
			AuxiliaryItem,
		}
		player.UpdateStats()
		return nil
	default:
		return errors.New("not a ResetType")
	}
}

// UpdateStats updates the players current stats based on his inventory.
func (player *Player) UpdateStats() {
	player.Stats.Damage = player.baseStats.Damage + player.auxiliaryItemEffect.Damage
	player.Stats.Armor = player.baseStats.Armor + player.auxiliaryItemEffect.Armor
	player.Stats.Health = player.baseStats.Health + player.auxiliaryItemEffect.Health
	for i := 0; i < len(player.Inventory)-1; i++ { // Subtract one from limit to not include the auxiliary(consumable) item.
		player.Stats.Damage += player.Inventory[i].DamageIncrease
		player.Stats.Armor += player.Inventory[i].ArmorIncrease
		player.Stats.Health += player.Inventory[i].HealthIncrease
	}
}

// WireMoney updates the players current gold based on the amount.
func (player *Player) WireMoney(amount int) error {
	if player.Gold-amount >= 0 {
		player.Gold -= amount
		return nil
	}

	return errors.New("not enough money")
}

// BuyDamageItem replaces the current damage item in the inventory,
// and handles payment.
func (player *Player) BuyDamageItem(item Item) error {
	if err := player.WireMoney(item.Cost); err != nil {
		return err
	}
	player.Inventory[0] = item
	player.UpdateStats()

	return nil
}

// BuyArmorItem replaces the current armor item in the inventory,
// and handles payment.
func (player *Player) BuyArmorItem(item Item) error {
	err := player.WireMoney(item.Cost)
	if err != nil {
		return err
	}
	player.Inventory[1] = item
	player.UpdateStats()

	return nil
}

// BuyHealthItem replaces the current health item in the inventory,
// and handles payment.
func (player *Player) BuyHealthItem(item Item) error {
	err := player.WireMoney(item.Cost)
	if err != nil {
		return err
	}
	player.Inventory[2] = item
	player.UpdateStats()

	return nil
}

// BuyAuxiliaryItem replaces the current auxiliary item in the inventory,
// and handles payment.
func (player *Player) BuyAuxiliaryItem(item Item) error {
	err := player.WireMoney(item.Cost)
	if err != nil {
		return err
	}
	player.Inventory[3] = item
	player.UpdateStats()

	return nil
}

// SellDamageItem removes the current damage item in the inventory,
// and handles payment.
func (player *Player) SellDamageItem() {
	player.WireMoney(-player.Inventory[0].Cost)
	player.Inventory[0] = DamageItem
	player.UpdateStats()
}

// SellArmorItem removes the current armor item in the inventory,
// and handles payment.
func (player *Player) SellArmorItem() {
	player.WireMoney(-player.Inventory[1].Cost)
	player.Inventory[1] = ArmorItem
	player.UpdateStats()
}

// SellHealthItem removes the current health item in the inventory,
// and handles payment.
func (player *Player) SellHealthItem() {
	player.WireMoney(-player.Inventory[2].Cost)
	player.Inventory[2] = HealthItem
	player.UpdateStats()
}

// SellAuxiliaryItem removes the current auxiliary item in the inventory,
// and handles payment.
func (player *Player) SellAuxiliaryItem() {
	player.WireMoney(-player.Inventory[3].Cost)
	player.Inventory[3] = AuxiliaryItem
	player.UpdateStats()
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
	player.auxiliaryItemEffect = Stats{
		Damage: player.Inventory[3].DamageIncrease,
		Armor:  player.Inventory[3].ArmorIncrease,
		Health: player.Inventory[3].HealthIncrease,
	}
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

// RemoveAuxiliaryItemEffect removes the auxiliary item effect from the player after one turn. Will not remove health regen effects, only armor and damage increases.
func (player *Player) RemoveAuxiliaryItemEffect() {
	fmt.Printf("Removing auxiliary effect on player %v, %s\n", player.Number, player.Name)
	player.auxiliaryItemEffect = Stats{}
	player.UpdateStats()
}

func (player *Player) handleChangeTurn() {
	player.RemoveAuxiliaryItemEffect()
}

// Listen will listen for events and act accordingly.
func Listen() {
	for {
		if event, ok := <-events.EventStream; ok {
			if event == "NextTurn" {
				for _, p := range PlayerList {
					p.handleChangeTurn()
				}
			}
		}
	}
}
