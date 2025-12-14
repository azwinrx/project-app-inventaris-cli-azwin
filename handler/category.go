package handler

import (
	"fmt"
	"os"
	"project-app-inventaris-cli-azwin/service"
	"text/tabwriter"
)

type HandlerCategory struct {
	ServiceCategory service.ServiceCategoryInterface
}

func NewHandlerCategory(serviceCategory service.ServiceCategoryInterface) HandlerCategory {
	return HandlerCategory{
		ServiceCategory: serviceCategory,
	}
}

func (handlerCategory *HandlerCategory) GetCategory() {
	// call service category
	categories, err := handlerCategory.ServiceCategory.GetCategory()
	if err != nil {
		fmt.Println("Failed to retrieve data:", err)
	}

	fmt.Println("\n=== CATEGORIES ===")
	
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, "ID\tName\tDescription")
	fmt.Fprintln(w, "--\t----\t-----------")
	
	for _, c := range categories {
		fmt.Fprintf(w, "%d\t%s\t%s\n", c.ID, c.Name, c.Description)
	}
	w.Flush()
}

func (handlerCategory *HandlerCategory) AddCategory(name, description string) {
	err := handlerCategory.ServiceCategory.AddCategory(name, description)
	if err != nil {
		fmt.Println("Failed to add category:", err)
		return
	}
	fmt.Println("Category added successfully!")
}

func (handlerCategory *HandlerCategory) GetCategoryById(id int) {
	category, err := handlerCategory.ServiceCategory.GetCategoryById(id)
	if err != nil {
		fmt.Println("Failed to retrieve category:", err)
		return
	}

	fmt.Println("\n=== CATEGORY DETAIL ===")
	fmt.Printf("ID          : %d\n", category.ID)
	fmt.Printf("Name        : %s\n", category.Name)
	fmt.Printf("Description : %s\n", category.Description)
}

func (handlerCategory *HandlerCategory) UpdateCategory(id int, name, description string) {
	err := handlerCategory.ServiceCategory.UpdateCategory(id, name, description)
	if err != nil {
		fmt.Println("Failed to update category:", err)
		return
	}
	fmt.Println("Category updated successfully!")
}

func (handlerCategory *HandlerCategory) DeleteCategory(id int) {
	err := handlerCategory.ServiceCategory.DeleteCategory(id)
	if err != nil {
		fmt.Println("Failed to delete category:", err)
		return
	}
	fmt.Println("Category deleted successfully!")
}
