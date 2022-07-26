package users

import (
	"item_management/server/entities"
)

type UserPort interface {
	FindUserByID(id int) (entities.User, error)
	FindAllUsers() (entities.Users, error)
}
