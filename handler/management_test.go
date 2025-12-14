package handler

import (
	"bytes"
	"errors"
	"io"
	"os"
	"testing"

	"project-app-inventaris-cli-azwin/model"
)

// Mock Service for Handler Testing
type MockServiceManagement struct {
	GetAllItemsFunc       func() ([]model.Management, error)
	AddItemFunc           func(categoryId int, name string, price float64, purchaseDate string) error
	GetItemByIdFunc       func(id int) (model.Management, error)
	UpdateItemFunc        func(id int, categoryId int, name string, price float64, purchaseDate string) error
	DeleteItemFunc        func(id int) error
	SearchItemsByNameFunc func(keyword string) ([]model.Management, error)
}

func (m *MockServiceManagement) GetAllItems() ([]model.Management, error) {
	if m.GetAllItemsFunc != nil {
		return m.GetAllItemsFunc()
	}
	return nil, nil
}

func (m *MockServiceManagement) AddItem(categoryId int, name string, price float64, purchaseDate string) error {
	if m.AddItemFunc != nil {
		return m.AddItemFunc(categoryId, name, price, purchaseDate)
	}
	return nil
}

func (m *MockServiceManagement) GetItemById(id int) (model.Management, error) {
	if m.GetItemByIdFunc != nil {
		return m.GetItemByIdFunc(id)
	}
	return model.Management{}, nil
}

func (m *MockServiceManagement) UpdateItem(id int, categoryId int, name string, price float64, purchaseDate string) error {
	if m.UpdateItemFunc != nil {
		return m.UpdateItemFunc(id, categoryId, name, price, purchaseDate)
	}
	return nil
}

func (m *MockServiceManagement) DeleteItem(id int) error {
	if m.DeleteItemFunc != nil {
		return m.DeleteItemFunc(id)
	}
	return nil
}

func (m *MockServiceManagement) SearchItemsByName(keyword string) ([]model.Management, error) {
	if m.SearchItemsByNameFunc != nil {
		return m.SearchItemsByNameFunc(keyword)
	}
	return nil, nil
}

// Capture stdout for testing
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

// Test GetAllItems - Success
func TestHandlerGetAllItems_Success(t *testing.T) {
	mockService := &MockServiceManagement{
		GetAllItemsFunc: func() ([]model.Management, error) {
			return []model.Management{
				{ID: 1, CategoryName: "Electronics", Name: "Laptop", Price: 15000000, UsageDays: 30},
			}, nil
		},
	}

	handler := NewHandlerManagement(mockService)
	output := captureOutput(func() {
		handler.GetAllItems()
	})

	if output == "" {
		t.Error("Expected output, got empty string")
	}
}

// Test GetAllItems - Error
func TestHandlerGetAllItems_Error(t *testing.T) {
	mockService := &MockServiceManagement{
		GetAllItemsFunc: func() ([]model.Management, error) {
			return nil, errors.New("database error")
		},
	}

	handler := NewHandlerManagement(mockService)
	output := captureOutput(func() {
		handler.GetAllItems()
	})

	if output == "" {
		t.Error("Expected error message in output")
	}
}

// Test AddItem - Success
func TestHandlerAddItem_Success(t *testing.T) {
	mockService := &MockServiceManagement{
		AddItemFunc: func(categoryId int, name string, price float64, purchaseDate string) error {
			return nil
		},
	}

	handler := NewHandlerManagement(mockService)
	output := captureOutput(func() {
		handler.AddItem(1, "New Item", 10000, "2024-01-01")
	})

	if output == "" {
		t.Error("Expected success message in output")
	}
}

// Test AddItem - Error
func TestHandlerAddItem_Error(t *testing.T) {
	mockService := &MockServiceManagement{
		AddItemFunc: func(categoryId int, name string, price float64, purchaseDate string) error {
			return errors.New("add failed")
		},
	}

	handler := NewHandlerManagement(mockService)
	output := captureOutput(func() {
		handler.AddItem(1, "New Item", 10000, "2024-01-01")
	})

	if output == "" {
		t.Error("Expected error message in output")
	}
}

// Test SearchItems - Success
func TestHandlerSearchItems_Success(t *testing.T) {
	mockService := &MockServiceManagement{
		SearchItemsByNameFunc: func(keyword string) ([]model.Management, error) {
			return []model.Management{
				{ID: 1, CategoryName: "Electronics", Name: "Laptop Dell", Price: 15000000, UsageDays: 30},
			}, nil
		},
	}

	handler := NewHandlerManagement(mockService)
	output := captureOutput(func() {
		handler.SearchItems("Laptop")
	})

	if output == "" {
		t.Error("Expected output, got empty string")
	}
}

// Test SearchItems - No Results
func TestHandlerSearchItems_NoResults(t *testing.T) {
	mockService := &MockServiceManagement{
		SearchItemsByNameFunc: func(keyword string) ([]model.Management, error) {
			return []model.Management{}, nil
		},
	}

	handler := NewHandlerManagement(mockService)
	output := captureOutput(func() {
		handler.SearchItems("NonExistent")
	})

	if output == "" {
		t.Error("Expected no results message in output")
	}
}

// Test SearchItems - Error
func TestHandlerSearchItems_Error(t *testing.T) {
	mockService := &MockServiceManagement{
		SearchItemsByNameFunc: func(keyword string) ([]model.Management, error) {
			return nil, errors.New("search failed")
		},
	}

	handler := NewHandlerManagement(mockService)
	output := captureOutput(func() {
		handler.SearchItems("Laptop")
	})

	if output == "" {
		t.Error("Expected error message in output")
	}
}
