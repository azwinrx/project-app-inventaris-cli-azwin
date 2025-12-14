package handler

import (
	"fmt"
	"project-app-inventaris-cli-azwin/service"
)

type HandlerOld struct {
	ServiceOld service.ServiceOldInterface
}

func NewHandlerOld(serviceOld service.ServiceOldInterface) HandlerOld {
	return HandlerOld{
		ServiceOld: serviceOld,
	}
}

func (handlerOld *HandlerOld) GetOldItems() {
	items, err := handlerOld.ServiceOld.GetOldItems()
	if err != nil {
		fmt.Println("Failed to retrieve old items:", err)
		return
	}

	if len(items) == 0 {
		fmt.Println("\n=== NO ITEMS NEED REPLACEMENT ===")
		fmt.Println("All items are still in good condition (less than 100 days old).")
		return
	}

	fmt.Println("\n=== ITEMS THAT NEED REPLACEMENT (> 100 DAYS) ===")
	fmt.Printf("%-5s %-20s %-20s %-15s %-15s %-10s\n", "ID", "Category", "Name", "Price", "Purchase Date", "Usage Days")
	fmt.Println("----------------------------------------------------------------------------------------------------")
	for _, item := range items {
		fmt.Printf("%-5d %-20s %-20s Rp%-13.2f %-15s %-10d\n", 
			item.ID, 
			item.CategoryName, 
			item.Name, 
			item.Price, 
			item.PurchaseDate.Format("2006-01-02"), 
			item.UsageDays)
	}
	fmt.Printf("\nTotal items need replacement: %d\n", len(items))
}
