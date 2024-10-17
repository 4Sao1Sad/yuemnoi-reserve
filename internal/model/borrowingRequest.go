package model

import (
	"gorm.io/gorm"
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
