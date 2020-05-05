package models

import "testing"

func TestStats_BaseStats(t *testing.T) {
	baseStats := baseStats

	if baseStats.Damage != 0 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", baseStats.Damage, 0)
	}

	if baseStats.Armor != 0 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", baseStats.Armor, 0)
	}

	if baseStats.Health != 100 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", baseStats.Health, 100)
	}
}

func TestStats_NewStats(t *testing.T) {
	stats := NewStats(10, 10, 110)

	if stats.Damage != 10 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", stats.Damage, 10)
	}

	if stats.Armor != 10 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", stats.Armor, 10)
	}

	if stats.Health != 110 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", stats.Health, 110)
	}
}

func TestStats_NewBaseStats(t *testing.T) {
	stats := NewBaseStats()

	if stats.Damage != 0 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", stats.Damage, 0)
	}

	if stats.Armor != 0 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", stats.Armor, 0)
	}

	if stats.Health != 100 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", stats.Health, 100)
	}
}
