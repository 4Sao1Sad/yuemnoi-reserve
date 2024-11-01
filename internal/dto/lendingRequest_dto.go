package dto

import "github.com/KKhimmoon/yuemnoi-reserve/internal/model"

type GetMyLendingRequestsResponse struct {
	ID              uint                `json:"id"`
	BorrowingUserID uint                `json:"borrowing_user_id"`
	LendingUserID   uint                `json:"lending_user_id"`
	BorrowingPostID uint                `json:"borrowing_post_id"`
	LendingPostID   uint                `json:"lending_post_id"`
	Status          model.RequestStatus `json:"status"`
	ActiveStatus    bool                `json:"active_status"`
	LendingPost     interface{}         `json:"lending_post"`
	BorrowingPost   interface{}         `json:"borrowing_post"`
}
type GetMyBorrowingPostsResponse struct {
	ID              uint                `json:"id"`
	BorrowingUserID uint                `json:"borrowing_user_id"`
	LendingUserID   uint                `json:"lending_user_id"`
	BorrowingPostID uint                `json:"borrowing_post_id"`
	LendingPostID   uint                `json:"lending_post_id"`
	Status          model.RequestStatus `json:"status"`
	ActiveStatus    bool                `json:"active_status"`
	LendingPost     interface{}         `json:"lending_post"`
	BorrowingPost   interface{}         `json:"borrowing_post"`
}

type CreateLendingRequestInput struct {
	BorrowingUserID uint `json:"borrowing_user_id"`
	BorrowingPostID uint `json:"borrowing_post_id"`
	LendingPostID   uint `json:"lending_post_id"`
}
