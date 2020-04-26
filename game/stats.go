package game

import (
	"fmt"

	"github.com/mradrianhh/go-fighter-game/commentary"
)

// ShowStats presents the user with his stats.
// Here the user can see his damage, health etc.
func ShowStats() {
	for {
		fmt.Print(separator)
		fmt.Print("\nStats\n")
		fmt.Print("\n\tGold: 10\n")
		fmt.Print("\tDamage: 10\n")
		fmt.Print("\tHealth: 10\n")
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
