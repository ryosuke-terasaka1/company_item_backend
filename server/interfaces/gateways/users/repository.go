package users

import (
	"item_management/server/entities"
	"item_management/server/interfaces/gateways"
)

type UserRepository struct {
	gateways.SQLHandler
}

func (repo *UserRepository) FindUserByID(id int) (user entities.User, err error) {
	if err = repo.Find(&user, id).Error; err != nil {
		return
	}
	return
}

func (repo *UserRepository) FindAllUsers() (users entities.Users, err error) {
	if err = repo.Find(&users).Error; err != nil {
		return
	}
	return
}
