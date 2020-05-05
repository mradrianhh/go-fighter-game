package models

import "testing"

func TestPlayer_NewPlayer(t *testing.T) {
	got := NewPlayer("Adrian", 1)
	expected := Player{
		Name:   "Adrian",
		Number: 1,
		baseStats: Stats{
			Damage: 0,
			Armor:  0,
			Health: 100,
		},
		Stats: Stats{
			Damage: 0,
			Armor:  0,
			Health: 100,
		},
		auxiliaryItemEffect: Stats{
			Damage: 0,
			Armor:  0,
			Health: 0,
		},
		Gold: 0,
		Inventory: Inventory{
			DamageItem,
			ArmorItem,
			HealthItem,
			AuxiliaryItem,
		},
	}

	if got.Name != expected.Name {
		t.Errorf("Name of players don't match. Got: '%s', expected: '%s'.", got.Name, expected.Name)
	}

	if got.Number != expected.Number {
		t.Errorf("Number of players don't match. Got: '%v', expected: '%v'.", got.Number, expected.Number)
	}

	if got.Gold != expected.Gold {
		t.Errorf("Gold of players don't match. Got: '%v', expected: '%v'.", got.Gold, expected.Gold)
	}

	if got.baseStats.Damage != expected.baseStats.Damage {
		t.Errorf("Damage-baseStat of players don't match. Got: '%v', expected: '%v'.", got.baseStats.Damage, expected.baseStats.Damage)
	}

	if got.baseStats.Armor != expected.baseStats.Armor {
		t.Errorf("Armor-baseStat of players don't match. Got: '%v', expected: '%v'.", got.baseStats.Armor, expected.baseStats.Armor)
	}

	if got.baseStats.Health != expected.baseStats.Health {
		t.Errorf("Health-baseStat of players don't match. Got: '%v', expected: '%v'.", got.baseStats.Health, expected.baseStats.Health)
	}

	if got.Stats.Damage != expected.Stats.Damage {
		t.Errorf("Damage-baseStat of players don't match. Got: '%v', expected: '%v'.", got.Stats.Damage, expected.Stats.Damage)
	}

	if got.Stats.Armor != expected.Stats.Armor {
		t.Errorf("Armor-baseStat of players don't match. Got: '%v', expected: '%v'.", got.Stats.Armor, expected.Stats.Armor)
	}

	if got.Stats.Health != expected.Stats.Health {
		t.Errorf("Health-baseStat of players don't match. Got: '%v', expected: '%v'.", got.Stats.Health, expected.Stats.Health)
	}

	if got.auxiliaryItemEffect.Damage != expected.auxiliaryItemEffect.Damage {
		t.Errorf("Damage-baseStat of players don't match. Got: '%v', expected: '%v'.", got.auxiliaryItemEffect.Damage, expected.auxiliaryItemEffect.Damage)
	}

	if got.auxiliaryItemEffect.Armor != expected.auxiliaryItemEffect.Armor {
		t.Errorf("Armor-baseStat of players don't match. Got: '%v', expected: '%v'.", got.auxiliaryItemEffect.Armor, expected.auxiliaryItemEffect.Armor)
	}

	if got.auxiliaryItemEffect.Health != expected.auxiliaryItemEffect.Health {
		t.Errorf("Health-baseStat of players don't match. Got: '%v', expected: '%v'.", got.auxiliaryItemEffect.Health, expected.auxiliaryItemEffect.Health)
	}

	if got.Inventory[0] != expected.Inventory[0] {
		t.Errorf("Damage-slot of players don't match. Got: '%s', expected: '%s'.", got.Inventory[0].Name, expected.Inventory[0].Name)
	}

	if got.Inventory[1] != expected.Inventory[1] {
		t.Errorf("Armor-slot of players don't match. Got: '%s', expected: '%s'.", got.Inventory[1].Name, expected.Inventory[1].Name)
	}

	if got.Inventory[2] != expected.Inventory[2] {
		t.Errorf("Health-slot of players don't match. Got: '%s', expected: '%s'.", got.Inventory[2].Name, expected.Inventory[2].Name)
	}

	if got.Inventory[3] != expected.Inventory[3] {
		t.Errorf("AuxiliaryItem-slot of players don't match. Got: '%s', expected: '%s'.", got.Inventory[3].Name, expected.Inventory[3].Name)
	}
}

