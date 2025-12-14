package service

import (
	"errors"
	"testing"

	"project-app-inventaris-cli-azwin/model"
)

// Mock Repository for Category Service
type MockRepoCategoryService struct {
	GetCategoryFunc      func() ([]model.Category, error)
	AddCategoryFunc      func(name, description string) error
	GetCategoryByIdFunc  func(id int) (model.Category, error)
	UpdateCategoryFunc   func(id int, name, description string) error
	DeleteCategoryFunc   func(id int) error
}

func (m *MockRepoCategoryService) GetCategory() ([]model.Category, error) {
	if m.GetCategoryFunc != nil {
		return m.GetCategoryFunc()
	}
	return nil, nil
}

func (m *MockRepoCategoryService) AddCategory(name, description string) error {
	if m.AddCategoryFunc != nil {
		return m.AddCategoryFunc(name, description)
	}
	return nil
}

func (m *MockRepoCategoryService) GetCategoryById(id int) (model.Category, error) {
	if m.GetCategoryByIdFunc != nil {
		return m.GetCategoryByIdFunc(id)
	}
	return model.Category{}, nil
}

func (m *MockRepoCategoryService) UpdateCategory(id int, name, description string) error {
	if m.UpdateCategoryFunc != nil {
		return m.UpdateCategoryFunc(id, name, description)
	}
	return nil
}

func (m *MockRepoCategoryService) DeleteCategory(id int) error {
	if m.DeleteCategoryFunc != nil {
		return m.DeleteCategoryFunc(id)
	}
	return nil
}

// Test GetCategory - Success
func TestServiceGetCategory_Success(t *testing.T) {
	mockRepo := &MockRepoCategoryService{
		GetCategoryFunc: func() ([]model.Category, error) {
			return []model.Category{
				{ID: 1, Name: "Electronics", Description: "Electronic items"},
				{ID: 2, Name: "Furniture", Description: "Office furniture"},
			}, nil
		},
	}

	service := NewServiceCategory(mockRepo)
	categories, err := service.GetCategory()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(categories) != 2 {
		t.Errorf("Expected 2 categories, got %d", len(categories))
	}
}

// Test GetCategory - Error
func TestServiceGetCategory_Error(t *testing.T) {
	mockRepo := &MockRepoCategoryService{
		GetCategoryFunc: func() ([]model.Category, error) {
			return nil, errors.New("database error")
		},
	}

	service := NewServiceCategory(mockRepo)
	_, err := service.GetCategory()

	if err == nil {
		t.Error("Expected an error, got nil")
	}
}

// Test AddCategory - Success
func TestServiceAddCategory_Success(t *testing.T) {
	mockRepo := &MockRepoCategoryService{
		AddCategoryFunc: func(name, description string) error {
			return nil
		},
	}

	service := NewServiceCategory(mockRepo)
	err := service.AddCategory("New Category", "New Description")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

// Test AddCategory - Error
func TestServiceAddCategory_Error(t *testing.T) {
	mockRepo := &MockRepoCategoryService{
		AddCategoryFunc: func(name, description string) error {
			return errors.New("duplicate category")
		},
	}

	service := NewServiceCategory(mockRepo)
	err := service.AddCategory("Existing Category", "Description")

	if err == nil {
		t.Error("Expected an error, got nil")
	}
}

// Test GetCategoryById - Success
func TestServiceGetCategoryById_Success(t *testing.T) {
	mockRepo := &MockRepoCategoryService{
		GetCategoryByIdFunc: func(id int) (model.Category, error) {
			return model.Category{ID: 1, Name: "Electronics", Description: "Electronic items"}, nil
		},
	}

	service := NewServiceCategory(mockRepo)
	category, err := service.GetCategoryById(1)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if category.Name != "Electronics" {
		t.Errorf("Expected category name to be 'Electronics', got '%s'", category.Name)
	}
}

// Test GetCategoryById - Error
func TestServiceGetCategoryById_Error(t *testing.T) {
	mockRepo := &MockRepoCategoryService{
		GetCategoryByIdFunc: func(id int) (model.Category, error) {
			return model.Category{}, errors.New("not found")
		},
	}

	service := NewServiceCategory(mockRepo)
	_, err := service.GetCategoryById(999)

	if err == nil {
		t.Error("Expected an error, got nil")
	}
}

// Test UpdateCategory - Success
func TestServiceUpdateCategory_Success(t *testing.T) {
	mockRepo := &MockRepoCategoryService{
		UpdateCategoryFunc: func(id int, name, description string) error {
			return nil
		},
	}

	service := NewServiceCategory(mockRepo)
	err := service.UpdateCategory(1, "Updated Category", "Updated Description")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

// Test UpdateCategory - Error
func TestServiceUpdateCategory_Error(t *testing.T) {
	mockRepo := &MockRepoCategoryService{
		UpdateCategoryFunc: func(id int, name, description string) error {
			return errors.New("update failed")
		},
	}

	service := NewServiceCategory(mockRepo)
	err := service.UpdateCategory(1, "Category", "Description")

	if err == nil {
		t.Error("Expected an error, got nil")
	}
}

// Test DeleteCategory - Success
func TestServiceDeleteCategory_Success(t *testing.T) {
	mockRepo := &MockRepoCategoryService{
		DeleteCategoryFunc: func(id int) error {
			return nil
		},
	}

	service := NewServiceCategory(mockRepo)
	err := service.DeleteCategory(1)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

// Test DeleteCategory - Error
func TestServiceDeleteCategory_Error(t *testing.T) {
	mockRepo := &MockRepoCategoryService{
		DeleteCategoryFunc: func(id int) error {
			return errors.New("delete failed")
		},
	}

	service := NewServiceCategory(mockRepo)
	err := service.DeleteCategory(1)

	if err == nil {
		t.Error("Expected an error, got nil")
	}
}
