package service

import (
	"sapgo/entity"
	"sapgo/repository"
)

type CompanyService interface {
	GetCompany() []entity.Company
}
type companyService struct {
	companyRepo repository.CompanyRepository
}

func NewCompanyService(company repository.CompanyRepository) CompanyService {
	return &companyService{
		companyRepo: company,
	}
}

func (c *companyService) GetCompany() []entity.Company {
	return c.companyRepo.GetCompany()
}
