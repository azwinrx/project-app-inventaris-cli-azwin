package repository

import (
	"context"
	"project-app-inventaris-cli-azwin/database"
	"project-app-inventaris-cli-azwin/model"
)

type RepositoryOldInterface interface {
	GetOldItems() ([]model.Management, error)
}

type RepositoryOld struct {
	DB database.PgxIface
}

func NewRepositoryOld(db database.PgxIface) RepositoryOld {
	return RepositoryOld{
		DB: db,
	}
}

// Fungsi untuk mendapatkan item yang telah digunakan lebih dari 100 hari
func (repo *RepositoryOld) GetOldItems() ([]model.Management, error) {
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
		WHERE CURRENT_DATE - i.purchase_date > 100
		ORDER BY usage_days DESC`

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

