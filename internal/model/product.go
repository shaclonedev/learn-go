package model

type Product struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Price        int    `json:"price"`
	Stock        int    `json:"stock"`
	CategoryName string `json:"category_name"`
}
