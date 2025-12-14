package cmd

import (
	"fmt"
	"os"
	"project-app-inventaris-cli-azwin/handler"
)

func Category(handler handler.HandlerCategory) {
	fmt.Println("\n=== VIEW CATEGORY MENU ===")
		fmt.Println("1. View Categories")
		fmt.Println("2. Add Category")
		fmt.Println("3. View Category Detail")
		fmt.Println("4. Update Category")
		fmt.Println("5. Delete Category")
		fmt.Println("6. Exit")
		fmt.Print("Choose an option: ")
	var choice int
	fmt.Print("Masukkan Menu yang kamu pilih: ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		ViewCategory(handler)
	case 2:
		AddCategory(handler)	
	case 3:
		ViewCategoryById(handler)
	case 4:
		UpdateCategory(handler)
	case 5:
		DeleteCategory(handler)
	case 6:
		os.Exit(0)
	default:
		fmt.Println("Pilihan kamu salah tolong masukkan ulang")
	}
}