func TestPlayer_NewPlayerTrainer(t *testing.T) {
	got := NewPlayerTrainer()
	expected := Player{
		Name:   "Player",
		Number: 1,
		baseStats: Stats{
			Damage: 0,
			Armor:  0,
			Health: 100,
		},
		Stats: Stats{
			Damage: 10,
			Armor:  10,
			Health: 110,
		},
		auxiliaryItemEffect: Stats{
			Damage: 0,
			Armor:  0,
			Health: 0,
		},
		Gold: 0,
		Inventory: Inventory{
			DamageItems["sword"],
			ArmorItems["chainvest"],
			HealthItems["lifefountain"],
			AuxiliaryItems["healthpot"],
		},
	}

	if got.Name != expected.Name {
		t.Errorf("Name of players don't match. Got: '%s', expected: '%s'.", got.Name, expected.Name)
	}

	if got.Number != expected.Number {
		t.Errorf("Number of players don't match. Got: '%v', expected: '%v'.", got.Number, expected.Number)
	}

	if got.Gold != expected.Gold {
		t.Errorf("Gold of players don't match. Got: '%v', expected: '%v'.", got.Gold, expected.Gold)
	}

	if got.baseStats.Damage != expected.baseStats.Damage {
		t.Errorf("Damage-baseStat of players don't match. Got: '%v', expected: '%v'.", got.baseStats.Damage, expected.baseStats.Damage)
	}

	if got.baseStats.Armor != expected.baseStats.Armor {
		t.Errorf("Armor-baseStat of players don't match. Got: '%v', expected: '%v'.", got.baseStats.Armor, expected.baseStats.Armor)
	}

	if got.baseStats.Health != expected.baseStats.Health {
		t.Errorf("Health-baseStat of players don't match. Got: '%v', expected: '%v'.", got.baseStats.Health, expected.baseStats.Health)
	}

	if got.Stats.Damage != expected.Stats.Damage {
		t.Errorf("Damage-Stat of players don't match. Got: '%v', expected: '%v'.", got.Stats.Damage, expected.Stats.Damage)
	}

	if got.Stats.Armor != expected.Stats.Armor {
		t.Errorf("Armor-Stat of players don't match. Got: '%v', expected: '%v'.", got.Stats.Armor, expected.Stats.Armor)
	}

	if got.Stats.Health != expected.Stats.Health {
		t.Errorf("Health-Stat of players don't match. Got: '%v', expected: '%v'.", got.Stats.Health, expected.Stats.Health)
	}

	if got.auxiliaryItemEffect.Damage != expected.auxiliaryItemEffect.Damage {
		t.Errorf("Damage-auxiliaryEffect of players don't match. Got: '%v', expected: '%v'.", got.auxiliaryItemEffect.Damage, expected.auxiliaryItemEffect.Damage)
	}

	if got.auxiliaryItemEffect.Armor != expected.auxiliaryItemEffect.Armor {
		t.Errorf("Armor-auxiliaryEffect of players don't match. Got: '%v', expected: '%v'.", got.auxiliaryItemEffect.Armor, expected.auxiliaryItemEffect.Armor)
	}

	if got.auxiliaryItemEffect.Health != expected.auxiliaryItemEffect.Health {
		t.Errorf("Health-auxiliaryEffect of players don't match. Got: '%v', expected: '%v'.", got.auxiliaryItemEffect.Health, expected.auxiliaryItemEffect.Health)
	}

	if got.Inventory[0] != expected.Inventory[0] {
		t.Errorf("Damage-slot of players don't match. Got: '%s', expected: '%s'.", got.Inventory[0].Name, expected.Inventory[0].Name)
	}

	if got.Inventory[1] != expected.Inventory[1] {
		t.Errorf("Armor-slot of players don't match. Got: '%s', expected: '%s'.", got.Inventory[1].Name, expected.Inventory[1].Name)
	}

	if got.Inventory[2] != expected.Inventory[2] {
		t.Errorf("Health-slot of players don't match. Got: '%s', expected: '%s'.", got.Inventory[2].Name, expected.Inventory[2].Name)
	}

	if got.Inventory[3] != expected.Inventory[3] {
		t.Errorf("AuxiliaryItem-slot of players don't match. Got: '%s', expected: '%s'.", got.Inventory[3].Name, expected.Inventory[3].Name)
	}
}

