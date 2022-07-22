package gateways

import "item_management/server/entities"

type ItemRepository struct {
	SQLHandler
}

func (repo *ItemRepository) ItemFindByID(id int) (item entities.Item, err error) {
	if err = repo.Find(&item, id).Error; err != nil {
		return
	}
	return
}

func (repo *ItemRepository) FindAll() (items entities.Item, err error) {
	if err = repo.Find(&items).Error; err != nil {
		return
	}
	return
}
