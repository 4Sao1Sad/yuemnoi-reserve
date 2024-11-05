package util

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/KKhimmoon/yuemnoi-reserve/config"
)

func GetUserById(userId uint) (string, error) {
	cfg := config.Load()
	url := fmt.Sprintf("%s%d", cfg.UserInfoURL, userId)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Print(err)
		return "", err
	}

	// Set the header
	req.Header.Set("X-User-Id", strconv.Itoa(int(userId)))
	req.Header.Set("X-bypass-auth", "true")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
		return "", err
	}
	defer resp.Body.Close()

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
