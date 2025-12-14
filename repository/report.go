package repository

import (
	"context"
	"project-app-inventaris-cli-azwin/database"
	"project-app-inventaris-cli-azwin/model"
)

type RepositoryReportInterface interface {
	GetAllItemsForReport() ([]model.Management, error)
	GetItemByIdForReport(id int) (model.Management, error)
}

type RepositoryReport struct {
	DB database.PgxIface
}

func NewRepositoryReport(db database.PgxIface) RepositoryReport {
	return RepositoryReport{
		DB: db,
	}
}

// Function to get all items with usage days for depreciation calculation
func (repo *RepositoryReport) GetAllItemsForReport() ([]model.Management, error) {
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

// Function to get item by id for depreciation report
func (repo *RepositoryReport) GetItemByIdForReport(id int) (model.Management, error) {
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
