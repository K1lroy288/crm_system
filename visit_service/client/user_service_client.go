package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"visit-service/model"
)

type UserServiceClient interface {
	GetUserInfo(lastname string) (*model.MasterDTO, error)
	GetMastersByIDs(mastersIDs []uint) ([]model.MasterDTO, error)
}

type userServiceClient struct {
	baseURL string
	client  *http.Client
}

func NewUserServiceClient(baseURL string) *userServiceClient {
	return &userServiceClient{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

func (u *userServiceClient) GetUserInfo(lastname string) (*model.MasterDTO, error) {
	url := fmt.Sprintf("%s/user/%s", u.baseURL, "some_info")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := u.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("user with such lastname not found")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("user service error: %d", resp.StatusCode)
	}

	var res *model.MasterDTO
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	return res, nil
}

func (u *userServiceClient) GetMastersByIDs(mastersIDs []uint) ([]model.MasterDTO, error) {
	url := fmt.Sprintf("%s/user/mastersByIDs", u.baseURL)
	postBody, err := json.Marshal(mastersIDs)
	if err != nil {
		return nil, err
	}
	reqBody := bytes.NewBuffer(postBody)

	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := u.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("user service error: %d", resp.StatusCode)
	}

	var res []model.MasterDTO
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	return res, nil
}
