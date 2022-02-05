package service

import (
	"sapgo/entity"
	"sapgo/repository"
)

type CategoryService interface {
	GetCategory() []entity.Category
}
type categoryService struct {
	serviceRepository repository.CategoryRepository
}

func NewCategoryService(service repository.CategoryRepository) CategoryService {
	return &categoryService{
		serviceRepository: service,
	}
}

func (s *categoryService) GetCategory() []entity.Category {
	return s.serviceRepository.GetCategory()
}
