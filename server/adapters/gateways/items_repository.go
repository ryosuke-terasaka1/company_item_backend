package gateways

import "item_management/server/entities"

type GuestRepository struct {
	SQLHandler
}

func (repo *GuestRepository) FindByID(id int) (guest entities.Item, err error) {
	if err = repo.Find(&guest, id).Error; err != nil {
		return
	}
	return
}

func (repo *GuestRepository) FindAll() (guests entities.User, err error) {
	if err = repo.Find(&guests).Error; err != nil {
		return
	}
	return
}
