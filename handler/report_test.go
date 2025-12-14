package handler

import (
	"errors"
	"testing"

	"project-app-inventaris-cli-azwin/model"
)

// Mock Service for Report Handler
type MockServiceReport struct {
	GetInvestmentReportFunc       func() ([]model.Management, float64, float64, error)
	GetItemDepreciationReportFunc func(id int) (model.Management, float64, float64, error)
}

func (m *MockServiceReport) GetInvestmentReport() ([]model.Management, float64, float64, error) {
	if m.GetInvestmentReportFunc != nil {
		return m.GetInvestmentReportFunc()
	}
	return nil, 0, 0, nil
}

func (m *MockServiceReport) GetItemDepreciationReport(id int) (model.Management, float64, float64, error) {
	if m.GetItemDepreciationReportFunc != nil {
		return m.GetItemDepreciationReportFunc(id)
	}
	return model.Management{}, 0, 0, nil
}

// Test GetInvestmentReport - Success
func TestHandlerGetInvestmentReport_Success(t *testing.T) {
	mockService := &MockServiceReport{
		GetInvestmentReportFunc: func() ([]model.Management, float64, float64, error) {
			return []model.Management{
				{ID: 1, CategoryName: "Electronics", Name: "Laptop", Price: 15000000, UsageDays: 30},
			}, 15000000.0, 14000000.0, nil
		},
	}

	handler := NewHandlerReport(mockService)
	output := captureOutput(func() {
		handler.GetInvestmentReport()
	})

	if output == "" {
		t.Error("Expected output, got empty string")
	}
}

// Test GetInvestmentReport - Error
func TestHandlerGetInvestmentReport_Error(t *testing.T) {
	mockService := &MockServiceReport{
		GetInvestmentReportFunc: func() ([]model.Management, float64, float64, error) {
			return nil, 0, 0, errors.New("database error")
		},
	}

	handler := NewHandlerReport(mockService)
	output := captureOutput(func() {
		handler.GetInvestmentReport()
	})

	if output == "" {
		t.Error("Expected error message in output")
	}
}

// Test GetItemDepreciationReport - Success
func TestHandlerGetItemDepreciationReport_Success(t *testing.T) {
	mockService := &MockServiceReport{
		GetItemDepreciationReportFunc: func(id int) (model.Management, float64, float64, error) {
			return model.Management{
				ID:           1,
				CategoryName: "Electronics",
				Name:         "Laptop",
				Price:        15000000,
				UsageDays:    30,
			}, 14000000.0, 1000000.0, nil
		},
	}

	handler := NewHandlerReport(mockService)
	output := captureOutput(func() {
		handler.GetItemDepreciationReport(1)
	})

	if output == "" {
		t.Error("Expected output, got empty string")
	}
}

// Test GetItemDepreciationReport - Error
func TestHandlerGetItemDepreciationReport_Error(t *testing.T) {
	mockService := &MockServiceReport{
		GetItemDepreciationReportFunc: func(id int) (model.Management, float64, float64, error) {
			return model.Management{}, 0, 0, errors.New("item not found")
		},
	}

	handler := NewHandlerReport(mockService)
	output := captureOutput(func() {
		handler.GetItemDepreciationReport(1)
	})

	if output == "" {
		t.Error("Expected error message in output")
	}
}
