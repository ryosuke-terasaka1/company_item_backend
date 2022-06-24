package ports

import (
	"context"
	"item_management/entities"
)

type ItemInputPort interface {
	AddItem(ctx context.Context, item *entities.Item) error
	GetItems(ctx context.Context) error
}

type ItemOutputPort interface {
	OutputItems([]*entities.Item) error
	OutputError(error) error
}

type ItemRepository interface {
	AddItem(ctx context.Context, item *entities.Item) ([]*entities.Item, error)
	GetItems(ctx context.Context) ([]*entities.Item, error)
}
