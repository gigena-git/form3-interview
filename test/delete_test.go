package test

import (
	"testing"

	"github.com/gigena-git/form3-interview/accounts"
	"github.com/gigena-git/form3-interview/pkg/delete"
)

func TestDeleteAccount(t *testing.T) {
	create_request := generateCreateData()
	accounts.Create(create_request)
	delete_request := &delete.AccountDataRequest{
		ID:      create_request.ID,
		Version: create_request.Version,
	}
	err := accounts.Delete(delete_request)
	if err != nil {
		t.Fatalf("This test should not return an error: %v", err.Error())
	}
}

func TestDeleteAccountNotFound(t *testing.T) {
	id := "13b97dcd-eff2-4e30-ae0b-54629d70e56b"
	version := int64(0)
	delete_request := &delete.AccountDataRequest{
		ID:      id,
		Version: &version,
	}

	err := accounts.Delete(delete_request)
	if err != nil {
		switch err.StatusCode {
		case "404":
		default:
			t.Fatalf("Status code should be 404!")
		}
	} else {
		t.Fatalf("Status Code should be 404!")
	}
}
