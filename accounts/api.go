package accounts

/*
 * This file is the entry point of the library. Any external app invoking this library should either
 * call Fetch, Create, or Delete.
 */
import (
	"log"
	"strconv"

	"github.com/gigena-git/form3-interview/pkg/create"
	"github.com/gigena-git/form3-interview/pkg/delete"
	"github.com/gigena-git/form3-interview/pkg/fetch"
	"github.com/gigena-git/form3-interview/pkg/rest"
)

func Fetch(fetch_request *fetch.AccountDataRequest) (*fetch.AccountDataResponse, *rest.FetchAccountError) {
	res, err := rest.FetchRequest(fetch_request.ID)
	if err != nil {
		log.Println(err.Error())
		return nil, &rest.FetchAccountError{
			ErrorMessage: "An error occurred when trying to parse fetch request to json",
			StackTrace:   err.Error(),
			StatusCode:   "N/A",
		}
	}
	if res.StatusCode == 404 {
		return nil, &rest.FetchAccountError{
			ErrorMessage: "Account not Found!",
			StackTrace:   "N/A",
			StatusCode:   strconv.Itoa(res.StatusCode),
		}
	}

	return rest.ParseFetchResponse(res)
}

func Create(create_request *create.AccountDataRequest) (*create.AccountDataResponse, *rest.CreateAccountError) {
	request_json, err := rest.ParseCreateRequest(create_request)
	//log.Println(string(json_data))
	if err != nil {
		log.Println(err.Error())
		return nil, &rest.CreateAccountError{
			ErrorMessage: "An error occurred when trying to parse create request to json",
			StackTrace:   err.Error(),
			StatusCode:   "N/A",
		}
	}
	res, err := rest.CreateRequest(request_json)
	if err != nil {
		log.Println(err.Error())
		return nil, &rest.CreateAccountError{
			ErrorMessage: "Internal Server Error",
			StackTrace:   err.Error(),
			StatusCode:   strconv.Itoa(res.StatusCode),
		}
	}
	if res.StatusCode != 201 {
		return nil, &rest.CreateAccountError{
			ErrorMessage: "Status code is not 201",
			StackTrace:   "N/A",
			StatusCode:   strconv.Itoa(res.StatusCode),
		}
	}

	return rest.ParseCreateResponse(res)
}

func Delete(delete_request *delete.AccountDataRequest) *rest.DeleteAccountError {
	res, err := rest.DeleteRequest(delete_request.ID, *delete_request.Version)
	if err != nil {
		log.Println(err.Error())
		return &rest.DeleteAccountError{
			ErrorMessage: "Internal Server Error",
			StackTrace:   err.Error(),
			StatusCode:   strconv.Itoa(res.StatusCode),
		}
	}
	if res.StatusCode != 204 {
		return &rest.DeleteAccountError{
			ErrorMessage: "Status code is not 204",
			StackTrace:   "N/A",
			StatusCode:   strconv.Itoa(res.StatusCode),
		}
	}

	return rest.ParseDeleteResponse(res)
}
