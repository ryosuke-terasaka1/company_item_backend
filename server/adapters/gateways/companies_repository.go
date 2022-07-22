package gateways

import "item_management/server/entities"

type CompanyRepository struct {
	SQLHandler
}

func (repo *CompanyRepository) FindByID(id int) (company entities.Company, err error) {
	if err = repo.Find(&company, id).Error; err != nil {
		return
	}
	return
}

func (repo *CompanyRepository) FindAll() (companies entities.Company, err error) {
	if err = repo.Find(&companies).Error; err != nil {
		return
	}
	return
}
