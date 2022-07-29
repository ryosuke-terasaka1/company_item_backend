package routes

import (
	"item_management/server/drivers/database"
	items_controllers "item_management/server/interfaces/controllers/items"

	"github.com/labstack/echo/v4"
)

func InitItems(e *echo.Echo) {
	guestController := items_controllers.InitItemController(database.InitSQLHandler())
	e.GET("/guests/:id", func(c echo.Context) error { return guestController.Show(c) })
	e.GET("/guests", func(c echo.Context) error { return guestController.Index(c) })
}
