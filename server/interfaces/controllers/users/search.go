package users

import (
	"item_management/server/interfaces/gateways"
	users_gateway "item_management/server/interfaces/gateways/users"
	users_usecase "item_management/server/usecases/users"
	"strconv"

	"github.com/labstack/echo/v4"
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

func (controller *UserController) Show(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	company, err := controller.Interactor.UserByID(id)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, company)
	return
}

func (controller *UserController) Index(c echo.Context) (err error) {
	users, err := controller.Interactor.AllUsers()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, users)
	return
}
