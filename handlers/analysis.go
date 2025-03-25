package handlers

import (
	"SalesDetails-CSV/config"
	"encoding/json"
	"net/http"
)

// GetTotalRevenue calculates the total revenue from sales
func GetTotalRevenue(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var revenue float64
	err := config.DB.QueryRow(`
		SELECT SUM(p.unit_price * o.quantity_sold * (1 - o.discount) + o.shipping_cost) 
		FROM orders o
		JOIN products p ON o.product_id = p.product_id
	`).Scan(&revenue)

	if err != nil {
		http.Error(w, "Error fetching revenue: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]float64{"total_revenue": revenue})
}

// GetTopProducts fetches the top 5 best-selling products
func GetTopProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := config.DB.Query(`
		SELECT p.product_name, SUM(o.quantity_sold) AS total_sold
		FROM orders o
		JOIN products p ON o.product_id = p.product_id
		GROUP BY p.product_name
		ORDER BY total_sold DESC
		LIMIT 5
	`)
	if err != nil {
		http.Error(w, "Error fetching top products: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type Product struct {
		Name      string `json:"product_name"`
		TotalSold int    `json:"total_sold"`
	}

	var products []Product
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.Name, &p.TotalSold)
		if err != nil {
			http.Error(w, "Error scanning data: "+err.Error(), http.StatusInternalServerError)
			return
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Error reading data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(products)
}

// GetTotalCustomers calculates the total number of customers
func GetTotalCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var count int
	err := config.DB.QueryRow("SELECT COUNT(DISTINCT customer_id) FROM orders").Scan(&count)
	if err != nil {
		http.Error(w, "Error fetching customer count: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]int{"total_customers": count})
}

// LoadData loads CSV data
func LoadData(w http.ResponseWriter, r *http.Request) {
	err := config.LoadCSVData("/home/kundhavk/SalesDetails-CSV/handlers/sales_data.csv")
	if err != nil {
		http.Error(w, "Error loading CSV", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("✅ Data loaded successfully!"))
}
