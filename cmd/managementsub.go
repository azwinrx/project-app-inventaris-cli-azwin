package cmd

import (
	"fmt"
	"project-app-inventaris-cli-azwin/handler"
)

func ManagementSubmenu(handlerManagement handler.HandlerManagement, handlerCategory handler.HandlerCategory) {
	for {
		fmt.Println("\n=== MENU MANAJEMEN INVENTARIS ===")
		fmt.Println("1. Lihat Semua Barang")
		fmt.Println("2. Tambah Barang")
		fmt.Println("3. Lihat Detail Barang")
		fmt.Println("4. Update Barang")
		fmt.Println("5. Hapus Barang")
		fmt.Println("6. Cari Barang (Berdasarkan Nama)")
		fmt.Println("7. Kembali ke Menu Utama")
		fmt.Print("Pilih menu: ")

		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Input tidak valid. Harap masukkan angka.")
			continue
		}
		
		switch choice {
		case 1:
			ManagementMenu(handlerManagement)
		case 2:
			AddItem(handlerManagement, handlerCategory)
		case 3:
			GetItemById(handlerManagement)
		case 4:
			UpdateItem(handlerManagement, handlerCategory)
		case 5:
			DeleteItem(handlerManagement)
		case 6:
			SearchItems(handlerManagement)
		case 7:
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih antara 1 - 7.")
		}
	}
}
