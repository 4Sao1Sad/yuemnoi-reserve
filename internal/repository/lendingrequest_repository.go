package repository

import (
	"github.com/KKhimmoon/yuemnoi-reserve/internal/model"
	"gorm.io/gorm"
)

type LendingRequestRepositoryImpl struct {
	db *gorm.DB
}

type LendingRequestRepository interface {
	CreateLendingRequest(request model.Request) (*model.Request, error)
	GetLendingRequestById(id uint64) (*model.Request, error)
	RejectLendingRequest(id uint64) (*model.Request, error)
	AcceptLendingRequest(id uint64) (*model.Request, error)
}

func NewLendingRequestRepository(db *gorm.DB) LendingRequestRepository {
	return &LendingRequestRepositoryImpl{db}
}

func (r LendingRequestRepositoryImpl) CreateLendingRequest(request model.Request) (*model.Request, error) {
	err := r.db.Create(&request).Error
	if err != nil {
		return nil, err
	}
	return &request, nil
}

func (r LendingRequestRepositoryImpl) GetLendingRequestById(id uint64) (*model.Request, error) {
	var request model.Request
	err := r.db.First(&request, id).Error
	if err != nil {
		return nil, err
	}
	return &request, nil
}

func (r LendingRequestRepositoryImpl) RejectLendingRequest(id uint64) (*model.Request, error) {
	var request model.Request
	err := r.db.First(&request, id).Error
	if err != nil {
		return nil, err
	}
	request.Status = model.Reject
	err = r.db.Save(&request).Error
	if err != nil {
		return nil, err
	}
	return &request, nil
}

func (r LendingRequestRepositoryImpl) AcceptLendingRequest(id uint64) (*model.Request, error) {
	var request model.Request
	err := r.db.First(&request, id).Error
	if err != nil {
		return nil, err
	}
	request.Status = model.Accept
	err = r.db.Save(&request).Error
	if err != nil {
		return nil, err
	}
	return &request, nil
}
