package companies

import (
	"item_management/server/entities"
)

type CompanyPort interface {
	FindCompanyByID(id int) (entities.Company, error)
	FindAllCompanies() (entities.Companies, error)
}
