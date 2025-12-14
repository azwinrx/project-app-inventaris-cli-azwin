package handler

import (
	"errors"
	"testing"

	"project-app-inventaris-cli-azwin/model"
)

// Mock Service for Category Handler
type MockServiceCategory struct {
	GetCategoryFunc     func() ([]model.Category, error)
	AddCategoryFunc     func(name, description string) error
	GetCategoryByIdFunc func(id int) (model.Category, error)
	UpdateCategoryFunc  func(id int, name, description string) error
	DeleteCategoryFunc  func(id int) error
}

func (m *MockServiceCategory) GetCategory() ([]model.Category, error) {
	if m.GetCategoryFunc != nil {
		return m.GetCategoryFunc()
	}
	return nil, nil
}

func (m *MockServiceCategory) AddCategory(name, description string) error {
	if m.AddCategoryFunc != nil {
		return m.AddCategoryFunc(name, description)
	}
	return nil
}

func (m *MockServiceCategory) GetCategoryById(id int) (model.Category, error) {
	if m.GetCategoryByIdFunc != nil {
		return m.GetCategoryByIdFunc(id)
	}
	return model.Category{}, nil
}

func (m *MockServiceCategory) UpdateCategory(id int, name, description string) error {
	if m.UpdateCategoryFunc != nil {
		return m.UpdateCategoryFunc(id, name, description)
	}
	return nil
}

func (m *MockServiceCategory) DeleteCategory(id int) error {
	if m.DeleteCategoryFunc != nil {
		return m.DeleteCategoryFunc(id)
	}
	return nil
}

// Test GetCategory - Success
func TestHandlerGetCategory_Success(t *testing.T) {
	mockService := &MockServiceCategory{
		GetCategoryFunc: func() ([]model.Category, error) {
			return []model.Category{
				{ID: 1, Name: "Electronics", Description: "Electronic items"},
			}, nil
		},
	}

	handler := NewHandlerCategory(mockService)
	output := captureOutput(func() {
		handler.GetCategory()
	})

	if output == "" {
		t.Error("Expected output, got empty string")
	}
}

// Test GetCategory - Error
func TestHandlerGetCategory_Error(t *testing.T) {
	mockService := &MockServiceCategory{
		GetCategoryFunc: func() ([]model.Category, error) {
			return nil, errors.New("database error")
		},
	}

	handler := NewHandlerCategory(mockService)
	output := captureOutput(func() {
		handler.GetCategory()
	})

	if output == "" {
		t.Error("Expected error message in output")
	}
}

// Test AddCategory - Success
func TestHandlerAddCategory_Success(t *testing.T) {
	mockService := &MockServiceCategory{
		AddCategoryFunc: func(name, description string) error {
			return nil
		},
	}

	handler := NewHandlerCategory(mockService)
	output := captureOutput(func() {
		handler.AddCategory("New Category", "Description")
	})

	if output == "" {
		t.Error("Expected success message in output")
	}
}

// Test AddCategory - Error
func TestHandlerAddCategory_Error(t *testing.T) {
	mockService := &MockServiceCategory{
		AddCategoryFunc: func(name, description string) error {
			return errors.New("duplicate category")
		},
	}

	handler := NewHandlerCategory(mockService)
	output := captureOutput(func() {
		handler.AddCategory("Existing Category", "Description")
	})

	if output == "" {
		t.Error("Expected error message in output")
	}
}
