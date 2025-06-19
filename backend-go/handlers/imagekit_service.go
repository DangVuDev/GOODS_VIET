package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type ImageKitService struct {
	PrivateAPIKey string
	UploadURL     string
}

// UploadResponse là struct map với kết quả trả về từ ImageKit API
type UploadResponse struct {
	FileID string `json:"fileId"`
	URL    string `json:"url"`
	Name   string `json:"name"`
}

// NewImageKitService tạo service mới với Private API Key
func NewImageKitService(privateAPIKey string) *ImageKitService {
	return &ImageKitService{
		PrivateAPIKey: privateAPIKey,
		UploadURL:     "https://upload.imagekit.io/api/v1/files/upload",
	}
}

// UploadImage upload file local lên ImageKit
func (s *ImageKitService) UploadImage(filePath string, fileName string, folder string) (*UploadResponse, error) {
	client := resty.New()

	resp, err := client.R().
		SetBasicAuth(s.PrivateAPIKey, "").
		SetFile("file", filePath).
		SetFormData(map[string]string{
			"fileName": fileName,
			"folder":   folder,
		}).
		Post(s.UploadURL)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("Upload failed: %s - %s", resp.Status(), resp.String())
	}

	var uploadResp UploadResponse
	err = json.Unmarshal(resp.Body(), &uploadResp)
	if err != nil {
		return nil, err
	}

	return &uploadResp, nil
}
