package models

// Stats is a holder for the stats damage, health and armor.
type Stats struct {
	Damage int
	Armor  int
	Health int
}

// BaseStats is the default and standard stats a player-object should be initialized with.
var baseStats = Stats{
	Damage: 0,
	Armor:  0,
	Health: 100,
}

// NewStats returns a new stats-object with the stats specified.
func NewStats(damage, armor, health int) Stats {
	return Stats{
		Damage: damage,
		Armor:  armor,
		Health: health,
	}
}

// NewBaseStats returns a new stats-object with the "BaseStats"(10, 0, 100).
func NewBaseStats() Stats {
	return baseStats
}
