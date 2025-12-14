package service

import (
	"project-app-inventaris-cli-azwin/model"
	"project-app-inventaris-cli-azwin/repository"
)

type ServiceManagementInterface interface {
	GetAllItems() ([]model.Management, error)
	AddItem(categoryId int, name string, price float64, purchaseDate string) error
	GetItemById(id int) (model.Management, error)
	UpdateItem(id int, categoryId int, name string, price float64, purchaseDate string) error
	DeleteItem(id int) error
	SearchItemsByName(keyword string) ([]model.Management, error)
}

type ServiceManagement struct {
	RepoManagement repository.RepositoryManagementInterface
}

func NewServiceManagement(repoManagement repository.RepositoryManagementInterface) ServiceManagement {
	return ServiceManagement{
		RepoManagement: repoManagement,
	}
}

func (serviceManagement *ServiceManagement) GetAllItems() ([]model.Management, error) {
	return serviceManagement.RepoManagement.GetAllItems()
}

func (serviceManagement *ServiceManagement) AddItem(categoryId int, name string, price float64, purchaseDate string) error {
	return serviceManagement.RepoManagement.AddItem(categoryId, name, price, purchaseDate)
}

func (serviceManagement *ServiceManagement) GetItemById(id int) (model.Management, error) {
	return serviceManagement.RepoManagement.GetItemById(id)
}

func (serviceManagement *ServiceManagement) UpdateItem(id int, categoryId int, name string, price float64, purchaseDate string) error {
	return serviceManagement.RepoManagement.UpdateItem(id, categoryId, name, price, purchaseDate)
}

func (serviceManagement *ServiceManagement) DeleteItem(id int) error {
	return serviceManagement.RepoManagement.DeleteItem(id)
}
func (serviceManagement *ServiceManagement) SearchItemsByName(keyword string) ([]model.Management, error) {
	return serviceManagement.RepoManagement.SearchItemsByName(keyword)
}