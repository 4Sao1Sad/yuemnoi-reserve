package util

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/KKhimmoon/yuemnoi-reserve/config"
)

func GetUserById(userId uint) (string, error) {
	cfg := config.Load()
	url := fmt.Sprintf("%s%d", cfg.UserInfoURL, userId)
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to call user service: %v", err)
	}

	switch resp.StatusCode {
	case http.StatusOK:
		var user struct {
			Name string `json:"name"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
			return "", fmt.Errorf("failed to decode user data: %v", err)
		}

		return user.Name, nil

	case http.StatusNotFound:
		return "", fmt.Errorf("user with Id %d does not exist", userId)
	default:
		return "", fmt.Errorf("unexpected response from user service: status code %d", resp.StatusCode)
	}
}
