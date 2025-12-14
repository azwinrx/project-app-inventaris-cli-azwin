package model

import "time"

type Management struct {
	ID           int       `json:"id"`
	CategoryId   int       `json:"category_id"`
	CategoryName string    `json:"category_name"`
	Name         string    `json:"name"`
	Price        float64   `json:"price"`
	PurchaseDate time.Time `json:"purchase_date"`
	UsageDays    int       `json:"usage_days"`
	CreatedAt    string    `json:"created_at"`
	UpdatedAt    string    `json:"updated_at"`
}