package users

import (
	"item_management/server/entities"
	"item_management/server/interfaces/gateways"
)

type UserRepository struct {
	gateways.SQLHandler
}

func (repo *UserRepository) UserFindByID(id int) (user entities.User, err error) {
	if err = repo.Find(&user, id).Error; err != nil {
		return
	}
	return
}

func (repo *UserRepository) UserFindAll() (users entities.User, err error) {
	if err = repo.Find(&users).Error; err != nil {
		return
	}
	return
}
