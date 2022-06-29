package rest

/*
 * This file contains error types, as recommended by https://go.dev/blog/error-handling-and-go
 */
import (
	"fmt"
)

type error interface {
	Error() string
}

type FetchAccountError struct {
	ErrorMessage string
	StackTrace   string
	StatusCode   string
}

func (e *FetchAccountError) Error() string {
	return fmt.Sprintf("An error occurred while retrieving an account. StatusCode: %v\nErrorMessage: %v\n StackTrace: %v", e.StatusCode, e.ErrorMessage, e.StackTrace)
}

type CreateAccountError struct {
	ErrorMessage string
	StackTrace   string
	StatusCode   string
}

func (e *CreateAccountError) Error() string {
	return fmt.Sprintf("An error occurred while creating an account. StatusCode: %v\nErrorMessage: %v\n StackTrace: %v", e.StatusCode, e.ErrorMessage, e.StackTrace)
}

type DeleteAccountError struct {
	ErrorMessage string
	StackTrace   string
	StatusCode   string
}

func (e *DeleteAccountError) Error() string {
	return fmt.Sprintf("An error occurred while deleting an account. StatusCode: %v\nErrorMessage: %v\n StackTrace: %v", e.StatusCode, e.ErrorMessage, e.StackTrace)
}
