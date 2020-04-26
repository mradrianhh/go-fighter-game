package models

import (
	"fmt"
)

// Comment is a response to a users actions.
type Comment struct {
	Text string
}

// Speak presents the comment to the user on screen.
func (c *Comment) Speak() {
	fmt.Println(c.Text)
}
