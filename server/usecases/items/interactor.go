package items

import (
	"context"
	"item_management/server/entities"
)

type ItemInteractor struct {
	OutputPort ItemOutputPort
	Repository ItemRepository
}

func NewItemInputPort(outputPort ItemOutputPort, repository ItemRepository) ItemInputPort {
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
