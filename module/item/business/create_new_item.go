package business

import (
	"context"
	"errors"
	model "go-test/module/item/model"
)

type CreateItemStorage interface {
	CreateItem(ctx context.Context, data *model.Item) error
}

type createNewBiz struct {
	store CreateItemStorage
}

func NewCreateNewBiz(store CreateItemStorage) *createNewBiz {
	return &createNewBiz{store: store}
}

func (biz *createNewBiz) CreateNewItem(ctx context.Context, data *model.Item) error {
	if data.Name == "" {
		return errors.New("Name is required")
	}
	data.Status = "active"
	if err := biz.store.CreateItem(ctx, data); err != nil {
		return err
	}
	return nil
}
