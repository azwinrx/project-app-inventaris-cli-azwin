package cmd

import (
	"bufio"
	"fmt"
	"os"
	"project-app-inventaris-cli-azwin/handler"
	"strings"
)

func CategorySubmenu(handlerCategory handler.HandlerCategory) {
	for {
		fmt.Println("\n=== MENU KATEGORI ===")
		fmt.Println("1. Lihat Semua Kategori")
		fmt.Println("2. Tambah Kategori")
		fmt.Println("3. Lihat Detail Kategori")
		fmt.Println("4. Update Kategori")
		fmt.Println("5. Hapus Kategori")
		fmt.Println("6. Kembali ke Menu Utama")
		fmt.Print("Pilih menu: ")

		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Input tidak valid. Harap masukkan angka.")
			continue
		}
		
		switch choice {
		case 1:
			ViewCategory(handlerCategory)
		case 2:
			AddCategory(handlerCategory)
		case 3:
			ViewCategoryById(handlerCategory)
		case 4:
			UpdateCategory(handlerCategory)
		case 5:
			DeleteCategory(handlerCategory)
		case 6:
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih antara 1 - 6.")
		}
	}
}

func ViewCategory(handler handler.HandlerCategory) {
	fmt.Println("\n=== LIHAT KATEGORI ===")


	handler.GetCategory()

	var choice string
	println("")
	fmt.Print("Apakah kamu ingin melanjutkan ke halaman lain (ya/tidak)? ")
	fmt.Scanln(&choice)

	switch choice {
	case "ya":
		ClearScreen()
	case "tidak":
		os.Exit(0)
	default:
		fmt.Println("Pilihan kamu salah, tolong masukkan ulang")
	}
}

func AddCategory(handler handler.HandlerCategory) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\n=== TAMBAH KATEGORI ===")
	
	// Input category name
	fmt.Print("Nama Kategori: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	
	// Input description
	fmt.Print("Deskripsi: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)
	
	// Validation
	if name == "" {
		fmt.Println("Nama kategori tidak boleh kosong!")
		return
	}
	
	// Call handler to add category
	handler.AddCategory(name, description)
}

func ViewCategoryById(handler handler.HandlerCategory) {
	fmt.Println("\n=== DETAIL KATEGORI ===")
	
	// Input category ID
	var id int
	fmt.Print("Masukkan ID Kategori: ")
	_, err := fmt.Scanln(&id)

	// Validation
	if err != nil {
		fmt.Println("Input tidak valid. Harap masukkan angka.")
		return
	}
	
	// Call handler to get category by id
	handler.GetCategoryById(id)
}

func UpdateCategory(handler handler.HandlerCategory) {
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Println("\n=== UPDATE KATEGORI ===")
	
	var id int
	fmt.Print("Masukkan ID Kategori: ")
	_, err := fmt.Scanln(&id)
	if err != nil {
		fmt.Println("Input tidak valid. Harap masukkan angka.")
		return
	}
	
	// Input new category name
	fmt.Print("Nama Kategori Baru: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	
	// Input new description
	fmt.Print("Deskripsi Baru: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)
	
	// Validation
	if name == "" {
		fmt.Println("Nama kategori tidak boleh kosong!")
		return
	}
	
	// Call handler to update category
	handler.UpdateCategory(id, name, description)
}

func DeleteCategory(handler handler.HandlerCategory) {
	fmt.Println("\n=== HAPUS KATEGORI ===")
	
	var id int
	fmt.Print("Masukkan ID Kategori yang akan dihapus: ")
	_, err := fmt.Scanln(&id)
	if err != nil {
		fmt.Println("Input tidak valid. Harap masukkan angka.")
		return
	}
	
	// Confirmation
	var confirm string
	fmt.Print("Apakah Anda yakin ingin menghapus kategori ini? (ya/tidak): ")
	fmt.Scanln(&confirm)
	
	if confirm != "ya" {
		fmt.Println("Penghapusan dibatalkan.")
		return
	}
	
	// Call handler to delete category
	handler.DeleteCategory(id)
}
