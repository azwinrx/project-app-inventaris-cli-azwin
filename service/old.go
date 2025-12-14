package service

import (
	"project-app-inventaris-cli-azwin/model"
	"project-app-inventaris-cli-azwin/repository"
)

type ServiceOldInterface interface {
	GetOldItems() ([]model.Management, error)
}

type ServiceOld struct {
	RepoOld repository.RepositoryOldInterface
}

func NewServiceOld(repoOld repository.RepositoryOldInterface) ServiceOld {
	return ServiceOld{
		RepoOld: repoOld,
	}
}

func (serviceOld *ServiceOld) GetOldItems() ([]model.Management, error) {
	return serviceOld.RepoOld.GetOldItems()
}

