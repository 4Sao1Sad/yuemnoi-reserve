package model

import (
	"gorm.io/gorm"
)

type RequestStatus string

const (
	Rejected RequestStatus = "Rejected"
	Pending  RequestStatus = "Pending"
	Accepted RequestStatus = "Accepted"
)

type BorrowingRequest struct {
	gorm.Model
	LendingUserID   uint          `gorm:"not null"`
	BorrowingUserID uint          `gorm:"not null"`
	LendingPostID   uint          `gorm:"not null"`
	BorrowingPostID uint          `gorm:"not null"`
	Status          RequestStatus `gorm:"type:varchar(255);not null"`
	ActiveStatus    bool          `gorm:"not null;default:false"`
}
