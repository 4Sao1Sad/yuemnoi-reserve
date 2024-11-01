package repository

import (
	"github.com/KKhimmoon/yuemnoi-reserve/internal/model"
	"gorm.io/gorm"
)

type LendingRequestRepositoryImpl struct {
	db *gorm.DB
}

func NewLendingRequestRepository(db *gorm.DB) *LendingRequestRepositoryImpl {
	return &LendingRequestRepositoryImpl{
		db: db,
	}
}

type LendingRequestRepository interface {
	GetMyLendingRequests(userId uint) ([]model.LendingRequest, error)
	GetMyBorrowingPosts(userId uint) ([]model.LendingRequest, error)
	AcceptLendingRequest(request model.LendingRequest) (model.LendingRequest, error)
	GetLendingRequestById(requestId uint) (model.LendingRequest, error)
	RejectLendingRequest(request model.LendingRequest) (model.LendingRequest, error)
	CreateLendingRequest(request model.LendingRequest) (model.LendingRequest, error)
	ReturnItemLendingRequest(request model.LendingRequest) (model.LendingRequest, error)
	GetMyActiveLendingRequests(userId uint) ([]model.LendingRequest, error)
	GetMyHistorryLendingRequests(userId uint) ([]model.LendingRequest, error)
}

func (r LendingRequestRepositoryImpl) GetMyLendingRequests(userId uint) ([]model.LendingRequest, error) {
	var requests []model.LendingRequest
	if err := r.db.Where("lending_user_id = ?", userId).Find(&requests).Error; err != nil {
		return nil, err
	}
	return requests, nil
}

func (r LendingRequestRepositoryImpl) GetMyBorrowingPosts(userId uint) ([]model.LendingRequest, error) {
	var requests []model.LendingRequest
	if err := r.db.Where("borrowing_user_id = ?", userId).Find(&requests).Error; err != nil {
		return nil, err
	}
	return requests, nil
}

func (r LendingRequestRepositoryImpl) AcceptLendingRequest(request model.LendingRequest) (model.LendingRequest, error) {
	if err := r.db.Model(&request).Update("status", model.Accepted).Error; err != nil {
		return model.LendingRequest{}, err
	}
	return request, nil
}

func (r LendingRequestRepositoryImpl) GetLendingRequestById(requestId uint) (model.LendingRequest, error) {
	var request model.LendingRequest
	if err := r.db.Where("id = ?", requestId).First(&request).Error; err != nil {
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

func (r LendingRequestRepositoryImpl) CreateLendingRequest(request model.LendingRequest) (model.LendingRequest, error) {
	if err := r.db.Create(&request).Error; err != nil {
		return model.LendingRequest{}, err
	}
	return request, nil
}

func (r LendingRequestRepositoryImpl) ReturnItemLendingRequest(request model.LendingRequest) (model.LendingRequest, error) {
	if err := r.db.Model(&request).Update("active_status", false).Error; err != nil {
		return model.LendingRequest{}, err
	}
	return request, nil
}

func (r LendingRequestRepositoryImpl) GetMyActiveLendingRequests(userId uint) ([]model.LendingRequest, error) {
	var requests []model.LendingRequest
	if err := r.db.Where("(lending_user_id = ? OR borrowing_user_id = ?) AND status = ? AND active_status = ?",
		userId, userId, model.Accepted, true).Find(&requests).Order("updated_at DESC").Error; err != nil {
		return nil, err
	}
	return requests, nil
}

func (r LendingRequestRepositoryImpl) GetMyHistorryLendingRequests(userId uint) ([]model.LendingRequest, error) {
	var requests []model.LendingRequest
	if err := r.db.Where("(lending_user_id = ? OR borrowing_user_id = ?) AND active_status = ?",
		userId, userId, false).Find(&requests).Order("updated_at DESC").Error; err != nil {
		return nil, err
	}
	return requests, nil
}
