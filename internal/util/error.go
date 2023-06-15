package util

import "errors"

type ErrorDomain struct {
	err     error
	Message string `json:"message"`
}

func (e ErrorDomain) Error() string {
	return e.err.Error()
}

func NewErrorDomain(errorMessage string) error {
	return &ErrorDomain{
		err:     errors.New(errorMessage),
		Message: errorMessage,
	}
}

func IsDomainError(e error) bool {
	_, err := e.(*ErrorDomain)
	return err
}
