package routes

import (
	"item_management/server/infrastructure/database"

	"github.com/labstack/echo/v4"
)

func InitUsers(e *echo.Echo) {
	guestController := controllers.InitGuestController(database.InitSQLHandler)
	e.GET("/guests/:id", func(c echo.Context) error { return guestController.Show(c) })
	e.GET("/guests", func(c echo.Context) error { return guestController.Index(c) })
}
