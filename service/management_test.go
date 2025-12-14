package service

import (
	"errors"
	"testing"
	"time"

	"project-app-inventaris-cli-azwin/model"
)

// Mock Repository untuk Service Management
type MockRepoManagementService struct {
	GetAllItemsFunc       func() ([]model.Management, error)
	AddItemFunc           func(categoryId int, name string, price float64, purchaseDate string) error
	GetItemByIdFunc       func(id int) (model.Management, error)
	UpdateItemFunc        func(id int, categoryId int, name string, price float64, purchaseDate string) error
	DeleteItemFunc        func(id int) error
	SearchItemsByNameFunc func(keyword string) ([]model.Management, error)
}

func (m *MockRepoManagementService) GetAllItems() ([]model.Management, error) {
	if m.GetAllItemsFunc != nil {
		return m.GetAllItemsFunc()
	}
	return nil, nil
}

func (m *MockRepoManagementService) AddItem(categoryId int, name string, price float64, purchaseDate string) error {
	if m.AddItemFunc != nil {
		return m.AddItemFunc(categoryId, name, price, purchaseDate)
	}
	return nil
}

func (m *MockRepoManagementService) GetItemById(id int) (model.Management, error) {
	if m.GetItemByIdFunc != nil {
		return m.GetItemByIdFunc(id)
	}
	return model.Management{}, nil
}

func (m *MockRepoManagementService) UpdateItem(id int, categoryId int, name string, price float64, purchaseDate string) error {
	if m.UpdateItemFunc != nil {
		return m.UpdateItemFunc(id, categoryId, name, price, purchaseDate)
	}
	return nil
}

func (m *MockRepoManagementService) DeleteItem(id int) error {
	if m.DeleteItemFunc != nil {
		return m.DeleteItemFunc(id)
	}
	return nil
}

func (m *MockRepoManagementService) SearchItemsByName(keyword string) ([]model.Management, error) {
	if m.SearchItemsByNameFunc != nil {
		return m.SearchItemsByNameFunc(keyword)
	}
	return nil, nil
}

// Test GetAllItems - Sukses
func TestServiceGetAllItems_Success(t *testing.T) {
	mockRepo := &MockRepoManagementService{
		GetAllItemsFunc: func() ([]model.Management, error) {
			return []model.Management{
				{ID: 1, CategoryId: 1, CategoryName: "Electronics", Name: "Laptop", Price: 15000000, PurchaseDate: time.Now(), UsageDays: 30},
				{ID: 2, CategoryId: 2, CategoryName: "Furniture", Name: "Chair", Price: 2000000, PurchaseDate: time.Now(), UsageDays: 60},
			}, nil
		},
	}

	service := NewServiceManagement(mockRepo)
	items, err := service.GetAllItems()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(items) != 2 {
		t.Errorf("Expected 2 items, got %d", len(items))
	}
}

// Test GetAllItems - Error
func TestServiceGetAllItems_Error(t *testing.T) {
	mockRepo := &MockRepoManagementService{
		GetAllItemsFunc: func() ([]model.Management, error) {
			return nil, errors.New("database error")
		},
	}

	service := NewServiceManagement(mockRepo)
	_, err := service.GetAllItems()

	if err == nil {
		t.Error("Expected an error, got nil")
	}
}

