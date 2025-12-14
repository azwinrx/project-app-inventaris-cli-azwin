package service

import (
	"project-app-inventaris-cli-azwin/model"
	"project-app-inventaris-cli-azwin/repository"
)

type ServiceCategoryInterface interface {
	GetCategory() ([]model.Category, error)
	AddCategory(name, description string) error
	GetCategoryById(id int) (model.Category, error)
	UpdateCategory(id int, name, description string) error
	DeleteCategory(id int) error
}

type ServiceCategory struct {
	RepoCategory repository.RepositoryCategoryInterface
}

func NewServiceCategory(repoCategory repository.RepositoryCategoryInterface) ServiceCategory {
	return ServiceCategory{
		RepoCategory: repoCategory,
	}
}

func (serviceCategory *ServiceCategory) GetCategory() ([]model.Category, error) {
	return serviceCategory.RepoCategory.GetCategory()
}

func (serviceCategory *ServiceCategory) AddCategory(name, description string) error {
	return serviceCategory.RepoCategory.AddCategory(name, description)
}

func (serviceCategory *ServiceCategory) GetCategoryById(id int) (model.Category, error) {
	return serviceCategory.RepoCategory.GetCategoryById(id)
}

func (serviceCategory *ServiceCategory) UpdateCategory(id int, name, description string) error {
	return serviceCategory.RepoCategory.UpdateCategory(id, name, description)
}

func (serviceCategory *ServiceCategory) DeleteCategory(id int) error {
	return serviceCategory.RepoCategory.DeleteCategory(id)
}
