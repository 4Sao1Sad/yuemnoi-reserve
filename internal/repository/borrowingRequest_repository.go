package repository

import (
	"github.com/KKhimmoon/yuemnoi-reserve/internal/model"
	"gorm.io/gorm"
)

type BorrowingRequestRepositoryImpl struct {
	db *gorm.DB
}

func NewBorrowingRequestRepository(db *gorm.DB) *BorrowingRequestRepositoryImpl {
	return &BorrowingRequestRepositoryImpl{
		db: db,
	}
}

type BorrowingRequestRepository interface {
	GetMyBorrowingRequests(userId uint) ([]model.BorrowingRequest, error)
	GetMyLendingPosts(userId uint) ([]model.BorrowingRequest, error)
	AcceptBorrowingRequest(request model.BorrowingRequest) (model.BorrowingRequest, error)
	GetBorrowingRequestById(requestId uint) (model.BorrowingRequest, error)
	RejectBorrowingRequest(request model.BorrowingRequest) (model.BorrowingRequest, error)
	CreateBorrowingRequest(request model.BorrowingRequest) (model.BorrowingRequest, error)
	ReturnItemBorrowingRequest(request model.BorrowingRequest) (model.BorrowingRequest, error)
	GetMyActiveBorrowingRequests(userId uint) ([]model.BorrowingRequest, error)
	GetMyHistorryBorrowingRequests(userId uint) ([]model.BorrowingRequest, error)
}

func (r BorrowingRequestRepositoryImpl) GetMyBorrowingRequests(userId uint) ([]model.BorrowingRequest, error) {
	var requests []model.BorrowingRequest
	if err := r.db.Where("borrowing_user_id = ? AND status = ?", userId, model.Pending).Find(&requests).Error; err != nil {
		return nil, err
	}
	return requests, nil
}

func (r BorrowingRequestRepositoryImpl) GetMyLendingPosts(userId uint) ([]model.BorrowingRequest, error) {
	var requests []model.BorrowingRequest
	if err := r.db.Where("lending_user_id = ? AND status = ?", userId, model.Pending).Find(&requests).Error; err != nil {
		return nil, err
	}
	return requests, nil
}

func (r BorrowingRequestRepositoryImpl) AcceptBorrowingRequest(request model.BorrowingRequest) (model.BorrowingRequest, error) {
	if err := r.db.Model(&request).Update("status", model.Accepted).Error; err != nil {
		return model.BorrowingRequest{}, err
	}
	return request, nil
}

func (r BorrowingRequestRepositoryImpl) GetBorrowingRequestById(requestId uint) (model.BorrowingRequest, error) {
	var request model.BorrowingRequest
	if err := r.db.Where("id = ?", requestId).First(&request).Error; err != nil {
		return model.BorrowingRequest{}, err
	}

	return request, nil
}

func (r BorrowingRequestRepositoryImpl) RejectBorrowingRequest(request model.BorrowingRequest) (model.BorrowingRequest, error) {
	if err := r.db.Model(&request).Updates(map[string]interface{}{
		"status":        model.Rejected,
		"active_status": false,
	}).Error; err != nil {
		return model.BorrowingRequest{}, err
	}
	return request, nil
}

func (r BorrowingRequestRepositoryImpl) CreateBorrowingRequest(request model.BorrowingRequest) (model.BorrowingRequest, error) {
	if err := r.db.Create(&request).Error; err != nil {
		return model.BorrowingRequest{}, err
	}
	return request, nil
}

func (r BorrowingRequestRepositoryImpl) ReturnItemBorrowingRequest(request model.BorrowingRequest) (model.BorrowingRequest, error) {
	if err := r.db.Model(&request).Update("active_status", false).Error; err != nil {
		return model.BorrowingRequest{}, err
	}
	return request, nil
}

func (r BorrowingRequestRepositoryImpl) GetMyActiveBorrowingRequests(userId uint) ([]model.BorrowingRequest, error) {
	var requests []model.BorrowingRequest
	if err := r.db.Where("(lending_user_id = ? OR borrowing_user_id = ?) AND status = ? AND active_status = ?",
		userId, userId, model.Accepted, true).Find(&requests).Order("updated_at DESC").Error; err != nil {
		return nil, err
	}
	return requests, nil
}

func (r BorrowingRequestRepositoryImpl) GetMyHistorryBorrowingRequests(userId uint) ([]model.BorrowingRequest, error) {
	var requests []model.BorrowingRequest
	if err := r.db.Where("(lending_user_id = ? OR borrowing_user_id = ?) AND active_status = ?",
		userId, userId, false).Find(&requests).Order("updated_at DESC").Error; err != nil {
		return nil, err
	}
	return requests, nil
}
