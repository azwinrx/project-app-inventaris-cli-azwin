package repository

import (
	"errors"
	"testing"
	"time"
)

// Test GetAllItemsForReport - Sukses
func TestGetAllItemsForReport_Success(t *testing.T) {
	purchaseDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	
	mockDB := &MockDB{
		queryResult: &MockRows{
			data: [][]interface{}{
				{1, 1, "Electronics", "Laptop", 15000000.0, purchaseDate, 30},
				{2, 2, "Furniture", "Chair", 2000000.0, purchaseDate, 60},
			},
		},
	}

	repo := NewRepositoryReport(mockDB)
	items, err := repo.GetAllItemsForReport()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(items) != 2 {
		t.Errorf("Expected 2 items, got %d", len(items))
	}

	if items[0].Name != "Laptop" {
		t.Errorf("Expected first item name to be 'Laptop', got '%s'", items[0].Name)
	}

	if items[0].Price != 15000000.0 {
		t.Errorf("Expected first item price to be 15000000.0, got %f", items[0].Price)
	}

	if items[0].UsageDays != 30 {
		t.Errorf("Expected first item usage days to be 30, got %d", items[0].UsageDays)
	}
}

// Test GetAllItemsForReport - Error
func TestGetAllItemsForReport_Error(t *testing.T) {
	mockDB := &MockDB{
		queryError: errors.New("database error"),
	}

	repo := NewRepositoryReport(mockDB)
	_, err := repo.GetAllItemsForReport()

	if err == nil {
		t.Error("Expected an error, got nil")
	}

	if err.Error() != "database error" {
		t.Errorf("Expected error message 'database error', got '%s'", err.Error())
	}
}

// Test GetAllItemsForReport - Hasil Kosong
func TestGetAllItemsForReport_EmptyResult(t *testing.T) {
	mockDB := &MockDB{
		queryResult: &MockRows{
			data: [][]interface{}{},
		},
	}

	repo := NewRepositoryReport(mockDB)
	items, err := repo.GetAllItemsForReport()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(items) != 0 {
		t.Errorf("Expected 0 items, got %d", len(items))
	}
}

// Test GetAllItemsForReport - Error Scan
func TestGetAllItemsForReport_ScanError(t *testing.T) {
	mockDB := &MockDB{
		queryResult: &MockRows{
			data: [][]interface{}{
				{1, 1, "Electronics", "Laptop", 15000000.0, time.Now(), 30},
			},
			err: errors.New("scan error"),
		},
	}

	repo := NewRepositoryReport(mockDB)
	_, err := repo.GetAllItemsForReport()

	if err == nil {
		t.Error("Expected an error, got nil")
	}
}

// Test GetItemByIdForReport - Sukses
func TestGetItemByIdForReport_Success(t *testing.T) {
	purchaseDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	
	mockDB := &MockDB{
		queryRowResult: &MockRow{
			data: []interface{}{1, 1, "Electronics", "Laptop", 15000000.0, purchaseDate, 30},
			err:  nil,
		},
	}

	repo := NewRepositoryReport(mockDB)
	item, err := repo.GetItemByIdForReport(1)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if item.ID != 1 {
		t.Errorf("Expected item ID to be 1, got %d", item.ID)
	}

	if item.Name != "Laptop" {
		t.Errorf("Expected item name to be 'Laptop', got '%s'", item.Name)
	}

	if item.CategoryName != "Electronics" {
		t.Errorf("Expected category name to be 'Electronics', got '%s'", item.CategoryName)
	}

	if item.Price != 15000000.0 {
		t.Errorf("Expected item price to be 15000000.0, got %f", item.Price)
	}

	if item.UsageDays != 30 {
		t.Errorf("Expected usage days to be 30, got %d", item.UsageDays)
	}
}

// Test GetItemByIdForReport - Tidak Ditemukan
func TestGetItemByIdForReport_NotFound(t *testing.T) {
	mockDB := &MockDB{
		queryRowResult: &MockRow{
			err: errors.New("no rows in result set"),
		},
	}

	repo := NewRepositoryReport(mockDB)
	_, err := repo.GetItemByIdForReport(999)

	if err == nil {
		t.Error("Expected an error for non-existent item, got nil")
	}

	if err.Error() != "no rows in result set" {
		t.Errorf("Expected error message 'no rows in result set', got '%s'", err.Error())
	}
}

// Test GetItemByIdForReport - Error Scan
func TestGetItemByIdForReport_ScanError(t *testing.T) {
	mockDB := &MockDB{
		queryRowResult: &MockRow{
			data: []interface{}{1, 1, "Electronics"}, // incomplete data
			err:  errors.New("scan error"),
		},
	}

	repo := NewRepositoryReport(mockDB)
	_, err := repo.GetItemByIdForReport(1)

	if err == nil {
		t.Error("Expected an error, got nil")
	}
}

// Test GetItemByIdForReport - Kategori Berbeda
func TestGetItemByIdForReport_DifferentCategory(t *testing.T) {
	purchaseDate := time.Date(2023, 6, 15, 0, 0, 0, 0, time.UTC)
	
	mockDB := &MockDB{
		queryRowResult: &MockRow{
			data: []interface{}{5, 3, "Furniture", "Office Desk", 5000000.0, purchaseDate, 150},
			err:  nil,
		},
	}

	repo := NewRepositoryReport(mockDB)
	item, err := repo.GetItemByIdForReport(5)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if item.CategoryName != "Furniture" {
		t.Errorf("Expected category name to be 'Furniture', got '%s'", item.CategoryName)
	}

	if item.Name != "Office Desk" {
		t.Errorf("Expected item name to be 'Office Desk', got '%s'", item.Name)
	}

	if item.UsageDays != 150 {
		t.Errorf("Expected usage days to be 150, got %d", item.UsageDays)
	}
}
