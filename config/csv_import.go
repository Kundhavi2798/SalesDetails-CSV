package config

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/google/uuid"
)

type SalesData struct {
	OrderID       uuid.UUID
	ProductID     uuid.UUID
	CustomerID    uuid.UUID
	ProductName   string
	Category      string
	Region        string
	DateOfSale    string
	QuantitySold  int
	UnitPrice     float64
	Discount      float64
	ShippingCost  float64
	PaymentMethod string
	CustomerName  string
	CustomerEmail string
	CustomerAddr  string
}

func LoadCSVData(filePath string) error {
	fmt.Println("load")
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for _, row := range records[1:] { // Skip header
		quantity, _ := strconv.Atoi(row[7])
		unitPrice, _ := strconv.ParseFloat(row[8], 64)
		discount, _ := strconv.ParseFloat(row[9], 64)
		shippingCost, _ := strconv.ParseFloat(row[10], 64)

		_, err := DB.Exec(`INSERT INTO customers (customer_name, customer_email, customer_address) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING`,
			row[12], row[13], row[14])

		_, err = DB.Exec(`INSERT INTO products (product_name, category, unit_price) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING`,
			row[3], row[4], unitPrice)

		_, err = DB.Exec(`INSERT INTO orders (date_of_sale, payment_method, shipping_cost) VALUES ($1, $2, $3)`,
			row[6], row[11], shippingCost)

		_, err = DB.Exec(`INSERT INTO sales (order_id, product_id, quantity_sold, discount) VALUES ($1, $2, $3, $4)`,
			row[0], row[1], quantity, discount)

		if err != nil {
			fmt.Println("Insert Error:", err)
		}
	}

	fmt.Println("✅ CSV Data Loaded Successfully!")
	return nil
}
