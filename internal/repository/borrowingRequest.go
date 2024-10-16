package repository

import (
	"github.com/KKhimmoon/yuemnoi-reserve/internal/model"
	"gorm.io/gorm"
)

type BorrowingRepositoryImpl struct {
	db *gorm.DB
}

type BorrowingRepository interface {
	CreateRequestFromBorrowingPost(request model.BorrowingRequest) (model.BorrowingRequest, error)
	GetRequestById(requestId uint) (model.BorrowingRequest, error)
	ConfirmBorrowingRequest(request model.BorrowingRequest) (model.BorrowingRequest, error)
	RejectBorrowingRequest(request model.BorrowingRequest) (model.BorrowingRequest, error)
	ReturnItemRequest(request model.BorrowingRequest) (model.BorrowingRequest, error)
}

func NewBorrowingRepository(db *gorm.DB) *BorrowingRepositoryImpl {
	return &BorrowingRepositoryImpl{db: db}
}

func (r BorrowingRepositoryImpl) CreateRequestFromBorrowingPost(request model.BorrowingRequest) (model.BorrowingRequest, error) {
	if err := r.db.Create(&request).Error; err != nil {
		return model.BorrowingRequest{}, err
	}
	return request, nil
}

func (r BorrowingRepositoryImpl) GetRequestById(requestId uint) (model.BorrowingRequest, error) {
	var request model.BorrowingRequest
	if err := r.db.Where("id = ?", requestId).First(&request).Error; err != nil {
		return model.BorrowingRequest{}, err
	}

	return request, nil
}

func (r BorrowingRepositoryImpl) ConfirmBorrowingRequest(request model.BorrowingRequest) (model.BorrowingRequest, error) {
	if err := r.db.Model(&request).Update("status", model.Accepted).Error; err != nil {
		return model.BorrowingRequest{}, err
	}
	return request, nil
}

func (r BorrowingRepositoryImpl) RejectBorrowingRequest(request model.BorrowingRequest) (model.BorrowingRequest, error) {
	if err := r.db.Model(&request).Updates(map[string]interface{}{
		"status":        model.Rejected,
		"active_status": false,
	}).Error; err != nil {
		return model.BorrowingRequest{}, err
	}
	return request, nil
}

func (r BorrowingRepositoryImpl) ReturnItemRequest(request model.BorrowingRequest) (model.BorrowingRequest, error) {
	if err := r.db.Model(&request).Update("active_status", false).Error; err != nil {
		return model.BorrowingRequest{}, err
	}
	return request, nil
}
