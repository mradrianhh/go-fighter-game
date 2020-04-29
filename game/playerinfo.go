package game

import (
	"fmt"

	"github.com/mradrianhh/go-fighter-game/models"
)

func showPlayerInfo(player models.Player) {
	fmt.Printf("\n\tGold: %d\n", player.Gold)
	fmt.Printf("\tDamage: %d\n", player.Stats.Damage)
	fmt.Printf("\tArmor: %d\n", player.Stats.Armor)
	fmt.Printf("\tHealth: %d\n", player.Stats.Health)
}
