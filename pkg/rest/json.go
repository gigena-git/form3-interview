package rest

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gigena-git/form3-interview/pkg/create"
	"github.com/gigena-git/form3-interview/pkg/fetch"
	"github.com/gigena-git/form3-interview/pkg/message"
)

func ParseFetchResponse(res *http.Response) (*fetch.AccountDataResponse, *FetchAccountError) {
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Println(err.Error())
		return nil, &FetchAccountError{
			ErrorMessage: "An error occurred when trying to read the response body",
			StackTrace:   err.Error(),
			StatusCode:   strconv.Itoa(res.StatusCode),
		}
	}
	var data message.AccountFetchResponse
	json.Unmarshal(body, &data)
	return data.Data, nil
}

func ParseCreateRequest(account_data *create.AccountDataRequest) ([]byte, error) {
	requestParse := message.AccountCreateRequest{
		Data: account_data,
	}
	return json.Marshal(requestParse)
}

func ParseCreateResponse(res *http.Response) (*create.AccountDataResponse, *CreateAccountError) {
	body, err := ioutil.ReadAll(res.Body)
	//log.Println(string(body))
	defer res.Body.Close()
	if err != nil {
		log.Println(err.Error())
		return nil, &CreateAccountError{
			ErrorMessage: "An error occurred when trying to read the response body",
			StackTrace:   err.Error(),
			StatusCode:   strconv.Itoa(res.StatusCode),
		}
	}
	var data message.AccountCreateResponse
	json.Unmarshal(body, &data)
	return data.Data, nil
}

func ParseDeleteResponse(res *http.Response) *DeleteAccountError {
	_, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Println(err.Error())
		return &DeleteAccountError{
			ErrorMessage: "An error occurred when trying to read the response body",
			StackTrace:   err.Error(),
			StatusCode:   strconv.Itoa(res.StatusCode),
		}
	}
	return nil
}
