package test

import (
	"testing"

	"github.com/gigena-git/form3-interview/accounts"
	"github.com/gigena-git/form3-interview/pkg/delete"
)

func TestCreateAccount(t *testing.T) {
	create_request := generateCreateData()
	delete_request := &delete.AccountDataRequest{
		ID:      create_request.ID,
		Version: create_request.Version,
	}
	defer accounts.Delete(delete_request)
	_, err := accounts.Create(create_request)
	if err != nil {
		t.Fatalf("This test should not return an error: %v", err.Error())
	}
}

func TestCreateAccountButConflict(t *testing.T) {
	create_request := generateCreateData()
	accounts.Create(create_request)
	delete_request := &delete.AccountDataRequest{
		ID:      create_request.ID,
		Version: create_request.Version,
	}
	defer accounts.Delete(delete_request)

	_, err := accounts.Create(create_request)
	if err != nil {
		switch err.StatusCode {
		case "409":
		default:
			t.Fatalf("This test must return a Conflict status code")
		}
	} else {
		t.Fatalf("This test must return an error ConflictError")
	}
}