func TestPlayer_NewEnemyTrainer(t *testing.T) {
	got := NewEnemyTrainer()
	expected := Player{
		Name:   "Enemy",
		Number: 2,
		baseStats: Stats{
			Damage: 0,
			Armor:  0,
			Health: 100,
		},
		Stats: Stats{
			Damage: 0,
			Armor:  0,
			Health: 100,
		},
		auxiliaryItemEffect: Stats{
			Damage: 0,
			Armor:  0,
			Health: 0,
		},
		Gold: 0,
		Inventory: Inventory{
			DamageItem,
			ArmorItem,
			HealthItem,
			AuxiliaryItem,
		},
	}

	if got.Name != expected.Name {
		t.Errorf("Name of players don't match. Got: '%s', expected: '%s'.", got.Name, expected.Name)
	}

	if got.Number != expected.Number {
		t.Errorf("Number of players don't match. Got: '%v', expected: '%v'.", got.Number, expected.Number)
	}

	if got.Gold != expected.Gold {
		t.Errorf("Gold of players don't match. Got: '%v', expected: '%v'.", got.Gold, expected.Gold)
	}

	if got.baseStats.Damage != expected.baseStats.Damage {
		t.Errorf("Damage-baseStat of players don't match. Got: '%v', expected: '%v'.", got.baseStats.Damage, expected.baseStats.Damage)
	}

	if got.baseStats.Armor != expected.baseStats.Armor {
		t.Errorf("Armor-baseStat of players don't match. Got: '%v', expected: '%v'.", got.baseStats.Armor, expected.baseStats.Armor)
	}

	if got.baseStats.Health != expected.baseStats.Health {
		t.Errorf("Health-baseStat of players don't match. Got: '%v', expected: '%v'.", got.baseStats.Health, expected.baseStats.Health)
	}

	if got.Stats.Damage != expected.Stats.Damage {
		t.Errorf("Damage-Stat of players don't match. Got: '%v', expected: '%v'.", got.Stats.Damage, expected.Stats.Damage)
	}

	if got.Stats.Armor != expected.Stats.Armor {
		t.Errorf("Armor-Stat of players don't match. Got: '%v', expected: '%v'.", got.Stats.Armor, expected.Stats.Armor)
	}

	if got.Stats.Health != expected.Stats.Health {
		t.Errorf("Health-Stat of players don't match. Got: '%v', expected: '%v'.", got.Stats.Health, expected.Stats.Health)
	}

	if got.auxiliaryItemEffect.Damage != expected.auxiliaryItemEffect.Damage {
		t.Errorf("Damage-auxiliaryEffect of players don't match. Got: '%v', expected: '%v'.", got.auxiliaryItemEffect.Damage, expected.auxiliaryItemEffect.Damage)
	}

	if got.auxiliaryItemEffect.Armor != expected.auxiliaryItemEffect.Armor {
		t.Errorf("Armor-auxiliaryEffect of players don't match. Got: '%v', expected: '%v'.", got.auxiliaryItemEffect.Armor, expected.auxiliaryItemEffect.Armor)
	}

	if got.auxiliaryItemEffect.Health != expected.auxiliaryItemEffect.Health {
		t.Errorf("Health-auxiliaryEffect of players don't match. Got: '%v', expected: '%v'.", got.auxiliaryItemEffect.Health, expected.auxiliaryItemEffect.Health)
	}

	if got.Inventory[0] != expected.Inventory[0] {
		t.Errorf("Damage-slot of players don't match. Got: '%s', expected: '%s'.", got.Inventory[0].Name, expected.Inventory[0].Name)
	}

	if got.Inventory[1] != expected.Inventory[1] {
		t.Errorf("Armor-slot of players don't match. Got: '%s', expected: '%s'.", got.Inventory[1].Name, expected.Inventory[1].Name)
	}

	if got.Inventory[2] != expected.Inventory[2] {
		t.Errorf("Health-slot of players don't match. Got: '%s', expected: '%s'.", got.Inventory[2].Name, expected.Inventory[2].Name)
	}

	if got.Inventory[3] != expected.Inventory[3] {
		t.Errorf("AuxiliaryItem-slot of players don't match. Got: '%s', expected: '%s'.", got.Inventory[3].Name, expected.Inventory[3].Name)
	}
}

func TestPlayer_ResetToDefault(t *testing.T) {
	got := NewPlayer("Adrian", 1)
	got.Reset("ToDefault")
	expected := Player{
		Name:   "Adrian",
		Number: 1,
		baseStats: Stats{
			Damage: 0,
			Armor:  0,
			Health: 100,
		},
		Stats: Stats{
			Damage: 0,
			Armor:  0,
			Health: 100,
		},
		auxiliaryItemEffect: Stats{
			Damage: 0,
			Armor:  0,
			Health: 0,
		},
		Gold: 0,
		Inventory: Inventory{
			DamageItem,
			ArmorItem,
			HealthItem,
			AuxiliaryItem,
		},
	}

	if got.Name != expected.Name {
		t.Errorf("Name of players don't match. Got: '%s', expected: '%s'.", got.Name, expected.Name)
	}

	if got.Number != expected.Number {
		t.Errorf("Number of players don't match. Got: '%v', expected: '%v'.", got.Number, expected.Number)
	}

	if got.Gold != expected.Gold {
		t.Errorf("Gold of players don't match. Got: '%v', expected: '%v'.", got.Gold, expected.Gold)
	}

	if got.baseStats.Damage != expected.baseStats.Damage {
		t.Errorf("Damage-baseStat of players don't match. Got: '%v', expected: '%v'.", got.baseStats.Damage, expected.baseStats.Damage)
	}

	if got.baseStats.Armor != expected.baseStats.Armor {
		t.Errorf("Armor-baseStat of players don't match. Got: '%v', expected: '%v'.", got.baseStats.Armor, expected.baseStats.Armor)
	}

	if got.baseStats.Health != expected.baseStats.Health {
		t.Errorf("Health-baseStat of players don't match. Got: '%v', expected: '%v'.", got.baseStats.Health, expected.baseStats.Health)
	}

	if got.Stats.Damage != expected.Stats.Damage {
		t.Errorf("Damage-Stat of players don't match. Got: '%v', expected: '%v'.", got.Stats.Damage, expected.Stats.Damage)
	}

	if got.Stats.Armor != expected.Stats.Armor {
		t.Errorf("Armor-Stat of players don't match. Got: '%v', expected: '%v'.", got.Stats.Armor, expected.Stats.Armor)
	}

	if got.Stats.Health != expected.Stats.Health {
		t.Errorf("Health-Stat of players don't match. Got: '%v', expected: '%v'.", got.Stats.Health, expected.Stats.Health)
	}

	if got.auxiliaryItemEffect.Damage != expected.auxiliaryItemEffect.Damage {
		t.Errorf("Damage-auxiliaryEffect of players don't match. Got: '%v', expected: '%v'.", got.auxiliaryItemEffect.Damage, expected.auxiliaryItemEffect.Damage)
	}

	if got.auxiliaryItemEffect.Armor != expected.auxiliaryItemEffect.Armor {
		t.Errorf("Armor-auxiliaryEffect of players don't match. Got: '%v', expected: '%v'.", got.auxiliaryItemEffect.Armor, expected.auxiliaryItemEffect.Armor)
	}

	if got.auxiliaryItemEffect.Health != expected.auxiliaryItemEffect.Health {
		t.Errorf("Health-auxiliaryEffect of players don't match. Got: '%v', expected: '%v'.", got.auxiliaryItemEffect.Health, expected.auxiliaryItemEffect.Health)
	}

	if got.Inventory[0] != expected.Inventory[0] {
		t.Errorf("Damage-slot of players don't match. Got: '%s', expected: '%s'.", got.Inventory[0].Name, expected.Inventory[0].Name)
	}

	if got.Inventory[1] != expected.Inventory[1] {
		t.Errorf("Armor-slot of players don't match. Got: '%s', expected: '%s'.", got.Inventory[1].Name, expected.Inventory[1].Name)
	}

	if got.Inventory[2] != expected.Inventory[2] {
		t.Errorf("Health-slot of players don't match. Got: '%s', expected: '%s'.", got.Inventory[2].Name, expected.Inventory[2].Name)
	}

	if got.Inventory[3] != expected.Inventory[3] {
		t.Errorf("AuxiliaryItem-slot of players don't match. Got: '%s', expected: '%s'.", got.Inventory[3].Name, expected.Inventory[3].Name)
	}
}

