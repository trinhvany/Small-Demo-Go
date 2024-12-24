package types

import "time"

type Order struct {
	ID             int
	CustomerID     int
	ProductID      int
	ProductName    string
	Quantity       int
	TotalAmount    float64
	Discount       float64
	Tax            float64
	Status         int
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Email      	   string
}
