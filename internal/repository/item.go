package repository

import (
	"github.com/KKhimmoon/yuemnoi-reserve/internal/model"
	"gorm.io/gorm"
)

type ItemRepositoryImpl struct {
	db *gorm.DB
}

type ItemRepository interface {
	CreateItem(item model.Item) (error)
}

func NewItemRepository(db *gorm.DB)  *ItemRepositoryImpl {
	return &ItemRepositoryImpl{db: db}
}

func (r ItemRepositoryImpl) CreateItem(item model.Item) (error) {
	err := r.db.Create(&item).Error
	if err != nil {
		return err
	}

	return nil
}