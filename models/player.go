package models

// Player will be controlled by the user.
//
// The player object will have a damage stat, health stat
// and a gold stat. Future changes may incurr.
type Player struct {
	Name      string
	Damage    int
	Health    int
	Armor     int
	Gold      int
	Inventory Inventory
}

// NewPlayer creates and initializes a new player with the correct default values and given name.
func NewPlayer(name string) Player {
	return Player{
		Name:   name,
		Damage: 10,
		Health: 100,
		Armor:  20,
		Gold:   0,
		Inventory: Inventory{
			DamageItem,
			ArmorItem,
			HealthItem,
			AuxiliaryItem,
		},
	}
}
