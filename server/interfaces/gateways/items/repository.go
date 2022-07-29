package items

import (
	"item_management/server/entities"
	"item_management/server/interfaces/gateways"
)

type ItemRepository struct {
	gateways.SQLHandler
}

func (repo *ItemRepository) FindItemByID(id int) (item entities.Item, err error) {
	if err = repo.Find(&item, id).Error; err != nil {
		return
	}
	return
}

func (repo *ItemRepository) FindAllItems() (items entities.Items, err error) {
	if err = repo.Find(&items).Error; err != nil {
		return
	}
	return
}
