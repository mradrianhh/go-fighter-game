package game

import (
	"fmt"

	"github.com/mradrianhh/go-fighter-game/commentary"

	"github.com/mradrianhh/go-fighter-game/models"
)

// ShowShop presents the user with a shop-menu.
// Here, the user can choose to view the shop's items, view his own items, or exit.
//
// On choosing to view the shop's items, the user will be presented with another menu.
// In this menu, he can buy an item.
//
// On choosing to view his own items, the user will be presented with another menu.
// In this menu, he can sell an item.
func ShowShop() {
	for {
		fmt.Print(separator)
		fmt.Printf("\n\tPlayer %d\n", currentPlayer.Number)
		fmt.Print(separator)
		fmt.Print("\nShop\n")
		fmt.Print("\n1 - View Shop Items | 2 - View Own Items | 3 - Back\n")
		if _, err := fmt.Scan(&response); err == nil {
			switch response {
			case 1:
				ShowShopItems(&shop, &currentPlayer)
			case 2:
				ShowOwnItems(&currentPlayer)
			case 3:
				goto end
			default:
				commentary.WrongInput()
			}
		}
	}
end:
}

// ShowShopItems presents the user with the items available in the shop.
func ShowShopItems(shop *models.Shop, player *models.Player) {
	for {
		fmt.Print(separator)
		fmt.Print("\nShop Items\n")

		fmt.Print("\n(Choose item to buy it.)")
		fmt.Printf("\n1 - %s(+%d Damage) | 2 - %s(+%d Armor) | 3 - %s(+%d Health) | 4 - %s(%s) | 5 - Back\n", shop.Inventory[0].Name, shop.Inventory[0].DamageIncrease,
			shop.Inventory[1].Name, shop.Inventory[1].ArmorIncrease, shop.Inventory[2].Name, shop.Inventory[2].HealthIncrease,
			shop.Inventory[3].Name, shop.Inventory[3].Description)
		if _, err := fmt.Scan(&response); err == nil {
			fmt.Print("\nAre you sure you want to buy it?(Y/N)")

			switch response {
			case 1:
				if _, err := fmt.Scan(&charResponse); err == nil {
					switch charResponse {
					case "Y", "y":
						if shop.Inventory[0] == models.DamageItem {
							fmt.Println("This item can't be bought. It's just a placeholder.")
						} else {
							err := shop.BuyDamageItem(player)
							if err != nil {
								fmt.Println(err.Error())
							}
						}
					case "N", "n":
						continue
					default:
						commentary.WrongInput()
					}
				}
			case 2:
				if _, err := fmt.Scan(&charResponse); err == nil {
					switch charResponse {
					case "Y", "y":
						if shop.Inventory[1] == models.ArmorItem {
							fmt.Println("This item can't be bought. It's just a placeholder.")
						} else {
							err := shop.BuyArmorItem(player)
							if err != nil {
								fmt.Println(err.Error())
							}
						}
					case "N", "n":
						continue
					default:
						commentary.WrongInput()
					}
				}
			case 3:
				if _, err := fmt.Scan(&charResponse); err == nil {
					switch charResponse {
					case "Y", "y":
						if shop.Inventory[2] == models.HealthItem {
							fmt.Println("This item can't be bought. It's just a placeholder.")
						} else {
							err := shop.BuyHealthItem(player)
							if err != nil {
								fmt.Println(err.Error())
							}
						}
					case "N", "n":
						continue
					default:
						commentary.WrongInput()
					}
				}
			case 4:
				if _, err := fmt.Scan(&charResponse); err == nil {
					switch charResponse {
					case "Y", "y":
						if shop.Inventory[3] == models.AuxiliaryItem {
							fmt.Println("This item can't be bought. It's just a placeholder.")
						} else {
							err := shop.BuyAuxiliaryItem(player)
							if err != nil {
								fmt.Println(err.Error())
							}
						}
					case "N", "n":
						continue
					default:
						commentary.WrongInput()
					}
				}
			case 5:
				goto end // Leaves the infinite loop and reassigns control to the previous process.
			default:
				commentary.WrongInput()
			}
		}
	}
end:
}

// ShowOwnItems presents the user with his own items. In this menu, he can choose to sell selected items.
func ShowOwnItems(p *models.Player) {
	for {
		fmt.Print(separator)
		fmt.Print("\nOwn Items\n")

		fmt.Print("\n(Choose item to sell it.)")
		fmt.Printf("\n1 - %s(+%d Damage) | 2 - %s(+%d Armor) | 3 - %s(+%d Health) | 4 - %s(%s) | 5 - Back\n", p.Inventory[0].Name, p.Inventory[0].DamageIncrease,
			p.Inventory[1].Name, p.Inventory[1].ArmorIncrease, p.Inventory[2].Name, p.Inventory[2].HealthIncrease,
			p.Inventory[3].Name, p.Inventory[3].Description)

		if _, err := fmt.Scan(&response); err == nil {
			fmt.Print("\nAre you sure you want to sell it?(Y/N)")

			switch response {
			case 1:
				if _, err := fmt.Scan(&charResponse); err == nil {
					switch charResponse {
					case "Y", "y":
						if p.Inventory[0] == models.DamageItem {
							fmt.Println("This item can't be sold. It's just a placeholder.")
						} else {
							p.SellDamageItem()
						}
					case "N", "n":
						continue
					default:
						commentary.WrongInput()
					}
				}
			case 2:
				if _, err := fmt.Scan(&charResponse); err == nil {
					switch charResponse {
					case "Y", "y":
						if p.Inventory[1] == models.ArmorItem {
							fmt.Println("This item can't be sold. It's just a placeholder.")
						} else {
							p.SellArmorItem()
						}
					case "N", "n":
						continue
					default:
						commentary.WrongInput()
					}
				}
			case 3:
				if _, err := fmt.Scan(&charResponse); err == nil {
					switch charResponse {
					case "Y", "y":
						if p.Inventory[2] == models.HealthItem {
							fmt.Println("This item can't be sold. It's just a placeholder.")
						} else {
							p.SellHealthItem()
						}
					case "N", "n":
						continue
					default:
						commentary.WrongInput()
					}
				}
			case 4:
				if _, err := fmt.Scan(&charResponse); err == nil {
					switch charResponse {
					case "Y", "y":
						if p.Inventory[3] == models.AuxiliaryItem {
							fmt.Println("This item can't be sold. It's just a placeholder.")
						} else {
							p.SellAuxiliaryItem()
						}
					case "N", "n":
						continue
					default:
						commentary.WrongInput()
					}
				}
			case 5:
				goto end // Leaves the infinite loop and reassigns control to the previous process.
			default:
				commentary.WrongInput()
			}
		}
	}
end:
}
