package companies

import (
	"item_management/server/interfaces/gateways"
	companies_gateway "item_management/server/interfaces/gateways/companies"
	companies_usecase "item_management/server/usecases/companies"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CompanyController struct {
	Interactor companies_usecase.CompanyInteractor
}

func InitCompanyController(sqlHandler gateways.SQLHandler) *CompanyController {
	return &CompanyController{
		Interactor: companies_usecase.CompanyInteractor{
			CompanyPort: &companies_gateway.CompanyRepository{
				SQLHandler: sqlHandler,
			},
		},
	}
}

func (controller *CompanyController) Show(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	company, err := controller.Interactor.CompanyByID(id)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, company)
	return
}

func (controller *CompanyController) Index(c echo.Context) (err error) {
	companies, err := controller.Interactor.AllCompanies()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, companies)
	return
}
