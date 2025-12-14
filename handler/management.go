package handler

import (
	"fmt"
	"os"
	"project-app-inventaris-cli-azwin/service"
	"text/tabwriter"
)

type HandlerManagement struct {
	ServiceManagement service.ServiceManagementInterface
}

func NewHandlerManagement(serviceManagement service.ServiceManagementInterface) HandlerManagement {
	return HandlerManagement{
		ServiceManagement: serviceManagement,
	}
}

func (handlerManagement *HandlerManagement) GetAllItems() {
	items, err := handlerManagement.ServiceManagement.GetAllItems()
	if err != nil {
		fmt.Println("Failed to retrieve items:", err)
		return
	}

	fmt.Println("\n=== INVENTORY ITEMS ===")
	
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, "ID\tCategory\tName\tPrice\tPurchase Date\tUsage Days")
	fmt.Fprintln(w, "--\t--------\t----\t-----\t-------------\t----------")
	
	for _, item := range items {
		fmt.Fprintf(w, "%d\t%s\t%s\tRp %.2f\t%s\t%d days\n", 
			item.ID, 
			item.CategoryName, 
			item.Name, 
			item.Price, 
			item.PurchaseDate.Format("2006-01-02"), 
			item.UsageDays)
	}
	w.Flush()
}

func (handlerManagement *HandlerManagement) AddItem(categoryId int, name string, price float64, purchaseDate string) {
	err := handlerManagement.ServiceManagement.AddItem(categoryId, name, price, purchaseDate)
	if err != nil {
		fmt.Println("Failed to add item:", err)
		return
	}
	fmt.Println("Item added successfully!")
}

func (handlerManagement *HandlerManagement) GetItemById(id int) {
	item, err := handlerManagement.ServiceManagement.GetItemById(id)
	if err != nil {
		fmt.Println("Failed to retrieve item:", err)
		return
	}

	fmt.Println("\n=== ITEM DETAIL ===")
	fmt.Printf("ID            : %d\n", item.ID)
	fmt.Printf("Category      : %s (ID: %d)\n", item.CategoryName, item.CategoryId)
	fmt.Printf("Name          : %s\n", item.Name)
	fmt.Printf("Price         : Rp%.2f\n", item.Price)
	fmt.Printf("Purchase Date : %s\n", item.PurchaseDate.Format("2006-01-02"))
	fmt.Printf("Usage Days    : %d days\n", item.UsageDays)
}

func (handlerManagement *HandlerManagement) UpdateItem(id int, categoryId int, name string, price float64, purchaseDate string) {
	err := handlerManagement.ServiceManagement.UpdateItem(id, categoryId, name, price, purchaseDate)
	if err != nil {
		fmt.Println("Failed to update item:", err)
		return
	}
	fmt.Println("Item updated successfully!")
}

func (handlerManagement *HandlerManagement) DeleteItem(id int) {
	err := handlerManagement.ServiceManagement.DeleteItem(id)
	if err != nil {
		fmt.Println("Failed to delete item:", err)
		return
	}
	fmt.Println("Item deleted successfully!")
}
func (handlerManagement *HandlerManagement) SearchItems(keyword string) {
	items, err := handlerManagement.ServiceManagement.SearchItemsByName(keyword)
	if err != nil {
		fmt.Println("Failed to search items:", err)
		return
	}

	if len(items) == 0 {
		fmt.Printf("\nNo items found with keyword: '%s'\n", keyword)
		return
	}

	fmt.Printf("\n=== SEARCH RESULTS (keyword: '%s') ===\n", keyword)
	
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, "ID\tCategory\tName\tPrice\tPurchase Date\tUsage Days")
	fmt.Fprintln(w, "--\t--------\t----\t-----\t-------------\t----------")
	
	for _, item := range items {
		fmt.Fprintf(w, "%d\t%s\t%s\tRp %.2f\t%s\t%d days\n", 
			item.ID, 
			item.CategoryName, 
			item.Name, 
			item.Price, 
			item.PurchaseDate.Format("2006-01-02"), 
			item.UsageDays)
	}
	w.Flush()
	
	fmt.Printf("\nTotal items found: %d\n", len(items))
}