package models

import (
	"fmt"

	"github.com/mradrianhh/go-fighter-game/persist"
)

type comments []Comment

// PositiveComments is a collection of comments that are encouraging.
var PositiveComments comments

// NegativeComments is a collection of comments that are discouraging.
var NegativeComments comments

func init() {
	PositiveComments = comments{"Great job!", "Well done!", "Terrific!", "Wonderful!", "Fabulous!"}
	persist.Save("positivecomments", PositiveComments)
	NegativeComments = comments{"You're awful!", "Oh my god...", "Let me out of this!", "What is wrong with you?", "Are you incapable of doing anything right?"}
	persist.Save("negativecomments", NegativeComments)
}

// Comment is a response to a users actions.
type Comment string

// Speak presents the comment to the user on screen.
func (c *Comment) Speak() {
	fmt.Println(c)
}