func TestPlayer_ResetToTrainingOpponent(t *testing.T) {
	got := NewPlayer("Adrian", 1)
	got.Reset("ToTrainingOpponent")
	expected := Player{
		Name:   "Adrian",
		Number: 1,
		baseStats: Stats{
			Damage: 0,
			Armor:  0,
			Health: 100,
		},
		Stats: Stats{
			Damage: 0,
			Armor:  0,
			Health: 100,
		},
		auxiliaryItemEffect: Stats{
			Damage: 0,
			Armor:  0,
			Health: 0,
		},
		Gold: 0,
		Inventory: Inventory{
			DamageItem,
			ArmorItem,
			HealthItem,
			AuxiliaryItem,
		},
	}

	if got.Name != expected.Name {
		t.Errorf("Name of players don't match. Got: '%s', expected: '%s'.", got.Name, expected.Name)
	}

	if got.Number != expected.Number {
		t.Errorf("Number of players don't match. Got: '%v', expected: '%v'.", got.Number, expected.Number)
	}

	if got.Gold != expected.Gold {
		t.Errorf("Gold of players don't match. Got: '%v', expected: '%v'.", got.Gold, expected.Gold)
	}

	if got.baseStats.Damage != expected.baseStats.Damage {
		t.Errorf("Damage-baseStat of players don't match. Got: '%v', expected: '%v'.", got.baseStats.Damage, expected.baseStats.Damage)
	}

	if got.baseStats.Armor != expected.baseStats.Armor {
		t.Errorf("Armor-baseStat of players don't match. Got: '%v', expected: '%v'.", got.baseStats.Armor, expected.baseStats.Armor)
	}

	if got.baseStats.Health != expected.baseStats.Health {
		t.Errorf("Health-baseStat of players don't match. Got: '%v', expected: '%v'.", got.baseStats.Health, expected.baseStats.Health)
	}

	if got.Stats.Damage != expected.Stats.Damage {
		t.Errorf("Damage-Stat of players don't match. Got: '%v', expected: '%v'.", got.Stats.Damage, expected.Stats.Damage)
	}

	if got.Stats.Armor != expected.Stats.Armor {
		t.Errorf("Armor-Stat of players don't match. Got: '%v', expected: '%v'.", got.Stats.Armor, expected.Stats.Armor)
	}

	if got.Stats.Health != expected.Stats.Health {
		t.Errorf("Health-Stat of players don't match. Got: '%v', expected: '%v'.", got.Stats.Health, expected.Stats.Health)
	}

	if got.auxiliaryItemEffect.Damage != expected.auxiliaryItemEffect.Damage {
		t.Errorf("Damage-auxiliaryEffect of players don't match. Got: '%v', expected: '%v'.", got.auxiliaryItemEffect.Damage, expected.auxiliaryItemEffect.Damage)
	}

	if got.auxiliaryItemEffect.Armor != expected.auxiliaryItemEffect.Armor {
		t.Errorf("Armor-auxiliaryEffect of players don't match. Got: '%v', expected: '%v'.", got.auxiliaryItemEffect.Armor, expected.auxiliaryItemEffect.Armor)
	}

	if got.auxiliaryItemEffect.Health != expected.auxiliaryItemEffect.Health {
		t.Errorf("Health-auxiliaryEffect of players don't match. Got: '%v', expected: '%v'.", got.auxiliaryItemEffect.Health, expected.auxiliaryItemEffect.Health)
	}

	if got.Inventory[0] != expected.Inventory[0] {
		t.Errorf("Damage-slot of players don't match. Got: '%s', expected: '%s'.", got.Inventory[0].Name, expected.Inventory[0].Name)
	}

	if got.Inventory[1] != expected.Inventory[1] {
		t.Errorf("Armor-slot of players don't match. Got: '%s', expected: '%s'.", got.Inventory[1].Name, expected.Inventory[1].Name)
	}

	if got.Inventory[2] != expected.Inventory[2] {
		t.Errorf("Health-slot of players don't match. Got: '%s', expected: '%s'.", got.Inventory[2].Name, expected.Inventory[2].Name)
	}

	if got.Inventory[3] != expected.Inventory[3] {
		t.Errorf("AuxiliaryItem-slot of players don't match. Got: '%s', expected: '%s'.", got.Inventory[3].Name, expected.Inventory[3].Name)
	}
}

