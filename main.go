package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Item struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

var store = map[int]Item{
	1: {ID: 1, Name: "Widget", Price: 9.99, CreatedAt: time.Now()},
	2: {ID: 2, Name: "Gadget", Price: 19.99, CreatedAt: time.Now()},
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/items", itemsHandler)
	http.HandleFunc("/items/", itemHandler)
	http.HandleFunc("/discount", discountHandler)
	http.HandleFunc("/time", timeHandler)

	log.Printf("listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func itemsHandler(w http.ResponseWriter, r *http.Request) {
	items := make([]Item, 0, len(store))
	for _, v := range store {
		items = append(items, v)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func itemHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/items/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"error":"invalid id"}`, http.StatusBadRequest)
		return
	}
	item, ok := store[id]
	if !ok {
		http.Error(w, `{"error":"not found"}`, http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

// BUG: discount uses integer division, truncating results
func discountHandler(w http.ResponseWriter, r *http.Request) {
	priceStr := r.URL.Query().Get("price")
	pctStr := r.URL.Query().Get("percent")

	price, _ := strconv.ParseFloat(priceStr, 64)
	pct, _ := strconv.Atoi(pctStr)

	// Bug: integer division truncates
	discounted := price - price*float64(pct/100)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]float64{"original": price, "discounted": discounted})
}

// BUG: returns wrong status code (200 instead of 201) for POST-like time creation
func timeHandler(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	resp := map[string]interface{}{
		"unix":   now.Unix(),
		"formatted": now.Format(time.RFC3339),
	}

	if r.URL.Query().Get("create") == "true" {
		// Should return 201 Created, but returns 200
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) // BUG: should be StatusCreated
		json.NewEncoder(w).Encode(resp)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// ApplyDiscount is the buggy function exposed for testing
func ApplyDiscount(price float64, percent int) float64 {
	return price - price*float64(percent/100)
}

// FormatPrice formats price to 2 decimal string
func FormatPrice(price float64) string {
	return fmt.Sprintf("%.2f", price)
}

// ValidatePrice checks price is positive
func ValidatePrice(price float64) error {
	if price < 0 {
		return fmt.Errorf("price must be positive, got %.2f", price)
	}
	return nil
}
