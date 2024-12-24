package common

import (
	"myproject/types"
	"sync"
)

func ApplyDiscount(order *types.Order, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()

	if order.TotalAmount > 200 {
		order.Tax = order.TotalAmount * 0.1
		order.Discount = order.TotalAmount * 0.1
	}
}
