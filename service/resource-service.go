package service

import (
	"sapgo/entity"
	"sapgo/repository"
)

type ResourceService interface {
	GetResource() []entity.Resource
}
type resourceService struct {
	resRepo repository.ResourceRepository
}

func NewResourceService(company repository.ResourceRepository) ResourceService {
	return &resourceService{
		resRepo: company,
	}
}

func (c *resourceService) GetResource() []entity.Resource {
	return c.resRepo.GetResource()
}
