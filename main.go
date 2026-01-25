package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Category struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var id int64 = 3
var categories []Category = []Category{
	{ID: 1, Name: "A", Description: "DESC A"},
	{ID: 2, Name: "B", Description: "DESC B"},
	{ID: 3, Name: "C", Description: "DESC C"},
}

func getAll(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

func create(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type must be json!", http.StatusUnsupportedMediaType)
		return
	}

	var newCat Category
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&newCat)
	if err != nil {
		http.Error(w, "Invalid data: "+err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if newCat.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	id++
	newCat.ID = id
	categories = append(categories, newCat)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCat)
	log.Printf("Created product: %+v\n", newCat)
}

func categoriesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getAll(w)

	case http.MethodPost:
		create(w, r)

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func categoryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := parseCategoryID(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
	}

	for _, c := range categories {
		if c.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(c)
			return
		}
	}
}

func parseCategoryID(path string) (int64, error) {
	idStr := strings.TrimPrefix(path, "/categories/")
	return strconv.ParseInt(idStr, 10, 64)
}

func main() {
	http.HandleFunc("/categories", categoriesHandler)
	http.HandleFunc("/categories/", categoryHandler)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
