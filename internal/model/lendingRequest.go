package model

import (
	"gorm.io/gorm"
)

type RequestStatus string

const (
	Rejected RequestStatus = "Reject"
	Pending  RequestStatus = "Pending"
	Accepted RequestStatus = "Accept"
)

type LendingRequest struct {
	gorm.Model
	LendingUserID   uint          `gorm:"not null"`
	BorrowingUserID uint          `gorm:"not null"`
	PostID          uint          `gorm:"not null"`
	Status          RequestStatus `gorm:"type:varchar(255);not null"`
	ActiveStatus    bool          `gorm:"not null;default:false"`
}
