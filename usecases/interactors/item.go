package interactors

import (
	"context"
	"item_management/entities"
	"item_management/usecases/ports"
)

type ItemInteractor struct {
	OutputPort ports.ItemOutputPort
	Repository ports.ItemRepository
}

func NewItemInputPort(outputPort ports.ItemOutputPort, repository ports.ItemRepository) ports.ItemInputPort {
	return &ItemInteractor{
		OutputPort: outputPort,
		Repository: repository,
	}
}

func (u *ItemInteractor) AddItem(ctx context.Context, item *entities.Item) error {
	items, err := u.Repository.AddItem(ctx, item)
	if err != nil {
		return u.OutputPort.OutputError(err)
	}

	return u.OutputPort.OutputItems(items)
}

func (u *ItemInteractor) GetItems(ctx context.Context) error {
	items, err := u.Repository.GetItems(ctx)
	if err != nil {
		return u.OutputPort.OutputError(err)
	}

	return u.OutputPort.OutputItems(items)
}
