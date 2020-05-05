package models

import "testing"

func TestItem_DamageItem(t *testing.T) {
	damageitem := DamageItem

	if damageitem.Name != "Damage Item" {
		t.Errorf("Did not get expected result. Got: %s, expected: %s.", damageitem.Name, "Damage Item")
	}

	if damageitem.Cost != 0 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", damageitem.Cost, 0)
	}

	if damageitem.DamageIncrease != 0 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", damageitem.DamageIncrease, 0)
	}

	if damageitem.ArmorIncrease != 0 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", damageitem.ArmorIncrease, 0)
	}

	if damageitem.HealthIncrease != 0 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", damageitem.HealthIncrease, 0)
	}
}

func TestItem_ArmorItem(t *testing.T) {
	armoritem := ArmorItem

	if armoritem.Name != "Armor Item" {
		t.Errorf("Did not get expected result. Got: %s, expected: %s.", armoritem.Name, "Armor Item")
	}

	if armoritem.Cost != 0 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", armoritem.Cost, 0)
	}

	if armoritem.DamageIncrease != 0 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", armoritem.DamageIncrease, 0)
	}

	if armoritem.ArmorIncrease != 0 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", armoritem.ArmorIncrease, 0)
	}

	if armoritem.HealthIncrease != 0 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", armoritem.HealthIncrease, 0)
	}
}

func TestItem_HealthItem(t *testing.T) {
	healthitem := HealthItem

	if healthitem.Name != "Health Item" {
		t.Errorf("Did not get expected result. Got: %s, expected: %s.", healthitem.Name, "Health Item")
	}

	if healthitem.Cost != 0 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", healthitem.Cost, 0)
	}

	if healthitem.DamageIncrease != 0 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", healthitem.DamageIncrease, 0)
	}

	if healthitem.ArmorIncrease != 0 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", healthitem.ArmorIncrease, 0)
	}

	if healthitem.HealthIncrease != 0 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", healthitem.HealthIncrease, 0)
	}
}

func TestItem_AuxiliaryItem(t *testing.T) {
	auxiliaryitem := AuxiliaryItem

	if auxiliaryitem.Name != "Auxiliary Item" {
		t.Errorf("Did not get expected result. Got: %s, expected: %s.", auxiliaryitem.Name, "Auxiliary Item")
	}

	if auxiliaryitem.Cost != 0 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", auxiliaryitem.Cost, 0)
	}

	if auxiliaryitem.DamageIncrease != 0 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", auxiliaryitem.DamageIncrease, 0)
	}

	if auxiliaryitem.ArmorIncrease != 0 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", auxiliaryitem.ArmorIncrease, 0)
	}

	if auxiliaryitem.HealthIncrease != 0 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", auxiliaryitem.HealthIncrease, 0)
	}
}

func TestItem_SwordItem(t *testing.T) {
	sword := DamageItems["sword"]

	if sword.Name != "Sword" {
		t.Errorf("Did not get expected result. Got: %s, expected: %s.", sword.Name, "Sword")
	}

	if sword.Cost != 10 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", sword.Cost, 10)
	}

	if sword.DamageIncrease != 10 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", sword.DamageIncrease, 10)
	}

	if sword.ArmorIncrease != 0 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", sword.ArmorIncrease, 0)
	}

	if sword.HealthIncrease != 0 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", sword.HealthIncrease, 0)
	}
}

func TestItem_LifefountainItem(t *testing.T) {
	lifefountain := HealthItems["lifefountain"]

	if lifefountain.Name != "Life Fountain" {
		t.Errorf("Did not get expected result. Got: %s, expected: %s.", lifefountain.Name, "Life Fountain")
	}

	if lifefountain.Cost != 10 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", lifefountain.Cost, 10)
	}

	if lifefountain.DamageIncrease != 0 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", lifefountain.DamageIncrease, 0)
	}

	if lifefountain.ArmorIncrease != 0 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", lifefountain.ArmorIncrease, 0)
	}

	if lifefountain.HealthIncrease != 10 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", lifefountain.HealthIncrease, 10)
	}
}

func TestItem_HealthpotItem(t *testing.T) {
	healthpot := AuxiliaryItems["healthpot"]

	if healthpot.Name != "Health Pot" {
		t.Errorf("Did not get expected result. Got: %s, expected: %s.", healthpot.Name, "Health Pot")
	}

	if healthpot.Description != "+20 Health" {
		t.Errorf("Did not get expected result. Got: %s, expected: %s.", healthpot.Description, "+20 Health")
	}

	if healthpot.Cost != 1 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", healthpot.Cost, 1)
	}

	if healthpot.DamageIncrease != 0 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", healthpot.DamageIncrease, 0)
	}

	if healthpot.ArmorIncrease != 0 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", healthpot.ArmorIncrease, 0)
	}

	if healthpot.HealthIncrease != 20 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", healthpot.HealthIncrease, 20)
	}
}

func TestItem_WrathItem(t *testing.T) {
	wrath := AuxiliaryItems["wrath"]

	if wrath.Name != "Wrath" {
		t.Errorf("Did not get expected result. Got: %s, expected: %s.", wrath.Name, "Wrath")
	}

	if wrath.Description != "+20 Damage" {
		t.Errorf("Did not get expected result. Got: %s, expected: %s.", wrath.Description, "+20 Damage")
	}

	if wrath.Cost != 1 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", wrath.Cost, 1)
	}

	if wrath.DamageIncrease != 20 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", wrath.DamageIncrease, 20)
	}

	if wrath.ArmorIncrease != 0 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", wrath.ArmorIncrease, 0)
	}

	if wrath.HealthIncrease != 0 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", wrath.HealthIncrease, 0)
	}
}

func TestItem_MorphineItem(t *testing.T) {
	morphine := AuxiliaryItems["morphine"]

	if morphine.Name != "Morphine" {
		t.Errorf("Did not get expected result. Got: %s, expected: %s.", morphine.Name, "Morphine")
	}

	if morphine.Description != "+20 Armor" {
		t.Errorf("Did not get expected result. Got: %s, expected: %s.", morphine.Description, "+20 Armor")
	}

	if morphine.Cost != 1 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", morphine.Cost, 1)
	}

	if morphine.DamageIncrease != 0 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", morphine.DamageIncrease, 0)
	}

	if morphine.ArmorIncrease != 20 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", morphine.ArmorIncrease, 20)
	}

	if morphine.HealthIncrease != 0 {
		t.Errorf("Did not get expected result. Got: %v, expected: %v.", morphine.HealthIncrease, 0)
	}
}
