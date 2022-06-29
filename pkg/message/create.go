package message

import (
	"github.com/gigena-git/form3-interview/pkg/create"
)

/**
 * Account represents an account in the form3 org section.
 * See https://api-docs.form3.tech/api.html#organization-accounts for
 * more information about fields
 */
type AccountCreateResponse struct {
	Data  *create.AccountDataResponse `json:"data"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
}

// This type is invoked by the Create method to parse the account data to json,
// so that it can be sent to the rest api.
type AccountCreateRequest struct {
	Data *create.AccountDataRequest `json:"data"`
}
