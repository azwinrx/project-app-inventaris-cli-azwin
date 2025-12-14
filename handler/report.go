package handler

import (
	"fmt"
	"math"
	"os"
	"project-app-inventaris-cli-azwin/service"
	"text/tabwriter"
)

type HandlerReport struct {
	ServiceReport service.ServiceReportInterface
}

func NewHandlerReport(serviceReport service.ServiceReportInterface) HandlerReport {
	return HandlerReport{
		ServiceReport: serviceReport,
	}
}

// Calculate depreciation using declining balance method (20% per year)
func calculateDepreciationDisplay(price float64, usageDays int) (float64, float64) {
	years := float64(usageDays) / 365.0
	currentValue := price * math.Pow(0.80, years)
	depreciation := price - currentValue
	return currentValue, depreciation
}

func (handlerReport *HandlerReport) GetInvestmentReport() {
	items, totalInvestment, totalCurrentValue, err := handlerReport.ServiceReport.GetInvestmentReport()
	if err != nil {
		fmt.Println("Gagal mengambil laporan investasi:", err)
		return
	}

	fmt.Println("\n=== LAPORAN INVESTASI DAN DEPRESIASI ===")
	fmt.Println("Metode Depresiasi: Saldo Menurun 20% per tahun")
	
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, "ID\tKategori\tNama\tHarga Awal\tHari Pakai\tNilai Saat Ini\tDepresiasi")
	fmt.Fprintln(w, "--\t--------\t----\t----------\t----------\t--------------\t----------")
	
	for _, item := range items {
		currentValue, depreciation := calculateDepreciationDisplay(item.Price, item.UsageDays)
		fmt.Fprintf(w, "%d\t%s\t%s\tRp %.2f\t%d\tRp %.2f\tRp %.2f\n", 
			item.ID, 
			item.CategoryName, 
			item.Name, 
			item.Price,
			item.UsageDays,
			currentValue,
			depreciation)
	}
	w.Flush()
	
	fmt.Println("\n" + string(make([]byte, 100)))
	fmt.Printf("TOTAL INVESTASI AWAL         : Rp%.2f\n", totalInvestment)
	fmt.Printf("TOTAL NILAI SAAT INI         : Rp%.2f\n", totalCurrentValue)
	fmt.Printf("TOTAL DEPRESIASI             : Rp%.2f\n", totalInvestment - totalCurrentValue)
	fmt.Printf("PERSENTASE NILAI TERSISA     : %.2f%%\n", (totalCurrentValue/totalInvestment)*100)
}

func (handlerReport *HandlerReport) GetItemDepreciationReport(id int) {
	item, currentValue, depreciation, err := handlerReport.ServiceReport.GetItemDepreciationReport(id)
	if err != nil {
		fmt.Println("Gagal mengambil laporan barang:", err)
		return
	}

	years := float64(item.UsageDays) / 365.0

	fmt.Println("\n=== LAPORAN DEPRESIASI BARANG ===")
	fmt.Printf("ID Barang          : %d\n", item.ID)
	fmt.Printf("Nama Barang        : %s\n", item.Name)
	fmt.Printf("Kategori           : %s\n", item.CategoryName)
	fmt.Printf("Tanggal Pembelian  : %s\n", item.PurchaseDate.Format("2006-01-02"))
	fmt.Printf("Hari Penggunaan    : %d hari (%.2f tahun)\n", item.UsageDays, years)
	fmt.Println("\n--- PERHITUNGAN DEPRESIASI ---")
	fmt.Println("Metode             : Saldo Menurun")
	fmt.Println("Rate Depresiasi    : 20% per tahun")
	fmt.Printf("Harga Awal         : Rp%.2f\n", item.Price)
	fmt.Printf("Nilai Saat Ini     : Rp%.2f\n", currentValue)
	fmt.Printf("Total Depresiasi   : Rp%.2f\n", depreciation)
	fmt.Printf("Persentase Tersisa : %.2f%%\n", (currentValue/item.Price)*100)
	fmt.Println("\nFormula: Nilai = Harga Awal × (0.80)^tahun")
	fmt.Printf("Perhitungan: %.2f × (0.80)^%.2f = %.2f\n", item.Price, years, currentValue)
}
