package commentary

import "github.com/mradrianhh/go-fighter-game/models"

// Encourage chooses a positive comment
// and prints it to the screen with the Speak() command.
func Encourage() {
	c := models.Comment{Text: "Well done!"}
	c.Speak()
}

// Disgrace chooses a negative comment
// and prints it to the screen with the Speak() command.
func Disgrace() {
	c := models.Comment{Text: "You're useless!"}
	c.Speak()
}

// WrongInput informs the player that he has entered an option which is not viable.
func WrongInput() {
	c := models.Comment{Text: "Sorry, I can't understand. Try again."}
	c.Speak()
}
