package repository

import (
	"database/sql"
	"errors"
	"learn-go/internal/model"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetAll() ([]model.Category, error) {
	rows, err := r.db.Query("SELECT id, name, description FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []model.Category
	for rows.Next() {
		var category model.Category
		if err := rows.Scan(&category.ID, &category.Name, &category.Description); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (r *CategoryRepository) Create(category *model.Category) error {
	err := r.db.QueryRow("INSERT INTO categories (name, description) VALUES ($1, $2) RETURNING id", category.Name, category.Description).Scan(&category.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *CategoryRepository) GetByID(id int) (*model.Category, error) {
	var category model.Category
	err := r.db.QueryRow("SELECT id, name, description FROM categories WHERE id = $1", id).Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("category not found")
		}
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) Update(category *model.Category) error {
	result, err := r.db.Exec("UPDATE categories SET name = $1, description = $2 WHERE id = $3", category.Name, category.Description, category.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("category not found")
	}
	return nil
}

func (r *CategoryRepository) Delete(id int) error {
	result, err := r.db.Exec("DELETE FROM categories WHERE id = $1", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("category not found")
	}
	return nil
}