func TestPlayer_ResetToTrainingPlayer(t *testing.T) {
	got := NewPlayer("Adrian", 1)
	got.Reset("ToTrainingPlayer")
	expected := Player{
		Name:   "Adrian",
		Number: 1,
		baseStats: Stats{
			Damage: 0,
			Armor:  0,
			Health: 100,
		},
		Stats: Stats{
			Damage: 10,
			Armor:  10,
			Health: 110,
		},
		auxiliaryItemEffect: Stats{
			Damage: 0,
			Armor:  0,
			Health: 0,
		},
		Gold: 0,
		Inventory: Inventory{
			DamageItems["sword"],
			ArmorItems["chainvest"],
			HealthItems["lifefountain"],
			AuxiliaryItems["healthpot"],
		},
	}

	if got.Name != expected.Name {
		t.Errorf("Name of players don't match. Got: '%s', expected: '%s'.", got.Name, expected.Name)
	}

	if got.Number != expected.Number {
		t.Errorf("Number of players don't match. Got: '%v', expected: '%v'.", got.Number, expected.Number)
	}

	if got.Gold != expected.Gold {
		t.Errorf("Gold of players don't match. Got: '%v', expected: '%v'.", got.Gold, expected.Gold)
	}

	if got.baseStats.Damage != expected.baseStats.Damage {
		t.Errorf("Damage-baseStat of players don't match. Got: '%v', expected: '%v'.", got.baseStats.Damage, expected.baseStats.Damage)
	}

	if got.baseStats.Armor != expected.baseStats.Armor {
		t.Errorf("Armor-baseStat of players don't match. Got: '%v', expected: '%v'.", got.baseStats.Armor, expected.baseStats.Armor)
	}

	if got.baseStats.Health != expected.baseStats.Health {
		t.Errorf("Health-baseStat of players don't match. Got: '%v', expected: '%v'.", got.baseStats.Health, expected.baseStats.Health)
	}

	if got.Stats.Damage != expected.Stats.Damage {
		t.Errorf("Damage-Stat of players don't match. Got: '%v', expected: '%v'.", got.Stats.Damage, expected.Stats.Damage)
	}

	if got.Stats.Armor != expected.Stats.Armor {
		t.Errorf("Armor-Stat of players don't match. Got: '%v', expected: '%v'.", got.Stats.Armor, expected.Stats.Armor)
	}

	if got.Stats.Health != expected.Stats.Health {
		t.Errorf("Health-Stat of players don't match. Got: '%v', expected: '%v'.", got.Stats.Health, expected.Stats.Health)
	}

	if got.auxiliaryItemEffect.Damage != expected.auxiliaryItemEffect.Damage {
		t.Errorf("Damage-auxiliaryEffect of players don't match. Got: '%v', expected: '%v'.", got.auxiliaryItemEffect.Damage, expected.auxiliaryItemEffect.Damage)
	}

	if got.auxiliaryItemEffect.Armor != expected.auxiliaryItemEffect.Armor {
		t.Errorf("Armor-auxiliaryEffect of players don't match. Got: '%v', expected: '%v'.", got.auxiliaryItemEffect.Armor, expected.auxiliaryItemEffect.Armor)
	}

	if got.auxiliaryItemEffect.Health != expected.auxiliaryItemEffect.Health {
		t.Errorf("Health-auxiliaryEffect of players don't match. Got: '%v', expected: '%v'.", got.auxiliaryItemEffect.Health, expected.auxiliaryItemEffect.Health)
	}

	if got.Inventory[0] != expected.Inventory[0] {
		t.Errorf("Damage-slot of players don't match. Got: '%s', expected: '%s'.", got.Inventory[0].Name, expected.Inventory[0].Name)
	}

	if got.Inventory[1] != expected.Inventory[1] {
		t.Errorf("Armor-slot of players don't match. Got: '%s', expected: '%s'.", got.Inventory[1].Name, expected.Inventory[1].Name)
	}

	if got.Inventory[2] != expected.Inventory[2] {
		t.Errorf("Health-slot of players don't match. Got: '%s', expected: '%s'.", got.Inventory[2].Name, expected.Inventory[2].Name)
	}

	if got.Inventory[3] != expected.Inventory[3] {
		t.Errorf("AuxiliaryItem-slot of players don't match. Got: '%s', expected: '%s'.", got.Inventory[3].Name, expected.Inventory[3].Name)
	}
}

