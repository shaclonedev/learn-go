package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
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
func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "health!")
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
	log.Printf("Created category: %+v\n", newCat)
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

func updateCat(w http.ResponseWriter, r *http.Request) {
	id, err := parseCategoryID(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type must be json!", http.StatusUnsupportedMediaType)
		return
	}

	var updatedCat Category
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err = decoder.Decode(&updatedCat)
	if err != nil {
		http.Error(w, "Invalid data: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if updatedCat.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	found := false
	for i, cat := range categories {
		if cat.ID == id {
			updatedCat.ID = id
			categories[i] = updatedCat
			found = true
			break
		}
	}

	if !found {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedCat)

	log.Printf("Updated category: %+v\n", updatedCat)
}

func getCat(w http.ResponseWriter, r *http.Request) {
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

func deleteCat(w http.ResponseWriter, r *http.Request) {
	id, err := parseCategoryID(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
	}

	foundIndex := -1
	for i, cat := range categories {
		if cat.ID == id {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	categories = append(categories[:foundIndex], categories[foundIndex+1:]...)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "Category deleted successfully")
	log.Printf("Deleted category with ID: %d\n", id)
}

func categoryHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getCat(w, r)

	case http.MethodPut:
		updateCat(w, r)

	case http.MethodDelete:
		deleteCat(w, r)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

func parseCategoryID(path string) (int64, error) {
	idStr := strings.TrimPrefix(path, "/categories/")
	return strconv.ParseInt(idStr, 10, 64)
}

func main() {
	http.HandleFunc("/categories", categoriesHandler)
	http.HandleFunc("/categories/", categoryHandler)
	http.HandleFunc("/health", healthCheck)

	http.ListenAndServe("0.0.0.0:"+os.Getenv("PORT"), nil)
}
