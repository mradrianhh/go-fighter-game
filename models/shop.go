package models

// Shop is a place where the player can buy items.
type Shop struct {
	Inventory *Inventory
}

// NewShop initializes and returns a new shop.
func NewShop() Shop {
	var inventory Inventory

	inventory = Inventory{DamageItems["sword"], ArmorItems["chainvest"], HealthItems["lifefountain"], AuxiliaryItems["healthpot"]}

	shop := Shop{Inventory: &inventory}

	return shop
}

// BuyDamageItem handles the transaction between the shop and the player of
// a damage item.
func (shop *Shop) BuyDamageItem(player *Player) error {
	err := player.BuyDamageItem(shop.Inventory[0])
	if err != nil {
		return err
	}
	shop.Inventory[0] = DamageItem
	return nil
}

// BuyArmorItem handles the transaction between the shop and the player of
// a armor item.
func (shop *Shop) BuyArmorItem(player *Player) error {
	err := player.BuyArmorItem(shop.Inventory[1])
	if err != nil {
		return err
	}
	shop.Inventory[1] = DamageItem
	return nil
}

// BuyHealthItem handles the transaction between the shop and the player of
// a health item.
func (shop *Shop) BuyHealthItem(player *Player) error {
	err := player.BuyHealthItem(shop.Inventory[2])
	if err != nil {
		return err
	}
	shop.Inventory[2] = DamageItem
	return nil
}

// BuyAuxiliaryItem handles the transaction between the shop and the player of
// an auxiliary item.
func (shop *Shop) BuyAuxiliaryItem(player *Player) error {
	err := player.BuyAuxiliaryItem(shop.Inventory[3])
	if err != nil {
		return err
	}
	shop.Inventory[3] = DamageItem
	return nil
}
