package companies

import (
	"item_management/server/entities"
)

type CompanyInteractor struct {
	CompanyPort CompanyPort
}

func (interactor *CompanyInteractor) FindCompanyByID(id int) (Company entities.Company, err error) {
	Company, err = interactor.CompanyPort.FindCompanyByID(id)
	return
}

func (interactor *CompanyInteractor) FindAllCompanies() (Companies entities.Companies, err error) {
	Companies, err = interactor.CompanyPort.FindAllCompanies()
	return
}
