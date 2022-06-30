package test

import (
	"testing"

	"github.com/gigena-git/form3-interview/accounts"
	"github.com/gigena-git/form3-interview/pkg/delete"
	"github.com/gigena-git/form3-interview/pkg/fetch"
)

func TestFetchAccount(t *testing.T) {
	create_request := generateCreateData()
	accounts.Create(create_request)
	delete_request := &delete.AccountDataRequest{
		ID:      create_request.ID,
		Version: create_request.Version,
	}
	defer accounts.Delete(delete_request)
	fetch_request := &fetch.AccountDataRequest{
		ID: create_request.ID,
	}
	res, err := accounts.Fetch(fetch_request)
	if err != nil {
		t.Fatalf("Test returned an error: %v", err.Error())
	}

	if create_request.ID != res.ID {
		t.Fatalf("Account ID is different to Account ID fetched")
	}
}

func TestFetchAccountNotFound(t *testing.T) {
	fetch_request := &fetch.AccountDataRequest{
		ID: "13b97dcd-eff2-4e30-ae0b-54629d70e56b",
	}

	_, err := accounts.Fetch(fetch_request)
	if err != nil {
		switch err.StatusCode {
		case "404":
		default:
			t.Fatalf("This test must return an error NotFoundError")
		}
	} else {
		t.Fatalf("This test must return an error NotFoundError")
	}
}
