package repo

import "errors"

var (
	ErrRecordNotFound    = errors.New("Record not found")
	ErrEmailTaken        = errors.New("User email is already used")
	ErrNotImplementedYet = errors.New("This feature is not implemented yet")
)
