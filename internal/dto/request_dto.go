package dto

type ActiveRequestResponse struct {
	ID              uint        `json:"id"`
	BorrowingUserID uint        `json:"borrowing_user_id"`
	LendingUserID   uint        `json:"lending_user_id"`
	PostID          uint        `json:"post_id"`
	Role            string      `json:"role"`
	Post            interface{} `json:"post"`
}

type HistoryRequestResponse struct {
	ID              uint        `json:"id"`
	BorrowingUserID uint        `json:"borrowing_user_id"`
	LendingUserID   uint        `json:"lending_user_id"`
	PostID          uint        `json:"post_id"`
	IsReject        bool        `json:"is_reject"`
	Post            interface{} `json:"post"`
}
