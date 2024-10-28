package util

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func CheckUserExists(borrowingUserID uint, lendingUserID uint) error {
	if err := checkSingleUser(borrowingUserID); err != nil {
		return fmt.Errorf("borrowing user check failed: %w", err)
	}
	if err := checkSingleUser(lendingUserID); err != nil {
		return fmt.Errorf("lending user check failed: %w", err)
	}

	return nil
}
func checkSingleUser(userID uint) error {
	url := fmt.Sprintf("http://localhost:8080/user/%d", userID)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to call user service: %v", err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		var user map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
			return fmt.Errorf("failed to decode user data: %v", err)
		}

		if user["id"] == float64(0) {
			return fmt.Errorf("user with ID %d does not exist", userID)
		}

		return nil

	case http.StatusNotFound:
		return fmt.Errorf("user with ID %d does not exist", userID)

	default:
		return fmt.Errorf("unexpected response from user service: status code %d", resp.StatusCode)
	}
}
