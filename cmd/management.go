package cmd

import (
	"bufio"
	"fmt"
	"os"
	"project-app-inventaris-cli-azwin/handler"
	"strings"
)

func ManagementMenu(handlerManagement handler.HandlerManagement) {
	var choice string

	println("")
	handlerManagement.GetAllItems()
	println("")

	fmt.Print("Apakah kamu ingin melanjutkan ke halaman lain (ya/tidak)? ")
	fmt.Scanln(&choice)

	switch choice {
	case "ya":
		ClearScreen()
	case "tidak":
		os.Exit(0)
	default:
		fmt.Println("Pilihan kamu salah tolong masukkan ulang")
	}
}

func AddItem(handlerManagement handler.HandlerManagement, handlerCategory handler.HandlerCategory) {
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Println("\n=== TAMBAH BARANG INVENTARIS ===")
	
	// Show available categories
	fmt.Println("\nKategori yang Tersedia:")
	handlerCategory.GetCategory()
	
	// Input category ID
	var categoryId int
	fmt.Print("\nID Kategori: ")
	_, err := fmt.Scanln(&categoryId)
	if err != nil {
		fmt.Println("Input tidak valid. Harap masukkan angka.")
		return
	}
	
	// Input item name
	fmt.Print("Nama Barang: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	
	// Input price
	var price float64
	fmt.Print("Harga: ")
	_, err = fmt.Scanln(&price)
	if err != nil {
		fmt.Println("Input tidak valid. Harap masukkan harga yang benar.")
		return
	}
	
	// Input purchase date
	fmt.Print("Tanggal Pembelian (YYYY-MM-DD): ")
	purchaseDate, _ := reader.ReadString('\n')
	purchaseDate = strings.TrimSpace(purchaseDate)
	
	// Validation
	if name == "" {
		fmt.Println("Nama barang tidak boleh kosong!")
		return
	}
	
	// Call handler to add item
	handlerManagement.AddItem(categoryId, name, price, purchaseDate)
}

func GetItemById(handlerManagement handler.HandlerManagement) {
	fmt.Println("\n=== DETAIL BARANG ===")
	
	var id int
	fmt.Print("Masukkan ID Barang: ")
	_, err := fmt.Scanln(&id)
	if err != nil {
		fmt.Println("Input tidak valid. Harap masukkan angka.")
		return
	}
	
	handlerManagement.GetItemById(id)
}

func UpdateItem(handlerManagement handler.HandlerManagement, handlerCategory handler.HandlerCategory) {
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Println("\n=== UPDATE BARANG ===")
	
	var id int
	fmt.Print("Masukkan ID Barang: ")
	_, err := fmt.Scanln(&id)
	if err != nil {
		fmt.Println("Input tidak valid. Harap masukkan angka.")
		return
	}
	
	// Show available categories
	fmt.Println("\nKategori yang Tersedia:")
	handlerCategory.GetCategory()
	
	// Input new category ID
	var categoryId int
	fmt.Print("\nID Kategori Baru: ")
	_, err = fmt.Scanln(&categoryId)
	if err != nil {
		fmt.Println("Input tidak valid. Harap masukkan angka.")
		return
	}
	
	// Input new item name
	fmt.Print("Nama Barang Baru: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	
	// Input new price
	var price float64
	fmt.Print("Harga Baru: ")
	_, err = fmt.Scanln(&price)
	if err != nil {
		fmt.Println("Input tidak valid. Harap masukkan harga yang benar.")
		return
	}
	
	// Input new purchase date
	fmt.Print("Tanggal Pembelian Baru (YYYY-MM-DD): ")
	purchaseDate, _ := reader.ReadString('\n')
	purchaseDate = strings.TrimSpace(purchaseDate)
	
	// Validation
	if name == "" {
		fmt.Println("Nama barang tidak boleh kosong!")
		return
	}
	
	// Call handler to update item
	handlerManagement.UpdateItem(id, categoryId, name, price, purchaseDate)
}

func DeleteItem(handlerManagement handler.HandlerManagement) {
	fmt.Println("\n=== HAPUS BARANG ===")
	
	var id int
	fmt.Print("Masukkan ID Barang yang akan dihapus: ")
	_, err := fmt.Scanln(&id)
	if err != nil {
		fmt.Println("Input tidak valid. Harap masukkan angka.")
		return
	}
	
	// Confirmation
	var confirm string
	fmt.Print("Apakah Anda yakin ingin menghapus barang ini? (ya/tidak): ")
	fmt.Scanln(&confirm)
	
	if confirm != "ya" {
		fmt.Println("Penghapusan dibatalkan.")
		return
	}
	
	// Call handler to delete item
	handlerManagement.DeleteItem(id)
}

func OldItems(handlerOld handler.HandlerOld) {
	handlerOld.GetOldItems()
	
	var choice string
	println("")
	fmt.Print("Tekan Enter untuk melanjutkan...")
	fmt.Scanln(&choice)
}

func InvestmentReport(handlerReport handler.HandlerReport) {
	handlerReport.GetInvestmentReport()
	
	var choice string
	println("")
	fmt.Print("Tekan Enter untuk melanjutkan...")
	fmt.Scanln(&choice)
}

func ItemDepreciationReport(handlerReport handler.HandlerReport) {
	fmt.Println("\n=== LAPORAN DEPRESIASI BARANG ===")
	
	var id int
	fmt.Print("Masukkan ID Barang: ")
	_, err := fmt.Scanln(&id)
	if err != nil {
		fmt.Println("Input tidak valid. Harap masukkan angka.")
		return
	}
	
	handlerReport.GetItemDepreciationReport(id)
	
	var choice string
	println("")
	fmt.Print("Tekan Enter untuk melanjutkan...")
	fmt.Scanln(&choice)
}

func SearchItems(handlerManagement handler.HandlerManagement) {
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Println("\n=== PENCARIAN BARANG ===")
	fmt.Print("Masukkan kata kunci pencarian: ")
	keyword, _ := reader.ReadString('\n')
	keyword = strings.TrimSpace(keyword)
	
	if keyword == "" {
		fmt.Println("Kata kunci tidak boleh kosong!")
		return
	}
	
	handlerManagement.SearchItems(keyword)
	
	var choice string
	println("")
	fmt.Print("Tekan Enter untuk melanjutkan...")
	fmt.Scanln(&choice)
}