package service

import (
	"errors"
	"testing"

	"project-app-inventaris-cli-azwin/model"
)

// Mock Repository for Old Service
type MockRepoOldService struct {
	GetOldItemsFunc func() ([]model.Management, error)
}

func (m *MockRepoOldService) GetOldItems() ([]model.Management, error) {
	if m.GetOldItemsFunc != nil {
		return m.GetOldItemsFunc()
	}
	return nil, nil
}

// Test GetOldItems - Success
func TestServiceGetOldItems_Success(t *testing.T) {
	mockRepo := &MockRepoOldService{
		GetOldItemsFunc: func() ([]model.Management, error) {
			return []model.Management{
				{ID: 1, CategoryName: "Electronics", Name: "Old Laptop", Price: 15000000, UsageDays: 150},
				{ID: 2, CategoryName: "Furniture", Name: "Old Chair", Price: 2000000, UsageDays: 200},
			}, nil
		},
	}

	service := NewServiceOld(mockRepo)
	items, err := service.GetOldItems()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(items) != 2 {
		t.Errorf("Expected 2 items, got %d", len(items))
	}
}

// Test GetOldItems - Error
func TestServiceGetOldItems_Error(t *testing.T) {
	mockRepo := &MockRepoOldService{
		GetOldItemsFunc: func() ([]model.Management, error) {
			return nil, errors.New("database error")
		},
	}

	service := NewServiceOld(mockRepo)
	_, err := service.GetOldItems()

	if err == nil {
		t.Error("Expected an error, got nil")
	}
}

// Test GetOldItems - Empty Result
func TestServiceGetOldItems_EmptyResult(t *testing.T) {
	mockRepo := &MockRepoOldService{
		GetOldItemsFunc: func() ([]model.Management, error) {
			return []model.Management{}, nil
		},
	}

	service := NewServiceOld(mockRepo)
	items, err := service.GetOldItems()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(items) != 0 {
		t.Errorf("Expected 0 items, got %d", len(items))
	}
}
