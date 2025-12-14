package service

import (
	"math"
	"project-app-inventaris-cli-azwin/model"
	"project-app-inventaris-cli-azwin/repository"
)

type ServiceReportInterface interface {
	GetInvestmentReport() ([]model.Management, float64, float64, error)
	GetItemDepreciationReport(id int) (model.Management, float64, float64, error)
}

type ServiceReport struct {
	RepoReport repository.RepositoryReportInterface
}

func NewServiceReport(repoReport repository.RepositoryReportInterface) ServiceReport {
	return ServiceReport{
		RepoReport: repoReport,
	}
}

// Hitung depresiasi menggunakan metode saldo menurun (20% per tahun)
func calculateDepreciation(price float64, usageDays int) float64 {
	// Hitung tahun (termasuk pecahan tahun)
	years := float64(usageDays) / 365.0
	
	// Saldo menurun: Nilai Sekarang = Nilai Awal × (1 - tingkat)^tahun
	// Tingkat = 20% = 0.20
	currentValue := price * math.Pow(0.80, years)
	
	return currentValue
}

// Dapatkan laporan total investasi
func (serviceReport *ServiceReport) GetInvestmentReport() ([]model.Management, float64, float64, error) {
	items, err := serviceReport.RepoReport.GetAllItemsForReport()
	if err != nil {
		return nil, 0, 0, err
	}

	var totalInvestment float64
	var totalCurrentValue float64

	for i := range items {
		currentValue := calculateDepreciation(items[i].Price, items[i].UsageDays)
		totalInvestment += items[i].Price
		totalCurrentValue += currentValue
	}

	return items, totalInvestment, totalCurrentValue, nil
}

// Dapatkan laporan depresiasi untuk item tertentu
func (serviceReport *ServiceReport) GetItemDepreciationReport(id int) (model.Management, float64, float64, error) {
	item, err := serviceReport.RepoReport.GetItemByIdForReport(id)
	if err != nil {
		return model.Management{}, 0, 0, err
	}

	currentValue := calculateDepreciation(item.Price, item.UsageDays)
	depreciation := item.Price - currentValue

	return item, currentValue, depreciation, nil
}
