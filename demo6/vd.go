package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// Struct để đại diện cho một đối tượng trong cơ sở dữ liệu
type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// Biến db để đại diện cho kết nối cơ sở dữ liệu
var db *sql.DB

func main() {
	// Khởi tạo kết nối đến cơ sở dữ liệu MySQL
	var err error
	db, err = sql.Open("mysql", "@tcp(localhost:3306)/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Khởi động HTTP server
	http.HandleFunc("/items", handleItems)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Xử lý yêu cầu tới endpoint "/items"
func handleItems(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		getItems(w, req)
	case http.MethodPost:
		createItem(w, req)
	case http.MethodPut:
		updateItem(w, req)
	case http.MethodDelete:
		deleteItem(w, req)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Xử lý yêu cầu GET để lấy danh sách các item
func getItems(w http.ResponseWriter, _ *http.Request) {
	rows, err := db.Query("SELECT id, name, price FROM items")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.ID, &item.Name, &item.Price); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		items = append(items, item)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// Xử lý yêu cầu POST để tạo một item mới
func createItem(w http.ResponseWriter, req *http.Request) {
	var newItem Item
	if err := json.NewDecoder(req.Body).Decode(&newItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := db.Exec("INSERT INTO items (name, price) VALUES (?, ?)", newItem.Name, newItem.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Xử lý yêu cầu PUT để cập nhật một item
func updateItem(w http.ResponseWriter, req *http.Request) {
	// Code xử lý cập nhật item ở đây
}

// Xử lý yêu cầu DELETE để xóa một item
func deleteItem(w http.ResponseWriter, req *http.Request) {
	// Code xử lý xóa item ở đây
}
