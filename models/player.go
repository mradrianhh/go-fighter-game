package models

// Player will be controlled by the user.
//
// The player object will have a damage stat, health stat
// and a gold stat. Future changes may incurr.
type Player struct {
	Name   string
	Damage int
	Health int
	Gold   int
}
