package cmd

import (
	"fmt"
	"project-app-inventaris-cli-azwin/handler"

	"github.com/spf13/cobra"
)

// Global variables untuk handlers (akan diinject dari main)
var (
	handlerCategory   *handler.HandlerCategory
	handlerManagement *handler.HandlerManagement
	handlerOld        *handler.HandlerOld
	handlerReport     *handler.HandlerReport
)

// Flags untuk kategori
var (
	categoryID          int
	categoryName        string
	categoryDescription string
)

// Flags untuk item inventaris
var (
	itemID          int
	itemCategoryID  int
	itemName        string
	itemPrice       float64
	itemPurchaseDate string
	searchKeyword   string
)

// ==================== CATEGORY COMMANDS ====================

var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "Manajemen kategori barang",
}

var categoryListCmd = &cobra.Command{
	Use:   "list",
	Short: "Menampilkan semua kategori",
	Run: func(cmd *cobra.Command, args []string) {
		if handlerCategory == nil {
			fmt.Println("Error: Handler category belum diinisialisasi")
			return
		}
		handlerCategory.GetCategory()
	},
}

var categoryAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Menambahkan kategori baru",
	Run: func(cmd *cobra.Command, args []string) {
		if handlerCategory == nil {
			fmt.Println("Error: Handler category belum diinisialisasi")
			return
		}

		if categoryName == "" {
			fmt.Println("Error: Nama kategori tidak boleh kosong")
			return
		}

		handlerCategory.AddCategory(categoryName, categoryDescription)
	},
}

var categoryDetailCmd = &cobra.Command{
	Use:   "detail",
	Short: "Menampilkan detail kategori berdasarkan ID",
	Run: func(cmd *cobra.Command, args []string) {
		if handlerCategory == nil {
			fmt.Println("Error: Handler category belum diinisialisasi")
			return
		}

		if categoryID <= 0 {
			fmt.Println("Error: ID kategori harus lebih dari 0")
			return
		}

		handlerCategory.GetCategoryById(categoryID)
	},
}

var categoryUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Mengupdate kategori",
	Run: func(cmd *cobra.Command, args []string) {
		if handlerCategory == nil {
			fmt.Println("Error: Handler category belum diinisialisasi")
			return
		}

		if categoryID <= 0 {
			fmt.Println("Error: ID kategori harus lebih dari 0")
			return
		}

		if categoryName == "" {
			fmt.Println("Error: Nama kategori tidak boleh kosong")
			return
		}

		handlerCategory.UpdateCategory(categoryID, categoryName, categoryDescription)
	},
}

var categoryDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Menghapus kategori",
	Run: func(cmd *cobra.Command, args []string) {
		if handlerCategory == nil {
			fmt.Println("Error: Handler category belum diinisialisasi")
			return
		}

		if categoryID <= 0 {
			fmt.Println("Error: ID kategori harus lebih dari 0")
			return
		}

		handlerCategory.DeleteCategory(categoryID)
	},
}

// ==================== ITEM COMMANDS ====================

var itemCmd = &cobra.Command{
	Use:   "item",
	Short: "Manajemen barang inventaris",
}

var itemListCmd = &cobra.Command{
	Use:   "list",
	Short: "Menampilkan semua barang inventaris",
	Run: func(cmd *cobra.Command, args []string) {
		if handlerManagement == nil {
			fmt.Println("Error: Handler management belum diinisialisasi")
			return
		}
		handlerManagement.GetAllItems()
	},
}

var itemAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Menambahkan barang inventaris baru",
	Run: func(cmd *cobra.Command, args []string) {
		if handlerManagement == nil {
			fmt.Println("Error: Handler management belum diinisialisasi")
			return
		}

		if itemCategoryID <= 0 {
			fmt.Println("Error: ID kategori harus lebih dari 0")
			return
		}

		if itemName == "" {
			fmt.Println("Error: Nama barang tidak boleh kosong")
			return
		}

		if itemPrice <= 0 {
			fmt.Println("Error: Harga barang harus lebih dari 0")
			return
		}

		if itemPurchaseDate == "" {
			fmt.Println("Error: Tanggal pembelian tidak boleh kosong (format: YYYY-MM-DD)")
			return
		}

		handlerManagement.AddItem(itemCategoryID, itemName, itemPrice, itemPurchaseDate)
	},
}

