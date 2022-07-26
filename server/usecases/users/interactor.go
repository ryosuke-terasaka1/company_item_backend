package users

import (
	"item_management/server/entities"
)

type UserInteractor struct {
	UserPort UserPort
}

func (interactor *UserInteractor) FindUserByID(id int) (user entities.User, err error) {
	user, err = interactor.UserPort.FindUserByID(id)
	return
}

func (interactor *UserInteractor) FindAllUsers() (users entities.Users, err error) {
	users, err = interactor.UserPort.FindAllUsers()
	return
}
