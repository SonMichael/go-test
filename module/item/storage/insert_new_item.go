package storage

import (
	"context"
	"go-test/module/item/model"
)

func (s *mysqlStorage) CreateItem(ctx context.Context, data *model.Item) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}
	return nil
}