func TestPlayer_ResetDefault(t *testing.T) {
	got := NewPlayer("Adrian", 1)
	err := got.Reset("not a ResetType")
	if err.Error() != "not a ResetType" {
		t.Errorf("Did not get the expected result. Got: %s, expected: %s.", err.Error(), "not a ResetType")
	}
}

func TestPlayer_UpdateStats(t *testing.T) {
	got := NewPlayer("Adrian", 1)
	got.Inventory[0] = DamageItems["sword"]
	got.UpdateStats()
	expected := Player{
		Name:   "Adrian",
		Number: 1,
		baseStats: Stats{
			Damage: 0,
			Armor:  0,
			Health: 100,
		},
		Stats: Stats{
			Damage: 10,
			Armor:  0,
			Health: 100,
		},
		auxiliaryItemEffect: Stats{
			Damage: 0,
			Armor:  0,
			Health: 0,
		},
		Gold: 0,
		Inventory: Inventory{
			DamageItems["sword"],
			ArmorItem,
			HealthItem,
			AuxiliaryItem,
		},
	}

	if got.Name != expected.Name {
		t.Errorf("Name of players don't match. Got: '%s', expected: '%s'.", got.Name, expected.Name)
	}

	if got.Number != expected.Number {
		t.Errorf("Number of players don't match. Got: '%v', expected: '%v'.", got.Number, expected.Number)
	}

	if got.Gold != expected.Gold {
		t.Errorf("Gold of players don't match. Got: '%v', expected: '%v'.", got.Gold, expected.Gold)
	}

	if got.baseStats.Damage != expected.baseStats.Damage {
		t.Errorf("Damage-baseStat of players don't match. Got: '%v', expected: '%v'.", got.baseStats.Damage, expected.baseStats.Damage)
	}

	if got.baseStats.Armor != expected.baseStats.Armor {
		t.Errorf("Armor-baseStat of players don't match. Got: '%v', expected: '%v'.", got.baseStats.Armor, expected.baseStats.Armor)
	}

	if got.baseStats.Health != expected.baseStats.Health {
		t.Errorf("Health-baseStat of players don't match. Got: '%v', expected: '%v'.", got.baseStats.Health, expected.baseStats.Health)
	}

	if got.Stats.Damage != expected.Stats.Damage {
		t.Errorf("Damage-Stat of players don't match. Got: '%v', expected: '%v'.", got.Stats.Damage, expected.Stats.Damage)
	}

	if got.Stats.Armor != expected.Stats.Armor {
		t.Errorf("Armor-Stat of players don't match. Got: '%v', expected: '%v'.", got.Stats.Armor, expected.Stats.Armor)
	}

	if got.Stats.Health != expected.Stats.Health {
		t.Errorf("Health-Stat of players don't match. Got: '%v', expected: '%v'.", got.Stats.Health, expected.Stats.Health)
	}

	if got.auxiliaryItemEffect.Damage != expected.auxiliaryItemEffect.Damage {
		t.Errorf("Damage-auxiliaryEffect of players don't match. Got: '%v', expected: '%v'.", got.auxiliaryItemEffect.Damage, expected.auxiliaryItemEffect.Damage)
	}

	if got.auxiliaryItemEffect.Armor != expected.auxiliaryItemEffect.Armor {
		t.Errorf("Armor-auxiliaryEffect of players don't match. Got: '%v', expected: '%v'.", got.auxiliaryItemEffect.Armor, expected.auxiliaryItemEffect.Armor)
	}

	if got.auxiliaryItemEffect.Health != expected.auxiliaryItemEffect.Health {
		t.Errorf("Health-auxiliaryEffect of players don't match. Got: '%v', expected: '%v'.", got.auxiliaryItemEffect.Health, expected.auxiliaryItemEffect.Health)
	}

	if got.Inventory[0] != expected.Inventory[0] {
		t.Errorf("Damage-slot of players don't match. Got: '%s', expected: '%s'.", got.Inventory[0].Name, expected.Inventory[0].Name)
	}

	if got.Inventory[1] != expected.Inventory[1] {
		t.Errorf("Armor-slot of players don't match. Got: '%s', expected: '%s'.", got.Inventory[1].Name, expected.Inventory[1].Name)
	}

	if got.Inventory[2] != expected.Inventory[2] {
		t.Errorf("Health-slot of players don't match. Got: '%s', expected: '%s'.", got.Inventory[2].Name, expected.Inventory[2].Name)
	}

	if got.Inventory[3] != expected.Inventory[3] {
		t.Errorf("AuxiliaryItem-slot of players don't match. Got: '%s', expected: '%s'.", got.Inventory[3].Name, expected.Inventory[3].Name)
	}
}

