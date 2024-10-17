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

type LendingRequest struct {
	gorm.Model
	LendingUserID   uint64        `gorm:"not null"`
	BorrowingUserID uint64        `gorm:"not null"`
	PostID          uint64        `gorm:"not null"`
	Status          RequestStatus `gorm:"type:varchar(255);not null"`
	ActiveStatus    bool          `gorm:"not null;default:false"`
}
