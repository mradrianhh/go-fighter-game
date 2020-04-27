package game

import (
	"fmt"

	"github.com/mradrianhh/go-fighter-game/commentary"
	"github.com/mradrianhh/go-fighter-game/models"
)

// ShowStats presents the user with his stats.
// Here the user can see his damage, health etc.
func ShowStats(player *models.Player) {
	for {
		fmt.Print(separator)
		fmt.Printf("\n\tPlayer %d\n", currentPlayer.Number)
		fmt.Print(separator)
		fmt.Print("\nStats\n")
		fmt.Printf("\n\tGold: %d\n", player.Gold)
		fmt.Printf("\tDamage: %d\n", player.Damage)
		fmt.Printf("\tArmor: %d\n", player.Armor)
		fmt.Printf("\tHealth: %d\n", player.Health)
		fmt.Printf("\n\tInventory: \n")
		fmt.Printf("\t1 - %s(+%d Damage) | 2 - %s(+%d Armor) | 3 - %s(+%d Health) | 4 - %s(%s)\n", player.Inventory[0].Name, player.Inventory[0].DamageIncrease,
			player.Inventory[1].Name, player.Inventory[1].ArmorIncrease, player.Inventory[2].Name, player.Inventory[2].HealthIncrease,
			player.Inventory[3].Name, player.Inventory[3].Description)
		fmt.Print("\n1 - Continue\n")
		if _, err := fmt.Scan(&response); err == nil {
			switch response {
			case 1:
				goto end
			default:
				commentary.WrongInput()
			}
		}
	}
end:
}