// Test AddItem - Sukses
func TestServiceAddItem_Success(t *testing.T) {
	mockRepo := &MockRepoManagementService{
		AddItemFunc: func(categoryId int, name string, price float64, purchaseDate string) error {
			return nil
		},
	}

	service := NewServiceManagement(mockRepo)
	err := service.AddItem(1, "New Item", 10000, "2024-01-01")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

// Test AddItem - Error
func TestServiceAddItem_Error(t *testing.T) {
	mockRepo := &MockRepoManagementService{
		AddItemFunc: func(categoryId int, name string, price float64, purchaseDate string) error {
			return errors.New("duplicate item")
		},
	}

	service := NewServiceManagement(mockRepo)
	err := service.AddItem(1, "Existing Item", 10000, "2024-01-01")

	if err == nil {
		t.Error("Expected an error, got nil")
	}
}

// Test GetItemById - Sukses
func TestServiceGetItemById_Success(t *testing.T) {
	mockRepo := &MockRepoManagementService{
		GetItemByIdFunc: func(id int) (model.Management, error) {
			return model.Management{ID: 1, CategoryId: 1, CategoryName: "Electronics", Name: "Laptop", Price: 15000000}, nil
		},
	}

	service := NewServiceManagement(mockRepo)
	item, err := service.GetItemById(1)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if item.Name != "Laptop" {
		t.Errorf("Expected item name to be 'Laptop', got '%s'", item.Name)
	}
}

// Test GetItemById - Error
func TestServiceGetItemById_Error(t *testing.T) {
	mockRepo := &MockRepoManagementService{
		GetItemByIdFunc: func(id int) (model.Management, error) {
			return model.Management{}, errors.New("not found")
		},
	}

	service := NewServiceManagement(mockRepo)
	_, err := service.GetItemById(999)

	if err == nil {
		t.Error("Expected an error, got nil")
	}
}

// Test UpdateItem - Sukses
func TestServiceUpdateItem_Success(t *testing.T) {
	mockRepo := &MockRepoManagementService{
		UpdateItemFunc: func(id int, categoryId int, name string, price float64, purchaseDate string) error {
			return nil
		},
	}

	service := NewServiceManagement(mockRepo)
	err := service.UpdateItem(1, 1, "Updated Item", 20000, "2024-01-01")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

// Test UpdateItem - Error
func TestServiceUpdateItem_Error(t *testing.T) {
	mockRepo := &MockRepoManagementService{
		UpdateItemFunc: func(id int, categoryId int, name string, price float64, purchaseDate string) error {
			return errors.New("update failed")
		},
	}

	service := NewServiceManagement(mockRepo)
	err := service.UpdateItem(1, 1, "Item", 20000, "2024-01-01")

	if err == nil {
		t.Error("Expected an error, got nil")
	}
}

// Test DeleteItem - Sukses
func TestServiceDeleteItem_Success(t *testing.T) {
	mockRepo := &MockRepoManagementService{
		DeleteItemFunc: func(id int) error {
			return nil
		},
	}

	service := NewServiceManagement(mockRepo)
	err := service.DeleteItem(1)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

// Test DeleteItem - Error
func TestServiceDeleteItem_Error(t *testing.T) {
	mockRepo := &MockRepoManagementService{
		DeleteItemFunc: func(id int) error {
			return errors.New("delete failed")
		},
	}

	service := NewServiceManagement(mockRepo)
	err := service.DeleteItem(1)

	if err == nil {
		t.Error("Expected an error, got nil")
	}
}

// Test SearchItemsByName - Sukses
func TestServiceSearchItemsByName_Success(t *testing.T) {
	mockRepo := &MockRepoManagementService{
		SearchItemsByNameFunc: func(keyword string) ([]model.Management, error) {
			return []model.Management{
				{ID: 1, CategoryId: 1, CategoryName: "Electronics", Name: "Laptop Dell", Price: 15000000},
			}, nil
		},
	}

	service := NewServiceManagement(mockRepo)
	items, err := service.SearchItemsByName("Laptop")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(items) != 1 {
		t.Errorf("Expected 1 item, got %d", len(items))
	}

	if items[0].Name != "Laptop Dell" {
		t.Errorf("Expected item name to be 'Laptop Dell', got '%s'", items[0].Name)
	}
}

// Test SearchItemsByName - Error
func TestServiceSearchItemsByName_Error(t *testing.T) {
	mockRepo := &MockRepoManagementService{
		SearchItemsByNameFunc: func(keyword string) ([]model.Management, error) {
			return nil, errors.New("search failed")
		},
	}

	service := NewServiceManagement(mockRepo)
	_, err := service.SearchItemsByName("Laptop")

	if err == nil {
		t.Error("Expected an error, got nil")
	}
}

// Test SearchItemsByName - Hasil Kosong
func TestServiceSearchItemsByName_EmptyResult(t *testing.T) {
	mockRepo := &MockRepoManagementService{
		SearchItemsByNameFunc: func(keyword string) ([]model.Management, error) {
			return []model.Management{}, nil
		},
	}

	service := NewServiceManagement(mockRepo)
	items, err := service.SearchItemsByName("NonExistent")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(items) != 0 {
		t.Errorf("Expected 0 items, got %d", len(items))
	}
}
