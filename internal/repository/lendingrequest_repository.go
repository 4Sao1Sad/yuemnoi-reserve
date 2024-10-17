package repository

import (
	"github.com/KKhimmoon/yuemnoi-reserve/internal/model"
	"gorm.io/gorm"
)

type LendingRequestRepositoryImpl struct {
	db *gorm.DB
}

type LendingRequestRepository interface {
	CreateLendingRequest(request model.LendingRequest) (model.LendingRequest, error)
	GetLendingRequestById(requestId uint64) (model.LendingRequest, error)
	RejectLendingRequest(request model.LendingRequest) (model.LendingRequest, error)
	AcceptLendingRequest(request model.LendingRequest) (model.LendingRequest, error)
	ReturnItemRequest(request model.LendingRequest) (model.LendingRequest, error)
}

func NewLendingRequestRepository(db *gorm.DB) *LendingRequestRepositoryImpl {
	return &LendingRequestRepositoryImpl{db: db}
}

func (r LendingRequestRepositoryImpl) CreateLendingRequest(request model.LendingRequest) (model.LendingRequest, error) {
	err := r.db.Create(&request).Error
	if err != nil {
		return model.LendingRequest{}, err
	}
	return request, nil
}

func (r LendingRequestRepositoryImpl) GetLendingRequestById(requestId uint64) (model.LendingRequest, error) {
	var request model.LendingRequest
	err := r.db.First(&request, requestId).Error
	if err != nil {
		return model.LendingRequest{}, err
	}
	return request, nil
}

func (r LendingRequestRepositoryImpl) RejectLendingRequest(request model.LendingRequest) (model.LendingRequest, error) {
	if err := r.db.Model(&request).Updates(map[string]interface{}{
		"status":        model.Rejected,
		"active_status": false,
	}).Error; err != nil {
		return model.LendingRequest{}, err
	}
	return request, nil
}

func (r LendingRequestRepositoryImpl) AcceptLendingRequest(request model.LendingRequest) (model.LendingRequest, error) {
	if err := r.db.Model(&request).Update("status", model.Accepted).Error; err != nil {
		return model.LendingRequest{}, err
	}
	return request, nil
}

func (r LendingRequestRepositoryImpl) ReturnItemRequest(request model.LendingRequest) (model.LendingRequest, error) {
	if err := r.db.Model(&request).Update("active_status", false).Error; err != nil {
		return model.LendingRequest{}, err
	}
	return request, nil
}
