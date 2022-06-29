package rest

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
)

var host = getHost()
var api_path = "/v1/organisation/accounts"

//If the variable SERVER_PATH is not set, assume it is localhost
func getHost() string {
	server_path := os.Getenv("SERVER_PATH")
	if server_path == "" {
		return "http://localhost:8080"
	}
	log.Printf("SERVER_PATH has been provided: %s", server_path)
	return server_path
}

func accountFetchUrl(id string) (*url.URL, error) {
	return url.Parse(fmt.Sprintf("%s%s/%s", host, api_path, id))
}

func accountCreateUrl() (*url.URL, error) {
	return url.Parse(fmt.Sprintf("%s%s", host, api_path))
}

func AccountDeleteUrl(id string, version int64) (*url.URL, error) {
	url_path, _ := url.Parse(fmt.Sprintf("%s%s/%s", host, api_path, id))
	params := url.Values{}
	params.Add("version", strconv.FormatInt(version, 10))

	url_path.RawQuery = params.Encode()
	return url_path, nil
}
