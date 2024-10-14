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

type LendingRequest struct {
	gorm.Model
	LendingUserID   string        `gorm:"type:varchar(255);not null"`
	BorrowingUserID string        `gorm:"type:varchar(255);not null"`
	PostID          string        `gorm:"type:varchar(255);not null"`
	Status          RequestStatus `gorm:"type:varchar(255);not null"`
	ActiveStatus    bool          `gorm:"not null;default:false"`
}
