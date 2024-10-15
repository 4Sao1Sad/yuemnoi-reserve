package model

import (
	"gorm.io/gorm"
)

type RequestStatus string

const (
	Reject  RequestStatus = "Reject"
	Pending RequestStatus = "Pending"
	Accept  RequestStatus = "Accept"
)

type BorrowingRequest struct {
	gorm.Model
	LendingUserID   string        `gorm:"type:varchar(255);not null"`
	BorrowingUserID string        `gorm:"type:varchar(255);not null"`
	LendingPostID   string        `gorm:"type:varchar(255);not null"`
	BorrowingPostID string        `gorm:"type:varchar(255);not null"`
	Status          RequestStatus `gorm:"type:varchar(255);not null"`
	ActiveStatus    bool          `gorm:"not null;default:false"`
}
