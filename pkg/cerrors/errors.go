package cerrors

import "errors"

var (
	ErrCEPNotFound = errors.New("can not find zipcode")
)
