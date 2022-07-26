package items

import (
	"item_management/server/entities"
)

type ItemPort interface {
	FindItemByID(id int) (entities.Item, error)
	FindAllItems() (entities.Items, error)
}
