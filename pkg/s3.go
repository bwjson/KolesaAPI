package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

type S3Client struct {
	KeyID       string
	AppKey      string
	AuthToken   string
	DownloadUrl string
	log         *slog.Logger
}

type AuthResponse struct {
	AuthToken   string `json:"authorizationToken"`
	DownloadUrl string `json:"downloadUrl"`
}

func NewS3Client(keyID, appKey, authToken, downloadUrl string, log *slog.Logger) (*S3Client, error) {
	client := S3Client{
		KeyID:       keyID,
		AppKey:      appKey,
		AuthToken:   authToken,
		DownloadUrl: downloadUrl,
		log:         log,
	}

	return &client, nil
}

func (s3 *S3Client) GetS3Credentials() (*AuthResponse, error) {
	req, err := http.NewRequest("GET", "https://api.backblazeb2.com/b2api/v2/b2_authorize_account", nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(s3.KeyID, s3.AppKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var authResp AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		return nil, err
	}

	return &authResp, nil
}

func (s3 *S3Client) DownloadFile(bucketName, fileId string) ([]byte, error) {
	url := fmt.Sprintf("%s/file/%s/%s", s3.DownloadUrl, bucketName, fileId)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	s3.log.Info("Using Auth Token: %s", s3.AuthToken)
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
