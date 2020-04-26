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
		fmt.Print("\nShop\n")
		fmt.Print("\n1 - View Shop Items | 2 - View Own Items | 3 - Back\n")
		if _, err := fmt.Scan(&response); err == nil {
			switch response {
			case 1:
				ShowShopItems(&shop, &Player)
			case 2:
				ShowOwnItems(&Player)
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
							player.Inventory[0] = shop.Inventory[0]
							shop.Inventory[0] = models.DamageItem
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
							player.Inventory[1] = shop.Inventory[1]
							shop.Inventory[1] = models.ArmorItem
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
							player.Inventory[2] = shop.Inventory[2]
							shop.Inventory[2] = models.HealthItem
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
							player.Inventory[3] = shop.Inventory[3]
							shop.Inventory[3] = models.AuxiliaryItem
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
							p.Inventory[0] = models.DamageItem
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
							p.Inventory[1] = models.ArmorItem
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
							p.Inventory[2] = models.HealthItem
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
							p.Inventory[3] = models.AuxiliaryItem
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
