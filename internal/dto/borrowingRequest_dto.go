package dto

import "github.com/KKhimmoon/yuemnoi-reserve/internal/model"

type GetMyBorrowingRequestsResponse struct {
	ID              uint                `json:"id"`
	BorrowingUserID uint                `json:"borrowing_user_id"`
	LendingUserID   uint                `json:"lending_user_id"`
	PostID          uint                `json:"post_id"`
	Status          model.RequestStatus `json:"status"`
	ActiveStatus    bool                `json:"active_status"`
	Post            interface{}         `json:"post"`
}

type GetMyLendingPostsResponse struct {
	ID              uint                `json:"id"`
	BorrowingUserID uint                `json:"borrowing_user_id"`
	LendingUserID   uint                `json:"lending_user_id"`
	PostID          uint                `json:"post_id"`
	Status          model.RequestStatus `json:"status"`
	ActiveStatus    bool                `json:"active_status"`
	Post            interface{}         `json:"post"`
	Borrower        string              `json:"borrower"`
}
type CreateBorrowingRequestInput struct {
	LendingUserID uint `json:"lending_user_id"`
	PostID        uint `json:"post_id"`
}
