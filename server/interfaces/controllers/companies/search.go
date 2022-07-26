package companies

import (
	"item_management/server/interfaces/gateways"
	companies_gateway "item_management/server/interfaces/gateways/companies"
	companies_usecase "item_management/server/usecases/companies"
)

type CompanyController struct {
	Interactor companies_usecase.CompanyInteractor
}

func InitCompanyController(sqlHandler gateways.SQLHandler) *CompanyController {
	return &CompanyController{
		Interactor: companies_usecase.CompanyInteractor{
			CompanyPort: &companies_gateway.CompanyRepository{
				SQLHandler: sqlHandler,
			},
		},
	}
}
