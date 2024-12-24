package db

import (
	"database/sql"
	"fmt"
	"myproject/types"
	"sync"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	// Dùng parseTime=true để MYSQL convert datetime sang time.Time
	connect := "<user>:<password>!@tcp(<host>)/<db_name>?parseTime=true"
	db, err := sql.Open("mysql", connect)
	if err != nil {
		return nil, err
	}
	return db, err
}

func GetListOrder(db *sql.DB) ([]types.Order, error) {
	rows, err := db.Query("SELECT orders.*, customer.email FROM orders JOIN customer ON orders.customer_id = customer.id WHERE orders.status = 0")
	if  err != nil {
		return nil, err
	}
	var orders []types.Order
	for rows.Next() {
		var order types.Order
		err := rows.Scan(
			&order.ID,
			&order.CustomerID,
			&order.Quantity,
			&order.ProductID,
			&order.TotalAmount,
			&order.Discount,
			&order.Tax,
			&order.Status,
			&order.UpdatedAt,
			&order.CreatedAt,
            &order.Email,
		)

		if err != nil {
			fmt.Println("Data mapping error:", err)
			return nil, err
		}

		orders = append(orders, order)
	}
	return orders, nil
}

func GetListProduct(db *sql.DB) ([]types.Product, error) {
	rows, err := db.Query("SELECT * FROM products")
	if  err != nil {
		return nil, err
	}
	var products []types.Product
	for rows.Next() {
		var product types.Product
		err := rows.Scan(
			&product.ID,
			&product.ProductName,
			&product.Quantity,
		)

		if err != nil {
			fmt.Println("Data mapping error:", err)
			return nil, err
		}

		products = append(products, product)
	}
	return products, nil
}

func UpdateStatusOrder (OrderId int,db *sql.DB, wg *sync.WaitGroup) error {
    defer wg.Done()
    query :="UPDATE orders SET status = 1 WHERE id = ?"
    _, err := db.Exec(query, OrderId)
    if err != nil {
        return fmt.Errorf("error updating order status: %v", err)
    }
    fmt.Printf("Successfully updated status of order ID %d to 1\n", OrderId)
	return nil
}