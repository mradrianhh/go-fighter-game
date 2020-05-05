package game

import (
	"fmt"
	"os"
	"time"
)

// ShowEndscreen presents the user with the end screen of the game.
func ShowEndscreen() {
	fmt.Print(separator)
	fmt.Printf("\n\tCongratulations %s, you are the winner!\n", Winner.Name)

	fmt.Print("\n\tReward: \n")
	fmt.Print("\n\tGold: 20")
	fmt.Print("\n\tExp: 50\n")
	fmt.Print("\nPress 1 to end the game...")
	if _, err := fmt.Scan(&response); err != nil {
		if response == 1 {
			fmt.Print("\n.")
			time.Sleep(100 * time.Millisecond)
			fmt.Print(".")
			time.Sleep(100 * time.Millisecond)
			fmt.Print(".")
			time.Sleep(100 * time.Millisecond)
			fmt.Print("Goodbye")
			time.Sleep(1 * time.Second)
			os.Exit(0)
		}
	}
}
