package cmd

import (
	"fmt"
	"os"
	"project-app-inventaris-cli-azwin/handler"
)

func Home(handlerCategory handler.HandlerCategory, handlerManagement handler.HandlerManagement, handlerOld handler.HandlerOld, handlerReport handler.HandlerReport) {
	for {
		fmt.Println("\n=== MENU UTAMA ===")
		fmt.Println("1. Fitur Kategori Barang")
		fmt.Println("2. Fitur Manajemen Barang Inventaris")
		fmt.Println("3. Barang yang Perlu Diganti (> 100 hari)")
		fmt.Println("4. Laporan Investasi dan Depresiasi")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih menu: ")

		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Input tidak valid. Harap masukkan angka.")
			continue
		}
		
		switch choice {
		case 1:
			CategorySubmenu(handlerCategory)
		case 2:
			ManagementSubmenu(handlerManagement, handlerCategory)
		case 3:
			OldItems(handlerOld)
		case 4:
			ReportSubmenu(handlerReport)
		case 5:
			fmt.Println("Keluar dari aplikasi...")
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih antara 1 - 5.")
		}
	}

}
