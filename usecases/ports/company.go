package ports

import (
	"context"
	"item_management/entities"
)

type CompanyInputPort interface {
	AddCompany(ctx context.Context, company *entities.Company) error
	GetCompanys(ctx context.Context) error
}

type CompanyOutputPort interface {
	OutputCompanys([]*entities.Company) error
	OutputError(error) error
}

type CompanyRepository interface {
	AddCompany(ctx context.Context, company *entities.Company) ([]*entities.Company, error)
	GetCompanys(ctx context.Context) ([]*entities.Company, error)
}
