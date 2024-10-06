package model

import (
	"gorm.io/gorm"
)

type RequestType string
type RequestStatus string

const (
	Lending   RequestType = "Lending"
	Borrowing RequestType = "Borrowing"

	Reject  RequestStatus = "Reject"
	Pending RequestStatus = "Pending"
	Accept  RequestStatus = "Accept"
)

type Request struct {
	gorm.Model
	RequestUserID  string        `gorm:"type:varchar(255);not null"`
	ResponseUserID string        `gorm:"type:varchar(255);not null"`
	PostID         string        `gorm:"type:varchar(255);not null"`
	Type           RequestType   `gorm:"type:varchar(255);not null"`
	Status         RequestStatus `gorm:"type:varchar(255);not null"`
}
