package companies

import (
	"context"
	"item_management/server/entities"
)

type CompanyInteractor struct {
	OutputPort CompanyOutputPort
	Repository CompanyRepository
}

func NewCompanyInputPort(outputPort CompanyOutputPort, repository CompanyRepository) CompanyInputPort {
	return &CompanyInteractor{
		OutputPort: outputPort,
		Repository: repository,
	}
}

func (u *CompanyInteractor) AddCompany(ctx context.Context, company *entities.Company) error {
	companies, err := u.Repository.AddCompany(ctx, company)
	if err != nil {
		return u.OutputPort.OutputError(err)
	}

	return u.OutputPort.OutputCompanys(companies)
}

func (u *CompanyInteractor) GetCompanys(ctx context.Context) error {
	companies, err := u.Repository.GetCompanys(ctx)
	if err != nil {
		return u.OutputPort.OutputError(err)
	}

	return u.OutputPort.OutputCompanys(companies)
}
