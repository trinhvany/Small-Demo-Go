package common

import (
	"myproject/types"
	"myproject/db"
	"sync"
	"fmt"
)

func ValidateOrder(order *types.Order, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	database, err := db.ConnectDB()
	defer database.Close()
	if err != nil {
		fmt.Errorf("Error database: %v", err)
	}
	products, err := db.GetListProduct(database);
	if err != nil {
		fmt.Errorf("Error when getting order list: %v", err)
	}
	for _, product := range products {
		if product.ID == order.ProductID {
			if product.Quantity < order.Quantity {
				message := fmt.Sprintf("There is not enough stock to create order for product %s with quantity of %v items", product.ProductName, order.Quantity)
				SendMail(order.Email, order.ProductName, order.Quantity, nil)
				ch <- message
				return
			}
			order.Status = 1
			order.ProductName = product.ProductName
			return
		}
	}
}
