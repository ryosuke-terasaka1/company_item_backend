package users

import (
	"item_management/server/interfaces/gateways"
	users_gateway "item_management/server/interfaces/gateways/users"
	users_usecase "item_management/server/usecases/users"
)

type UserController struct {
	Interactor users_usecase.UserInteractor
}

func InitUserController(sqlHandler gateways.SQLHandler) *UserController {
	return &UserController{
		Interactor: users_usecase.UserInteractor{
			UserPort: &users_gateway.UserRepository{
				SQLHandler: sqlHandler,
			},
		},
	}
}
