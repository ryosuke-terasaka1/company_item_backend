package routes

import (
	"item_management/server/drivers/database"

	"github.com/labstack/echo/v4"
)

func InitCompanies(e *echo.Echo) {
	companyController := controllers.InitCompanyController(database.InitSQLHandler)
	e.GET("/companies/:id", func(c echo.Context) error { return companyController.Show(c) })
	e.GET("/companies", func(c echo.Context) error { return companyController.Index(c) })
}