func TestPlayer_WireMoney_NotEnough(t *testing.T) {
	player := NewPlayer("Adrian", 1)
	got := player.WireMoney(10)
	expected := "not enough money"
	if got.Error() != expected {
		t.Errorf("Did not get the error expected. Got: %s, expected: %s.", got.Error(), expected)
	}
}

func TestPlayer_WireMoney_GiveAway(t *testing.T) {
	player := NewPlayer("Adrian", 1)
	player.Gold = 10
	player.WireMoney(10)
	if player.Gold != 0 {
		t.Errorf("Did not get the expected result. Got: %v gold, expected: %v gold.", player.Gold, 0)
	}
}

func TestPlayer_WireMoney_Receive(t *testing.T) {
	player := NewPlayer("Adrian", 1)
	player.WireMoney(-10)
	if player.Gold != 10 {
		t.Errorf("Did not get the expected result. Got: %v gold, expected: %v gold.", player.Gold, 10)
	}
}

func TestPlayer_BuyDamageItem_NotEnough(t *testing.T) {
	player := NewPlayer("Adrian", 1)
	got := player.BuyDamageItem(DamageItems["sword"])
	expected := "not enough money"
	if got.Error() != expected {
		t.Errorf("Did not get the error expected. Got: %s, expected: %s.", got.Error(), expected)
	}
}

func TestPlayer_BuyDamageItem(t *testing.T) {
	player := NewPlayer("Adrian", 1)
	player.Gold = 10
	player.BuyDamageItem(DamageItems["sword"])

	if player.Gold != 0 {
		t.Errorf("Incorrect gold amount. Got: %v gold, expected: %v gold.", player.Gold, 0)
	}

	if player.Inventory[0] != DamageItems["sword"] {
		t.Errorf("Incorrect item. Got: %s, expected: %s.", player.Inventory[0].Name, DamageItems["sword"].Name)
	}
}

func TestPlayer_BuyArmorItem_NotEnough(t *testing.T) {
	player := NewPlayer("Adrian", 1)
	got := player.BuyArmorItem(ArmorItems["chainvest"])
	expected := "not enough money"
	if got.Error() != expected {
		t.Errorf("Did not get the error expected. Got: %s, expected: %s.", got.Error(), expected)
	}
}

func TestPlayer_BuyArmorItem(t *testing.T) {
	player := NewPlayer("Adrian", 1)
	player.Gold = 10
	player.BuyArmorItem(ArmorItems["chainvest"])

	if player.Gold != 0 {
		t.Errorf("Incorrect gold amount. Got: %v gold, expected: %v gold.", player.Gold, 0)
	}

	if player.Inventory[1] != ArmorItems["chainvest"] {
		t.Errorf("Incorrect item. Got: %s, expected: %s.", player.Inventory[1].Name, ArmorItems["chainvest"].Name)
	}
}

func TestPlayer_BuyHealthItem_NotEnough(t *testing.T) {
	player := NewPlayer("Adrian", 1)
	got := player.BuyHealthItem(HealthItems["lifefountain"])
	expected := "not enough money"
	if got.Error() != expected {
		t.Errorf("Did not get the error expected. Got: %s, expected: %s.", got.Error(), expected)
	}
}

func TestPlayer_BuyHealthItem(t *testing.T) {
	player := NewPlayer("Adrian", 1)
	player.Gold = 10
	player.BuyHealthItem(HealthItems["lifefountain"])

	if player.Gold != 0 {
		t.Errorf("Incorrect gold amount. Got: %v gold, expected: %v gold.", player.Gold, 0)
	}

	if player.Inventory[2] != HealthItems["lifefountain"] {
		t.Errorf("Incorrect item. Got: %s, expected: %s.", player.Inventory[2].Name, HealthItems["lifefountain"].Name)
	}
}

func TestPlayer_BuyAuxiliaryItem_NotEnough(t *testing.T) {
	player := NewPlayer("Adrian", 1)
	got := player.BuyHealthItem(AuxiliaryItems["healthpot"])
	expected := "not enough money"
	if got.Error() != expected {
		t.Errorf("Did not get the error expected. Got: %s, expected: %s.", got.Error(), expected)
	}
}

func TestPlayer_BuyAuxiliaryItem(t *testing.T) {
	player := NewPlayer("Adrian", 1)
	player.Gold = 10
	player.BuyAuxiliaryItem(AuxiliaryItems["healthpot"])

	if player.Gold != 9 {
		t.Errorf("Incorrect gold amount. Got: %v gold, expected: %v gold.", player.Gold, 9)
	}

	if player.Inventory[3] != AuxiliaryItems["healthpot"] {
		t.Errorf("Incorrect item. Got: %s, expected: %s.", player.Inventory[3].Name, AuxiliaryItems["healthpot"].Name)
	}
}

