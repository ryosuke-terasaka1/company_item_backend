package items

import (
	"item_management/server/entities"
)

type ItemInteractor struct {
	ItemPort ItemPort
}

func (interactor *ItemInteractor) ItemByID(id int) (item entities.Item, err error) {
	item, err = interactor.ItemPort.FindItemByID(id)
	return
}

func (interactor *ItemInteractor) AllItems() (items entities.Items, err error) {
	items, err = interactor.ItemPort.FindAllItems()
	return
}
