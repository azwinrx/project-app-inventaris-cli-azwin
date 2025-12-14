package cmd

import (
	"fmt"
	"project-app-inventaris-cli-azwin/handler"
)

func ReportSubmenu(handlerReport handler.HandlerReport) {
	for {
		fmt.Println("\n=== MENU LAPORAN INVESTASI & DEPRESIASI ===")
		fmt.Println("1. Laporan Total Investasi dan Depresiasi")
		fmt.Println("2. Laporan Depresiasi Barang Tertentu")
		fmt.Println("3. Kembali ke Menu Utama")
		fmt.Print("Pilih menu: ")

		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Input tidak valid. Harap masukkan angka.")
			continue
		}
		
		switch choice {
		case 1:
			InvestmentReport(handlerReport)
		case 2:
			ItemDepreciationReport(handlerReport)
		case 3:
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih antara 1 - 3.")
		}
	}
}
