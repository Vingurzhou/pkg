package errors

import "errors"

var (
	ErrNoRowsAffected = errors.New("info.RowsAffected == 0")
)
