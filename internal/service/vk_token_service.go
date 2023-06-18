package service

import (
	"encoding/json"
	"fmt"
	"friendly/internal/utils"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type VkService struct {
	ApiURL     string
	ApiVersion string
}

type VkResponse struct {
	VkResult []struct {
		ID              int64  `json:"id"`
		FirstName       string `json:"first_name"`
		LastName        string `json:"last_name"`
		CanAccessClosed bool   `json:"can_access_closed"`
		IsClosed        bool   `json:"is_closed"`
		City            struct {
			ID    int    `json:"id"`
			Title string `json:"title"`
		} `json:"city"`
	} `json:"response"`
}

func (r *VkResponse) ToClaims() map[string]interface{} {
	var result = make(map[string]interface{})

	result["first_name"] = r.VkResult[0].FirstName
	result["id"] = r.VkResult[0].ID
	result["can-access-closed"] = r.VkResult[0].CanAccessClosed
	result["is_closed"] = r.VkResult[0].IsClosed
	result["last_name"] = r.VkResult[0].LastName
	result["city_title"] = r.VkResult[0].City.Title
	result["city_id"] = r.VkResult[0].City.ID

	return result
}

func NewVkService(apiURL string, apiVersion string) *VkService {
	return &VkService{ApiURL: apiURL, ApiVersion: apiVersion}
}

func (s *VkService) GetClaims(vkAccessToken string) (map[string]interface{}, error) {
	fields := "city,first_name,last_name,id"

	url := fmt.Sprintf("%s?access_token=%s&v=%s&fields=%s", s.ApiURL, vkAccessToken, s.ApiVersion, fields)

	resp, err := http.Get(url)
	if err != nil {
		logrus.Error("VkService.GetClaims error :", err)
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Error("VkService.GetClaims error :", err)
		return nil, err
	}

	var result VkResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		utils.Log("", err)
		return nil, err
	}

	logrus.Info(result)

	resp.Body.Close()

	return result.ToClaims(), nil
}
