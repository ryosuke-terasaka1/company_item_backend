package items

import (
	"item_management/server/interfaces/gateways"
	items_gateway "item_management/server/interfaces/gateways/items"
	items_usecase "item_management/server/usecases/items"
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
