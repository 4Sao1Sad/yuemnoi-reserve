package repository

import (
	"github.com/KKhimmoon/yuemnoi-reserve/internal/model"
	"gorm.io/gorm"
)

type BorrowingRepositoryImpl struct {
	db *gorm.DB
}

type BorrowingRepository interface {
	CreateBorrowingRequest(request model.BorrowingRequest) (model.BorrowingRequest, error)
	GetBorrowingRequestById(requestId uint) (model.BorrowingRequest, error)
	AcceptBorrowingRequest(request model.BorrowingRequest) (model.BorrowingRequest, error)
	RejectBorrowingRequest(request model.BorrowingRequest) (model.BorrowingRequest, error)
	ReturnItemBorrowingRequest(request model.BorrowingRequest) (model.BorrowingRequest, error)
	GetBorrowingRequests() ([]model.BorrowingRequest, error)
}

func NewBorrowingRepository(db *gorm.DB) *BorrowingRepositoryImpl {
	return &BorrowingRepositoryImpl{db: db}
}

func (r BorrowingRepositoryImpl) CreateBorrowingRequest(request model.BorrowingRequest) (model.BorrowingRequest, error) {
	if err := r.db.Create(&request).Error; err != nil {
		return model.BorrowingRequest{}, err
	}
	return request, nil
}

func (r BorrowingRepositoryImpl) GetBorrowingRequestById(requestId uint) (model.BorrowingRequest, error) {
	var request model.BorrowingRequest
	if err := r.db.Where("id = ?", requestId).First(&request).Error; err != nil {
		return model.BorrowingRequest{}, err
	}

	return request, nil
}

func (r BorrowingRepositoryImpl) GetBorrowingRequests() ([]model.BorrowingRequest, error) {
	var requests []model.BorrowingRequest
	if err := r.db.Find(&requests).Error; err != nil {
		return nil, err
	}
	return requests, nil
}

func (r BorrowingRepositoryImpl) AcceptBorrowingRequest(request model.BorrowingRequest) (model.BorrowingRequest, error) {
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

func (r BorrowingRepositoryImpl) ReturnItemBorrowingRequest(request model.BorrowingRequest) (model.BorrowingRequest, error) {
	if err := r.db.Model(&request).Update("active_status", false).Error; err != nil {
		return model.BorrowingRequest{}, err
	}
	return request, nil
}
