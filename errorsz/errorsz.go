package errorsz

import "errors"

var (
	ErrNoRowsAffected = errors.New("info.RowsAffected == 0")
	ErrInvalidMethod  = errors.New("invalid method")
)
