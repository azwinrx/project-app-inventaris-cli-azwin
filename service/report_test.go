package service

import (
	"errors"
	"testing"
	"time"

	"project-app-inventaris-cli-azwin/model"
)

// Mock Repository for Report Service
type MockRepoReportService struct {
	GetAllItemsForReportFunc func() ([]model.Management, error)
	GetItemByIdForReportFunc func(id int) (model.Management, error)
}

func (m *MockRepoReportService) GetAllItemsForReport() ([]model.Management, error) {
	if m.GetAllItemsForReportFunc != nil {
		return m.GetAllItemsForReportFunc()
	}
	return nil, nil
}

func (m *MockRepoReportService) GetItemByIdForReport(id int) (model.Management, error) {
	if m.GetItemByIdForReportFunc != nil {
		return m.GetItemByIdForReportFunc(id)
	}
	return model.Management{}, nil
}

// Test GetInvestmentReport - Success
func TestServiceGetInvestmentReport_Success(t *testing.T) {
	mockRepo := &MockRepoReportService{
		GetAllItemsForReportFunc: func() ([]model.Management, error) {
			return []model.Management{
				{ID: 1, CategoryName: "Electronics", Name: "Laptop", Price: 15000000, UsageDays: 365},
				{ID: 2, CategoryName: "Furniture", Name: "Chair", Price: 2000000, UsageDays: 730},
			}, nil
		},
	}

	service := NewServiceReport(mockRepo)
	items, totalInvestment, totalCurrentValue, err := service.GetInvestmentReport()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(items) != 2 {
		t.Errorf("Expected 2 items, got %d", len(items))
	}

	if totalInvestment != 17000000 {
		t.Errorf("Expected total investment 17000000, got %f", totalInvestment)
	}

	if totalCurrentValue <= 0 {
		t.Errorf("Expected positive current value, got %f", totalCurrentValue)
	}
}

// Test GetInvestmentReport - Error
func TestServiceGetInvestmentReport_Error(t *testing.T) {
	mockRepo := &MockRepoReportService{
		GetAllItemsForReportFunc: func() ([]model.Management, error) {
			return nil, errors.New("database error")
		},
	}

	service := NewServiceReport(mockRepo)
	_, _, _, err := service.GetInvestmentReport()

	if err == nil {
		t.Error("Expected an error, got nil")
	}
}

// Test GetInvestmentReport - Empty Items
func TestServiceGetInvestmentReport_EmptyItems(t *testing.T) {
	mockRepo := &MockRepoReportService{
		GetAllItemsForReportFunc: func() ([]model.Management, error) {
			return []model.Management{}, nil
		},
	}

	service := NewServiceReport(mockRepo)
	items, totalInvestment, totalCurrentValue, err := service.GetInvestmentReport()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(items) != 0 {
		t.Errorf("Expected 0 items, got %d", len(items))
	}

	if totalInvestment != 0 {
		t.Errorf("Expected total investment 0, got %f", totalInvestment)
	}

	if totalCurrentValue != 0 {
		t.Errorf("Expected current value 0, got %f", totalCurrentValue)
	}
}

// Test GetItemDepreciationReport - Success
func TestServiceGetItemDepreciationReport_Success(t *testing.T) {
	mockRepo := &MockRepoReportService{
		GetItemByIdForReportFunc: func(id int) (model.Management, error) {
			return model.Management{
				ID:           1,
				CategoryName: "Electronics",
				Name:         "Laptop",
				Price:        15000000,
				PurchaseDate: time.Now().AddDate(0, 0, -365),
				UsageDays:    365,
			}, nil
		},
	}

	service := NewServiceReport(mockRepo)
	item, currentValue, depreciation, err := service.GetItemDepreciationReport(1)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if item.Name != "Laptop" {
		t.Errorf("Expected item name 'Laptop', got '%s'", item.Name)
	}

	if currentValue <= 0 {
		t.Errorf("Expected positive current value, got %f", currentValue)
	}

	if depreciation <= 0 {
		t.Errorf("Expected positive depreciation, got %f", depreciation)
	}
}

// Test GetItemDepreciationReport - Error
func TestServiceGetItemDepreciationReport_Error(t *testing.T) {
	mockRepo := &MockRepoReportService{
		GetItemByIdForReportFunc: func(id int) (model.Management, error) {
			return model.Management{}, errors.New("item not found")
		},
	}

	service := NewServiceReport(mockRepo)
	_, _, _, err := service.GetItemDepreciationReport(999)

	if err == nil {
		t.Error("Expected an error, got nil")
	}
}
