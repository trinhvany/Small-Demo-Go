package main

import (
	"fmt"
	"myproject/common"
	"myproject/db"
	"sync"
)


func main() {
	database, err := db.ConnectDB()
	defer database.Close()
	if err != nil {
		fmt.Errorf("Error database: %v", err)
	}
	orders, err := db.GetListOrder(database);
	if err != nil {
		fmt.Errorf("Error when getting order list: %v", err)
	}
	var wg sync.WaitGroup;
	ch := make(chan string, 100)
	message := fmt.Sprintf("Processing %v orders", len(orders))
	fmt.Println(message)
	for i := range orders {
		wg.Add(2)
		go common.ValidateOrder(&orders[i], &wg, ch);
		go common.ApplyDiscount(&orders[i], &wg, ch);
	}
	
	wg.Wait()
	close(ch)

	for msg := range ch {
		fmt.Println(msg)
	}
	var wgMail sync.WaitGroup
	fmt.Println(orders)
	for i := range orders {
		if (orders[i].Status != 1) {
			continue
		}
		wgMail.Add(2)
		go common.SendMail(orders[i].Email, orders[i].ProductName, orders[i].Quantity, &wgMail)
		go db.UpdateStatusOrder(orders[i].ID, database, &wgMail)
	}
	wgMail.Wait()
	fmt.Println("Batch processing completed!")
}
