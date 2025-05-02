package s3

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
)

type S3Client struct {
	KeyID       string
	BucketID    string
	AppKey      string
	AuthToken   string
	ApiUrl      string
	DownloadUrl string
	UploadUrl   string
	log         *slog.Logger
}

type AuthResponse struct {
	AuthToken   string `json:"authorizationToken"`
	DownloadUrl string `json:"downloadUrl"`
}

type UploadUrlResponse struct {
	AuthorizationToken string `json:"authorizationToken"`
	BucketId           string `json:"bucketId"`
	UploadUrl          string `json:"uploadUrl"`
}

func NewS3Client(keyID, bucketID, appKey, authToken, apiUrl, downloadUrl, uploadUrl string, log *slog.Logger) (*S3Client, error) {
	client := S3Client{
		KeyID:       keyID,
		BucketID:    bucketID,
		AppKey:      appKey,
		AuthToken:   authToken,
		ApiUrl:      apiUrl,
		DownloadUrl: downloadUrl,
		UploadUrl:   uploadUrl,
		log:         log,
	}

	return &client, nil
}

func (s3 *S3Client) GetS3Credentials() (AuthResponse, error) {
	url := fmt.Sprintf("%s/b2api/v2/b2_authorize_account", s3.ApiUrl)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return AuthResponse{}, err
	}
	req.SetBasicAuth(s3.KeyID, s3.AppKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return AuthResponse{}, err
	}
	defer resp.Body.Close()

	var authResp AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		return AuthResponse{}, err
	}

	return authResp, nil
}

func (s3 *S3Client) DownloadFile(bucketName, fileId string) ([]byte, error) {
	url := fmt.Sprintf("%s/file/%s/%s", s3.ApiUrl, bucketName, fileId)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", s3.AuthToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("download failed: %s", string(body))
	}

	return io.ReadAll(resp.Body)
}

func (s3 *S3Client) getUploadUrl() (string, string, error) {
	url := fmt.Sprintf("%s/b2api/v3/b2_get_upload_url?bucketId=%s", s3.ApiUrl, s3.BucketID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", "", err
	}

	req.Header.Set("Authorization", s3.AuthToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return "", "", fmt.Errorf("failed to get upload url: %s", body)
	}
	defer resp.Body.Close()

	var response UploadUrlResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", "", err
	}

	return response.UploadUrl, response.AuthorizationToken, nil
}

func (s3 *S3Client) UploadFile(filename string, fileData []byte) (string, error) {
	url, authToken, err := s3.getUploadUrl()
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(fileData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", authToken)
	req.Header.Set("X-Bz-File-Name", "main_photos/"+filename+".jpg")
	req.Header.Set("Content-Type", "image/jpeg")
	req.Header.Set("X-Bz-Content-Sha1", "do_not_verify")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("failed to upload file: %s", resp.Status)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	log.Println(result)

	return result["fileId"].(string), nil
}
