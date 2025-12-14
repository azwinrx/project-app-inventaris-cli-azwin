package repository

import (
	"context"
	"errors"
	"fmt"
	"project-app-inventaris-cli-azwin/database"
	"project-app-inventaris-cli-azwin/model"
)

type RepositoryManagementInterface interface {
	GetAllItems() ([]model.Management, error)
	AddItem(categoryId int, name string, price float64, purchaseDate string) error
	GetItemById(id int) (model.Management, error)
	UpdateItem(id int, categoryId int, name string, price float64, purchaseDate string) error
	DeleteItem(id int) error
	SearchItemsByName(keyword string) ([]model.Management, error)
}

type RepositoryManagement struct {
	DB database.PgxIface
}

func NewrepoManagement(db database.PgxIface) RepositoryManagement {
	return RepositoryManagement{
		DB: db,
	}
}

// Fungsi untuk mendapatkan semua item dengan nama kategori dan hari pemakaian
func (repo *RepositoryManagement) GetAllItems() ([]model.Management, error) {
	query := `
		SELECT 
			i.id, 
			i.category_id, 
			c.name as category_name, 
			i.name, 
			i.price, 
			i.purchase_date,
			CURRENT_DATE - i.purchase_date as usage_days
		FROM items i
		JOIN categories c ON i.category_id = c.id
		ORDER BY i.id ASC`

	rows, err := repo.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []model.Management
	for rows.Next() {
		var item model.Management
		err := rows.Scan(
			&item.ID,
			&item.CategoryId,
			&item.CategoryName,
			&item.Name,
			&item.Price,
			&item.PurchaseDate,
			&item.UsageDays,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

// Fungsi untuk menambahkan item baru
func (repo *RepositoryManagement) AddItem(categoryId int, name string, price float64, purchaseDate string) error {
	// Validasi duplikasi nama barang dalam kategori yang sama
	checkQuery := `SELECT COUNT(*) FROM items WHERE LOWER(name) = LOWER($1) AND category_id = $2`
	var count int
	err := repo.DB.QueryRow(context.Background(), checkQuery, name, categoryId).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New(fmt.Sprintf("barang dengan nama '%s' sudah ada dalam kategori ini", name))
	}

	// Validasi harga tidak boleh negatif atau nol
	if price <= 0 {
		return errors.New("harga barang harus lebih dari 0")
	}

	query := `INSERT INTO items (category_id, name, price, purchase_date, created_at, updated_at) 
			VALUES ($1, $2, $3, $4, NOW(), NOW())`

	_, err = repo.DB.Exec(context.Background(), query, categoryId, name, price, purchaseDate)
	if err != nil {
		return err
	}

	return nil
}

// Fungsi untuk mendapatkan item berdasarkan id
func (repo *RepositoryManagement) GetItemById(id int) (model.Management, error) {
	query := `
		SELECT 
			i.id, 
			i.category_id, 
			c.name as category_name, 
			i.name, 
			i.price, 
			i.purchase_date,
			CURRENT_DATE - i.purchase_date as usage_days
		FROM items i
		JOIN categories c ON i.category_id = c.id
		WHERE i.id = $1`

	var item model.Management
	err := repo.DB.QueryRow(context.Background(), query, id).Scan(
		&item.ID,
		&item.CategoryId,
		&item.CategoryName,
		&item.Name,
		&item.Price,
		&item.PurchaseDate,
		&item.UsageDays,
	)
	if err != nil {
		return model.Management{}, err
	}

	return item, nil
}

// Fungsi untuk memperbarui item
func (repo *RepositoryManagement) UpdateItem(id int, categoryId int, name string, price float64, purchaseDate string) error {
	// Validasi duplikasi nama barang dalam kategori yang sama (kecuali ID yang sama)
	checkQuery := `SELECT COUNT(*) FROM items WHERE LOWER(name) = LOWER($1) AND category_id = $2 AND id != $3`
	var count int
	err := repo.DB.QueryRow(context.Background(), checkQuery, name, categoryId, id).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New(fmt.Sprintf("barang dengan nama '%s' sudah ada dalam kategori ini", name))
	}

	// Validasi harga tidak boleh negatif atau nol
	if price <= 0 {
		return errors.New("harga barang harus lebih dari 0")
	}

	query := `UPDATE items 
			SET category_id = $1, name = $2, price = $3, purchase_date = $4, updated_at = NOW() 
			WHERE id = $5`

	_, err = repo.DB.Exec(context.Background(), query, categoryId, name, price, purchaseDate, id)
	if err != nil {
		return err
	}

	return nil
}

// Fungsi untuk menghapus item
func (repo *RepositoryManagement) DeleteItem(id int) error {
	query := `DELETE FROM items WHERE id = $1`

	_, err := repo.DB.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}

	return nil
}

// Fungsi untuk mencari item berdasarkan kata kunci nama
func (repo *RepositoryManagement) SearchItemsByName(keyword string) ([]model.Management, error) {
	query := `
		SELECT 
			i.id, 
			i.category_id, 
			c.name as category_name, 
			i.name, 
			i.price, 
			i.purchase_date,
			CURRENT_DATE - i.purchase_date as usage_days
		FROM items i
		JOIN categories c ON i.category_id = c.id
		WHERE LOWER(i.name) LIKE LOWER($1)
		ORDER BY i.id ASC`

	rows, err := repo.DB.Query(context.Background(), query, "%"+keyword+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []model.Management
	for rows.Next() {
		var item model.Management
		err := rows.Scan(
			&item.ID,
			&item.CategoryId,
			&item.CategoryName,
			&item.Name,
			&item.Price,
			&item.PurchaseDate,
			&item.UsageDays,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

