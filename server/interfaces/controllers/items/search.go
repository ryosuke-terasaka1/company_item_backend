package items

import (
	"item_management/server/interfaces/gateways"
	items_gateway "item_management/server/interfaces/gateways/items"
	items_usecase "item_management/server/usecases/items"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ItemController struct {
	Interactor items_usecase.ItemInteractor
}

func InitItemController(sqlHandler gateways.SQLHandler) *ItemController {
	return &ItemController{
		Interactor: items_usecase.ItemInteractor{
			ItemPort: &items_gateway.ItemRepository{
				SQLHandler: sqlHandler,
			},
		},
	}
}

func (controller *ItemController) Show(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	item, err := controller.Interactor.ItemByID(id)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, item)
	return
}

func (controller *ItemController) Index(c echo.Context) (err error) {
	items, err := controller.Interactor.AllItems()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, items)
	return
}
