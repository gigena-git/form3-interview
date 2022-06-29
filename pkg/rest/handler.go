package rest

import (
	"bytes"
	"net/http"
)

func FetchRequest(id string) (*http.Response, error) {
	url_path, _ := accountFetchUrl(id)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url_path.String(), nil)
	req.Header.Set("Content-Type", "application/vnd.api+json")
	return client.Do(req)
}

func CreateRequest(json_data []byte) (*http.Response, error) {
	url_path, _ := accountCreateUrl()
	client := &http.Client{}
	req, _ := http.NewRequest("POST", url_path.String(), bytes.NewBuffer(json_data))
	req.Header.Set("Content-Type", "application/vnd.api+json")
	return client.Do(req)
}

func DeleteRequest(id string, version int64) (*http.Response, error) {
	url_path, _ := AccountDeleteUrl(id, version)

	client := &http.Client{}
	req, _ := http.NewRequest("DELETE", url_path.String(), nil)
	return client.Do(req)
}
