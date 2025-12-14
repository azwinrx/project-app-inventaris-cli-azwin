package repository

import (
	"errors"
	"testing"
	"time"
)

// Test GetOldItems - Sukses
func TestGetOldItems_Success(t *testing.T) {
	purchaseDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	
	mockDB := &MockDB{
		queryResult: &MockRows{
			data: [][]interface{}{
				{1, 1, "Electronics", "Old Laptop", 15000000.0, purchaseDate, 150},
				{2, 2, "Furniture", "Old Chair", 2000000.0, purchaseDate, 200},
			},
		},
	}

	repo := NewRepositoryOld(mockDB)
	items, err := repo.GetOldItems()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(items) != 2 {
		t.Errorf("Expected 2 items, got %d", len(items))
	}

	if items[0].UsageDays != 150 {
		t.Errorf("Expected first item usage days to be 150, got %d", items[0].UsageDays)
	}
}

// Test GetOldItems - Error
func TestGetOldItems_Error(t *testing.T) {
	mockDB := &MockDB{
		queryError: errors.New("database error"),
	}

	repo := NewRepositoryOld(mockDB)
	_, err := repo.GetOldItems()

	if err == nil {
		t.Error("Expected an error, got nil")
	}
}

// Test GetOldItems - Hasil Kosong
func TestGetOldItems_EmptyResult(t *testing.T) {
	mockDB := &MockDB{
		queryResult: &MockRows{
			data: [][]interface{}{},
		},
	}

	repo := NewRepositoryOld(mockDB)
	items, err := repo.GetOldItems()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(items) != 0 {
		t.Errorf("Expected 0 items, got %d", len(items))
	}
}

// Test GetOldItems - Error Scan
func TestGetOldItems_ScanError(t *testing.T) {
	mockDB := &MockDB{
		queryResult: &MockRows{
			data: [][]interface{}{
				{1, 1, "Electronics", "Old Laptop", 15000000.0, time.Now(), 150},
			},
			err: errors.New("scan error"),
		},
	}

	repo := NewRepositoryOld(mockDB)
	_, err := repo.GetOldItems()

	if err == nil {
		t.Error("Expected an error, got nil")
	}
}