var itemDetailCmd = &cobra.Command{
	Use:   "detail",
	Short: "Menampilkan detail barang berdasarkan ID",
	Run: func(cmd *cobra.Command, args []string) {
		if handlerManagement == nil {
			fmt.Println("Error: Handler management belum diinisialisasi")
			return
		}

		if itemID <= 0 {
			fmt.Println("Error: ID barang harus lebih dari 0")
			return
		}

		handlerManagement.GetItemById(itemID)
	},
}

var itemUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Mengupdate barang inventaris",
	Run: func(cmd *cobra.Command, args []string) {
		if handlerManagement == nil {
			fmt.Println("Error: Handler management belum diinisialisasi")
			return
		}

		if itemID <= 0 {
			fmt.Println("Error: ID barang harus lebih dari 0")
			return
		}

		if itemCategoryID <= 0 {
			fmt.Println("Error: ID kategori harus lebih dari 0")
			return
		}

		if itemName == "" {
			fmt.Println("Error: Nama barang tidak boleh kosong")
			return
		}

		if itemPrice <= 0 {
			fmt.Println("Error: Harga barang harus lebih dari 0")
			return
		}

		if itemPurchaseDate == "" {
			fmt.Println("Error: Tanggal pembelian tidak boleh kosong (format: YYYY-MM-DD)")
			return
		}

		handlerManagement.UpdateItem(itemID, itemCategoryID, itemName, itemPrice, itemPurchaseDate)
	},
}

var itemDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Menghapus barang inventaris",
	Run: func(cmd *cobra.Command, args []string) {
		if handlerManagement == nil {
			fmt.Println("Error: Handler management belum diinisialisasi")
			return
		}

		if itemID <= 0 {
			fmt.Println("Error: ID barang harus lebih dari 0")
			return
		}

		handlerManagement.DeleteItem(itemID)
	},
}

var itemSearchCmd = &cobra.Command{
	Use:   "search",
	Short: "Mencari barang berdasarkan nama",
	Run: func(cmd *cobra.Command, args []string) {
		if handlerManagement == nil {
			fmt.Println("Error: Handler management belum diinisialisasi")
			return
		}

		if searchKeyword == "" {
			fmt.Println("Error: Keyword pencarian tidak boleh kosong")
			return
		}

		handlerManagement.SearchItems(searchKeyword)
	},
}

// ==================== OLD ITEMS COMMAND ====================

var oldItemsCmd = &cobra.Command{
	Use:   "old-items",
	Short: "Menampilkan barang yang perlu diganti (> 100 hari)",
	Run: func(cmd *cobra.Command, args []string) {
		if handlerOld == nil {
			fmt.Println("Error: Handler old items belum diinisialisasi")
			return
		}
		handlerOld.GetOldItems()
	},
}

// ==================== REPORT COMMAND ====================

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Menampilkan laporan investasi dan depresiasi",
	Run: func(cmd *cobra.Command, args []string) {
		if handlerReport == nil {
			fmt.Println("Error: Handler report belum diinisialisasi")
			return
		}
		handlerReport.GetInvestmentReport()
	},
}

// ==================== INIT FUNCTIONS ====================

func InitHandlers(hCategory *handler.HandlerCategory, hManagement *handler.HandlerManagement, hOld *handler.HandlerOld, hReport *handler.HandlerReport) {
	handlerCategory = hCategory
	handlerManagement = hManagement
	handlerOld = hOld
	handlerReport = hReport
}

