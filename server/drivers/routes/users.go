package routes

import (
	"item_management/server/drivers/database"
	users_controllers "item_management/server/interfaces/controllers/users"

	"github.com/labstack/echo/v4"
)

func InitUsers(e *echo.Echo) {
	guestController := users_controllers.InitUserController(database.InitSQLHandler())
	e.GET("/guests/:id", func(c echo.Context) error { return guestController.Show(c) })
	e.GET("/guests", func(c echo.Context) error { return guestController.Index(c) })
}
