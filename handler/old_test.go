package handler

import (
	"errors"
	"testing"

	"project-app-inventaris-cli-azwin/model"
)

// Mock Service untuk Handler Old
type MockServiceOld struct {
	GetOldItemsFunc func() ([]model.Management, error)
}

func (m *MockServiceOld) GetOldItems() ([]model.Management, error) {
	if m.GetOldItemsFunc != nil {
		return m.GetOldItemsFunc()
	}
	return nil, nil
}

// Test GetOldItems - Sukses
func TestHandlerGetOldItems_Success(t *testing.T) {
	mockService := &MockServiceOld{
		GetOldItemsFunc: func() ([]model.Management, error) {
			return []model.Management{
				{ID: 1, CategoryName: "Electronics", Name: "Old Laptop", Price: 15000000, UsageDays: 150},
			}, nil
		},
	}

	handler := NewHandlerOld(mockService)
	output := captureOutput(func() {
		handler.GetOldItems()
	})

	if output == "" {
		t.Error("Expected output, got empty string")
	}
}

// Test GetOldItems - Error
func TestHandlerGetOldItems_Error(t *testing.T) {
	mockService := &MockServiceOld{
		GetOldItemsFunc: func() ([]model.Management, error) {
			return nil, errors.New("database error")
		},
	}

	handler := NewHandlerOld(mockService)
	output := captureOutput(func() {
		handler.GetOldItems()
	})

	if output == "" {
		t.Error("Expected error message in output")
	}
}

// Test GetOldItems - Tidak Ada Item
func TestHandlerGetOldItems_NoItems(t *testing.T) {
	mockService := &MockServiceOld{
		GetOldItemsFunc: func() ([]model.Management, error) {
			return []model.Management{}, nil
		},
	}

	handler := NewHandlerOld(mockService)
	output := captureOutput(func() {
		handler.GetOldItems()
	})

	if output == "" {
		t.Error("Expected output for no items, got empty string")
	}
}
