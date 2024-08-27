package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Item represents a sample data model
type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// In-memory store
var items = []Item{
	{ID: 1, Name: "Item 1", Price: 100},
	{ID: 2, Name: "Item 2", Price: 200},
}

// Handlers
func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for _, item := range items {
		if item.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Item not found", http.StatusNotFound)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	var newItem Item
	_ = json.NewDecoder(r.Body).Decode(&newItem)
	newItem.ID = len(items) + 1
	items = append(items, newItem)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newItem)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for index, item := range items {
		if item.ID == id {
			_ = json.NewDecoder(r.Body).Decode(&items[index])
			items[index].ID = id
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(items[index])
			return
		}
	}
	http.Error(w, "Item not found", http.StatusNotFound)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for index, item := range items {
		if item.ID == id {
			items = append(items[:index], items[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Item not found", http.StatusNotFound)
}

func main() {
	// Initialize the router
	router := mux.NewRouter()

	// Route handlers
	router.HandleFunc("/items", getItems).Methods("GET")
	router.HandleFunc("/items/{id}", getItem).Methods("GET")
	router.HandleFunc("/items", createItem).Methods("POST")
	router.HandleFunc("/items/{id}", updateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")

	// Start the server
	fmt.Println("Server running on port 8000")
	http.ListenAndServe(":8000", router)
}