func init() {
	// ===== Category Flags =====
	// Tambah
	categoryAddCmd.Flags().StringVarP(&categoryName, "name", "n", "", "Nama kategori")
	categoryAddCmd.Flags().StringVarP(&categoryDescription, "description", "d", "", "Deskripsi kategori")
	categoryAddCmd.MarkFlagRequired("name")

	// Detail
	categoryDetailCmd.Flags().IntVarP(&categoryID, "id", "i", 0, "ID kategori")
	categoryDetailCmd.MarkFlagRequired("id")

	// Perbarui
	categoryUpdateCmd.Flags().IntVarP(&categoryID, "id", "i", 0, "ID kategori")
	categoryUpdateCmd.Flags().StringVarP(&categoryName, "name", "n", "", "Nama kategori baru")
	categoryUpdateCmd.Flags().StringVarP(&categoryDescription, "description", "d", "", "Deskripsi kategori baru")
	categoryUpdateCmd.MarkFlagRequired("id")
	categoryUpdateCmd.MarkFlagRequired("name")

	// Hapus
	categoryDeleteCmd.Flags().IntVarP(&categoryID, "id", "i", 0, "ID kategori")
	categoryDeleteCmd.MarkFlagRequired("id")

	// Tambahkan subcommand ke category
	categoryCmd.AddCommand(categoryListCmd)
	categoryCmd.AddCommand(categoryAddCmd)
	categoryCmd.AddCommand(categoryDetailCmd)
	categoryCmd.AddCommand(categoryUpdateCmd)
	categoryCmd.AddCommand(categoryDeleteCmd)

	// ===== Item Flags =====
	// Tambah
	itemAddCmd.Flags().IntVarP(&itemCategoryID, "category", "c", 0, "ID kategori barang")
	itemAddCmd.Flags().StringVarP(&itemName, "name", "n", "", "Nama barang")
	itemAddCmd.Flags().Float64VarP(&itemPrice, "price", "p", 0, "Harga barang")
	itemAddCmd.Flags().StringVarP(&itemPurchaseDate, "date", "d", "", "Tanggal pembelian (YYYY-MM-DD)")
	itemAddCmd.MarkFlagRequired("category")
	itemAddCmd.MarkFlagRequired("name")
	itemAddCmd.MarkFlagRequired("price")
	itemAddCmd.MarkFlagRequired("date")

	// Detail
	itemDetailCmd.Flags().IntVarP(&itemID, "id", "i", 0, "ID barang")
	itemDetailCmd.MarkFlagRequired("id")

	// Perbarui
	itemUpdateCmd.Flags().IntVarP(&itemID, "id", "i", 0, "ID barang")
	itemUpdateCmd.Flags().IntVarP(&itemCategoryID, "category", "c", 0, "ID kategori barang")
	itemUpdateCmd.Flags().StringVarP(&itemName, "name", "n", "", "Nama barang")
	itemUpdateCmd.Flags().Float64VarP(&itemPrice, "price", "p", 0, "Harga barang")
	itemUpdateCmd.Flags().StringVarP(&itemPurchaseDate, "date", "d", "", "Tanggal pembelian (YYYY-MM-DD)")
	itemUpdateCmd.MarkFlagRequired("id")
	itemUpdateCmd.MarkFlagRequired("category")
	itemUpdateCmd.MarkFlagRequired("name")
	itemUpdateCmd.MarkFlagRequired("price")
	itemUpdateCmd.MarkFlagRequired("date")

	// Hapus
	itemDeleteCmd.Flags().IntVarP(&itemID, "id", "i", 0, "ID barang")
	itemDeleteCmd.MarkFlagRequired("id")

	// Cari
	itemSearchCmd.Flags().StringVarP(&searchKeyword, "keyword", "k", "", "Keyword pencarian")
	itemSearchCmd.MarkFlagRequired("keyword")

	// Tambahkan subcommand ke item
	itemCmd.AddCommand(itemListCmd)
	itemCmd.AddCommand(itemAddCmd)
	itemCmd.AddCommand(itemDetailCmd)
	itemCmd.AddCommand(itemUpdateCmd)
	itemCmd.AddCommand(itemDeleteCmd)
	itemCmd.AddCommand(itemSearchCmd)

	// ===== Add all commands to root =====
	rootCmd.AddCommand(categoryCmd)
	rootCmd.AddCommand(itemCmd)
	rootCmd.AddCommand(oldItemsCmd)
	rootCmd.AddCommand(reportCmd)
}