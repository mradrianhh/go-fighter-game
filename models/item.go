package models

// Item can be bought by the player to increase his stats.
// They can be bought in exchange for gold.
type Item struct {
	Name           string
	Cost           int
	DamageIncrease int
	HealthIncrease int
}