func TestPlayer_SellDamageItem(t *testing.T) {
	player := NewPlayer("Adrian", 1)
	player.Inventory[0] = DamageItems["sword"]
	player.SellDamageItem()

	if player.Gold != 10 {
		t.Errorf("Incorrect gold amount. Got: %v gold, expected: %v gold.", player.Gold, 10)
	}

	if player.Inventory[0] != DamageItem {
		t.Errorf("Incorrect item. Got: %s, expected: %s.", player.Inventory[0].Name, DamageItem.Name)
	}
}

func TestPlayer_SellArmorItem(t *testing.T) {
	player := NewPlayer("Adrian", 1)
	player.Inventory[1] = ArmorItems["chainvest"]
	player.SellArmorItem()

	if player.Gold != 10 {
		t.Errorf("Incorrect gold amount. Got: %v gold, expected: %v gold.", player.Gold, 10)
	}

	if player.Inventory[1] != ArmorItem {
		t.Errorf("Incorrect item. Got: %s, expected: %s.", player.Inventory[1].Name, ArmorItem.Name)
	}
}

func TestPlayer_SellHealthItem(t *testing.T) {
	player := NewPlayer("Adrian", 1)
	player.Inventory[2] = HealthItems["lifefountain"]
	player.SellHealthItem()

	if player.Gold != 10 {
		t.Errorf("Incorrect gold amount. Got: %v gold, expected: %v gold.", player.Gold, 10)
	}

	if player.Inventory[2] != HealthItem {
		t.Errorf("Incorrect item. Got: %s, expected: %s.", player.Inventory[2].Name, HealthItem.Name)
	}
}

func TestPlayer_SellAuxiliaryItem(t *testing.T) {
	player := NewPlayer("Adrian", 1)
	player.Inventory[3] = AuxiliaryItems["healthpot"]
	player.SellAuxiliaryItem()

	if player.Gold != 1 {
		t.Errorf("Incorrect gold amount. Got: %v gold, expected: %v gold.", player.Gold, 1)
	}

	if player.Inventory[3] != AuxiliaryItem {
		t.Errorf("Incorrect item. Got: %s, expected: %s.", player.Inventory[3].Name, AuxiliaryItem.Name)
	}
}

func TestPlayer_ConsumeAuxiliaryItem_ConsumePlaceholderItem(t *testing.T) {
	player := NewPlayer("Adrian", 1)
	got := player.ConsumeAuxiliaryItem()
	expected := "can't consume placeholder item"
	if got.Error() != expected {
		t.Errorf("Did not get the error expected. Got: %s, expected: %s.", got.Error(), expected)
	}
}

func TestPlayer_ConsumeAuxiliaryItem_Damage(t *testing.T) {
	player := NewPlayer("Adrian", 1)
	player.Inventory[3] = AuxiliaryItems["wrath"]
	player.ConsumeAuxiliaryItem()

	if player.Stats.Damage != 20 {
		t.Errorf("Did not get the expected damage. Got: %v, expected: %v.", player.Stats.Damage, 20)
	}
}

func TestPlayer_ConsumeAuxiliaryItem_Armor(t *testing.T) {
	player := NewPlayer("Adrian", 1)
	player.Inventory[3] = AuxiliaryItems["morphine"]
	player.ConsumeAuxiliaryItem()

	if player.Stats.Armor != 20 {
		t.Errorf("Did not get the expected armor. Got: %v, expected: %v.", player.Stats.Armor, 20)
	}
}

func TestPlayer_RemoveAuxiliaryItemEffect(t *testing.T) {
	player := NewPlayer("Adrian", 1)
	player.Inventory[3] = AuxiliaryItems["morphine"]
	player.ConsumeAuxiliaryItem()
	player.RemoveAuxiliaryItemEffect()

	if player.Stats.Armor != 0 {
		t.Errorf("Did not get the expected armor. Got: %v, expected: %v.", player.Stats.Armor, 0)
	}

	if player.Stats.Damage != 0 {
		t.Errorf("Did not get the expected damage. Got: %v, expected: %v.", player.Stats.Damage, 0)
	}

	if player.Stats.Health != 100 {
		t.Errorf("Did not get the expected health. Got: %v, expected: %v.", player.Stats.Health, 100)
	}

	if player.auxiliaryItemEffect.Armor != 0 {
		t.Errorf("Did not get the expected armor. Got: %v, expected: %v.", player.auxiliaryItemEffect.Armor, 0)
	}

	if player.auxiliaryItemEffect.Damage != 0 {
		t.Errorf("Did not get the expected damage. Got: %v, expected: %v.", player.auxiliaryItemEffect.Damage, 0)
	}

	if player.auxiliaryItemEffect.Health != 0 {
		t.Errorf("Did not get the expected health. Got: %v, expected: %v.", player.auxiliaryItemEffect.Health, 0)
	}
}
