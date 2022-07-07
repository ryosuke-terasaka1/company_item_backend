package routes

import "github.com/labstack/echo/v4"

func InitGuest(e *echo.Echo) {
	guestController := controller.InitGuestController(database.InitSQLHandler)
	e.GET("/guests/:id", func(c echo.Context) error { return guestController.Show(c) })
	e.GET("/guests", func(c echo.Context) error { return guestController.Index(c) })
}
