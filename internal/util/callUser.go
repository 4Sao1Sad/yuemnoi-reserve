package util

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/KKhimmoon/yuemnoi-reserve/config"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CheckUserExists(borrowingUserId uint, lendingUserId uint) error {
	if err := checkSingleUser(borrowingUserId); err != nil {
		return status.Errorf(codes.NotFound, "borrowing user check failed: %v", err.Error())
	}
	if err := checkSingleUser(lendingUserId); err != nil {
		return status.Errorf(codes.NotFound, "lending user check failed: %v", err.Error())
	}
	return nil
}
func checkSingleUser(userId uint) error {
	cfg := config.Load()
	url := fmt.Sprintf("%s%d", cfg.UserInfoURL, userId)
	resp, err := http.Get(url)
	if err != nil {
		return status.Errorf(codes.Internal, "failed to call user service: %v", err)
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case http.StatusOK:
		var user map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
			return status.Errorf(codes.Internal, "failed to decode user data: %v", err)
		}
		if user["id"] == float64(0) {
			return fmt.Errorf("user with Id %d does not exist", userId)
		}
		return nil
	case http.StatusNotFound:
		return fmt.Errorf("user with Id %d does not exist", userId)
	default:
		return status.Errorf(codes.Unknown, "unexpected response from user service: status code %d", resp.StatusCode)
	}
}
