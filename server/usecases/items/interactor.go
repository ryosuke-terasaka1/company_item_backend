package items

import (
	"item_management/server/entities"
)

type ItemInteractor struct {
	ItemPort ItemPort
}

func (interactor *ItemInteractor) FindItemByID(id int) (item entities.Item, err error) {
	item, err = interactor.ItemPort.FindItemByID(id)
	return
}

func (interactor *ItemInteractor) FindAllItems() (items entities.Items, err error) {
	items, err = interactor.ItemPort.FindAllItems()
	return
}
