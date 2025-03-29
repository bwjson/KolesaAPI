package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"mime/multipart"
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

func (s3 *S3Client) UploadFile(bucketId, fileName string, fileData []byte) (string, error) {
	// сделать функцию для получения uploadUrl (получение bucketId через bucketName -> получение uploadUrl)
	uploadUrl := "https://pod-060-1000-05.backblaze.com/b2api/v2/b2_upload_file/4a61547d0966c63d99500f10/c006_v0601000_t0057"

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		return "", err
	}
	_, err = part.Write(fileData)
	if err != nil {
		return "", err
	}
	writer.Close()

	req, err := http.NewRequest("POST", uploadUrl, body)
	if err != nil {
		return "", err
	}
	// ПОМЕНЯТЬ ЧТОБЫ AUTHTOKEN получался через ручку, а не брался из конфига s3 потому что он будет обновляться
	req.Header.Set("Authorization", s3.AuthToken)

	// Додумать как давать названия для файлов
	req.Header.Set("Authorization", uploadUrl)
	req.Header.Set("X-Bz-File-Name", fileName)
	req.Header.Set("Content-Type", "b2/x-auto")
	req.Header.Set("X-Bz-Content-Sha1", "do_not_verify")
	req.Header.Set("X-Bz-Info-Author", "golang")
	req.Header.Set("Content-Length", fmt.Sprintf("%d", body.Len()))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		bodyResp, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("upload failed: %s", string(bodyResp))
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result["fileId"].(string), nil

	// после завершения добавления файла не забыть связать все в car_photos бд (получать car_id??)
}
