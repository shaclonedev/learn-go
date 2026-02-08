package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"learn-go/internal/model"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) GetAll(name string) ([]model.Product, error) {
	query := "SELECT products.id, products.name, products.price, products.stock, products.description, categories.name as category_name, categories.id as category_id FROM products left join categories on products.category_id = categories.id"

	args := []interface{}{}
	if name != "" {
		query += " where name like $1"
		args = append(args, "%"+name+"%")
	}

	rows, err := repo.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]model.Product, 0)
	for rows.Next() {
		var p model.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.Description, &p.CategoryName, &p.CategoryID)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func (repo *ProductRepository) Create(product *model.Product) (int, error) {
	query := "INSERT INTO products (name, price, stock, description, category_id) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	err := repo.db.QueryRow(query, product.Name, product.Price, product.Stock, product.Description, product.CategoryID).Scan(&product.ID)
	fmt.Println("ID", product.ID)
	return product.ID, err
}

// GetByID - ambil produk by ID
func (repo *ProductRepository) GetByID(id int) (*model.Product, error) {
	query := "SELECT products.id, products.name, products.price, products.stock, products.description, categories.name as category_name, categories.id as category_id FROM products left join categories on products.category_id = categories.id WHERE products.id = $1"

	var p model.Product
	err := repo.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.Description, &p.CategoryName, &p.CategoryID)
	if err == sql.ErrNoRows {
		return nil, errors.New("product not found")
	}
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (repo *ProductRepository) Update(product *model.Product) error {
	query := "UPDATE products SET name = $1, price = $2, stock = $3, description = $4 WHERE id = $5"
	result, err := repo.db.Exec(query, product.Name, product.Price, product.Stock, product.Description, product.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("product not found")
	}

	return nil
}

func (repo *ProductRepository) Delete(id int) error {
	query := "DELETE FROM products WHERE id = $1"
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("product not found")
	}

	return err
}
