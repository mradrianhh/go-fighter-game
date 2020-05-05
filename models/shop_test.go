package models

import "testing"

func TestShop_NewShop(t *testing.T) {
	shop := NewShop()

	if shop.Inventory[0] != DamageItems["sword"] {
		t.Errorf("Did not get the expected result. Got: %s, expected: %s.", shop.Inventory[0].Name, "Sword")
	}

	if shop.Inventory[1] != ArmorItems["chainvest"] {
		t.Errorf("Did not get the expected result. Got: %s, expected: %s.", shop.Inventory[1].Name, "Chainvest")
	}

	if shop.Inventory[2] != HealthItems["lifefountain"] {
		t.Errorf("Did not get the expected result. Got: %s, expected: %s.", shop.Inventory[2].Name, "Life Fountain")
	}

	if shop.Inventory[3] != AuxiliaryItems["healthpot"] {
		t.Errorf("Did not get the expected result. Got: %s, expected: %s.", shop.Inventory[3].Name, "Health Pot")
	}
}

func TestShop_BuyDamageItem_Shop_NotEnough(t *testing.T) {
	player := NewPlayer("Adrian", 1)
	shop := NewShop()

	err := shop.BuyDamageItem(&player)

	if err.Error() != "not enough money" {
		t.Errorf("Did not get the expected result. Got: %s, expected: %s.", err.Error(), "not enough money")
	}

	if shop.Inventory[0] == DamageItem {
		t.Errorf("Did not get the expected result. Got: %s, expected random damage item.", shop.Inventory[0].Name)
	}
}

func TestShop_BuyDamageItem_Shop(t *testing.T) {
	player := NewPlayer("Adrian", 1)
	player.Gold = 10
	shop := NewShop()

	err := shop.BuyDamageItem(&player)

	if err != nil {
		t.Errorf("Did not get the expected result. Got: %s, expected no error.", err.Error())
	}

	if shop.Inventory[0] != DamageItem {
		t.Errorf("Did not get the expected result. Got: %s, expected: %s.", shop.Inventory[0].Name, DamageItem.Name)
	}
}

func TestShop_BuyArmorItem_Shop_NotEnough(t *testing.T) {
	player := NewPlayer("Adrian", 1)
	shop := NewShop()

	err := shop.BuyArmorItem(&player)

	if err.Error() != "not enough money" {
		t.Errorf("Did not get the expected result. Got: %s, expected: %s.", err.Error(), "not enough money")
	}

	if shop.Inventory[1] == ArmorItem {
		t.Errorf("Did not get the expected result. Got: %s, expected random armor item.", shop.Inventory[0].Name)
	}
}

func TestShop_BuyArmorItem_Shop(t *testing.T) {
	player := NewPlayer("Adrian", 1)
	player.Gold = 10
	shop := NewShop()

	err := shop.BuyArmorItem(&player)

	if err != nil {
		t.Errorf("Did not get the expected result. Got: %s, expected no error.", err.Error())
	}

	if shop.Inventory[1] != ArmorItem {
		t.Errorf("Did not get the expected result. Got: %s, expected: %s.", shop.Inventory[1].Name, ArmorItem.Name)
	}
}

func TestShop_BuyHealthItem_Shop_NotEnough(t *testing.T) {
	player := NewPlayer("Adrian", 1)
	shop := NewShop()

	err := shop.BuyHealthItem(&player)

	if err.Error() != "not enough money" {
		t.Errorf("Did not get the expected result. Got: %s, expected: %s.", err.Error(), "not enough money")
	}

	if shop.Inventory[2] == HealthItem {
		t.Errorf("Did not get the expected result. Got: %s, expected random health item.", shop.Inventory[2].Name)
	}
}

func TestShop_BuyHealthItem_Shop(t *testing.T) {
	player := NewPlayer("Adrian", 1)
	player.Gold = 10
	shop := NewShop()

	err := shop.BuyHealthItem(&player)

	if err != nil {
		t.Errorf("Did not get the expected result. Got: %s, expected no error.", err.Error())
	}

	if shop.Inventory[2] != HealthItem {
		t.Errorf("Did not get the expected result. Got: %s, expected: %s.", shop.Inventory[2].Name, HealthItem.Name)
	}
}

func TestShop_BuyAuxiliaryItem_Shop_NotEnough(t *testing.T) {
	player := NewPlayer("Adrian", 1)
	shop := NewShop()

	err := shop.BuyAuxiliaryItem(&player)

	if err.Error() != "not enough money" {
		t.Errorf("Did not get the expected result. Got: %s, expected: %s.", err.Error(), "not enough money")
	}

	if shop.Inventory[3] == AuxiliaryItem {
		t.Errorf("Did not get the expected result. Got: %s, expected random auxiliary item.", shop.Inventory[3].Name)
	}
}

func TestShop_BuyAuxiliaryItem_Shop(t *testing.T) {
	player := NewPlayer("Adrian", 1)
	player.Gold = 1
	shop := NewShop()

	err := shop.BuyAuxiliaryItem(&player)

	if err != nil {
		t.Errorf("Did not get the expected result. Got: %s, expected no error.", err.Error())
	}

	if shop.Inventory[3] != AuxiliaryItem {
		t.Errorf("Did not get the expected result. Got: %s, expected: %s.", shop.Inventory[3].Name, AuxiliaryItem.Name)
	}
}
