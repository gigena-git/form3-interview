package message

import (
	"github.com/gigena-git/form3-interview/pkg/fetch"
)

/**
 * Account represents an account in the form3 org section.
 * See https://api-docs.form3.tech/api.html#organization-accounts for
 * more information about fields
 */
type AccountFetchResponse struct {
	Data  *fetch.AccountDataResponse `json:"data"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
}

// This type is invoked by the Create method to parse the account data to json,
// so that it can be sent to the rest api.
type AccountFetchRequest struct {
	Data *fetch.AccountDataRequest `json:"data"`
}
