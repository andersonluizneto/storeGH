package errors

import "errors"

var (
	ErrCustomerNotFound = errors.New("costumer not found")
	ErrEmailAlreadyUsed = errors.New("email already exists")
	ErrCPFAlreadyUsed   = errors.New("cpf already exists")
	ErrInvalidUserName  = errors.New("invalid user name")
	ErrInvalidCPF       = errors.New("invalid user cpf")
	ErrInvalidEmail     = errors.New("invalid user email")
	ErrInvalidPhone     = errors.New("invalid user phone")
)
