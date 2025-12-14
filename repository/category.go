package repository

import (
	"context"
	"errors"
	"fmt"
	"project-app-inventaris-cli-azwin/database"
	"project-app-inventaris-cli-azwin/model"
)

type RepositoryCategoryInterface interface {
	GetCategory() ([]model.Category, error)
	AddCategory(name, description string) error
	GetCategoryById(id int) (model.Category, error)
	UpdateCategory(id int, name, description string) error
	DeleteCategory(id int) error
}

type RepositoryCateogry struct {
	DB database.PgxIface
}

func NewrepoCategory(db database.PgxIface) RepositoryCateogry {
	return RepositoryCateogry{
		DB: db,
	}
}

// Fungsi untuk melihat semua kategori
func (repo *RepositoryCateogry) GetCategory() ([]model.Category, error) {
	query := `SELECT id, name, description FROM categories ORDER BY id ASC;`

	rows, err := repo.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []model.Category
	var category model.Category
	for rows.Next() {
		err := rows.Scan(&category.ID, &category.Name, &category.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}


// Fungsi untuk menambahkan kategori baru
func (repo *RepositoryCateogry) AddCategory(name, description string) error {
	// Validasi duplikasi nama kategori
	checkQuery := `SELECT COUNT(*) FROM categories WHERE LOWER(name) = LOWER($1)`
	var count int
	err := repo.DB.QueryRow(context.Background(), checkQuery, name).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New(fmt.Sprintf("kategori dengan nama '%s' sudah ada", name))
	}

	query := `INSERT INTO categories (name, description, created_at, updated_at) 
			VALUES ($1, $2, NOW(), NOW())`

	_, err = repo.DB.Exec(context.Background(), query, name, description)
	if err != nil {
		return err
	}

	return nil
}

// Fungsi untuk melihat kategori berdasarkan id
func (repo *RepositoryCateogry) GetCategoryById(id int) (model.Category, error) {
	query := `SELECT id, name, description FROM categories WHERE id = $1`

	var category model.Category
	err := repo.DB.QueryRow(context.Background(), query, id).Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		return model.Category{}, err
	}

	return category, nil
}

// Fungsi untuk memperbarui kategori
func (repo *RepositoryCateogry) UpdateCategory(id int, name, description string) error {
	// Validasi duplikasi nama kategori (kecuali ID yang sama)
	checkQuery := `SELECT COUNT(*) FROM categories WHERE LOWER(name) = LOWER($1) AND id != $2`
	var count int
	err := repo.DB.QueryRow(context.Background(), checkQuery, name, id).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New(fmt.Sprintf("kategori dengan nama '%s' sudah ada", name))
	}

	query := `UPDATE categories SET name = $1, description = $2, updated_at = NOW() WHERE id = $3`

	_, err = repo.DB.Exec(context.Background(), query, name, description, id)
	if err != nil {
		return err
	}

	return nil
}

// Fungsi untuk menghapus kategori
func (repo *RepositoryCateogry) DeleteCategory(id int) error {
	query := `DELETE FROM categories WHERE id = $1`

	_, err := repo.DB.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}

	return nil
}