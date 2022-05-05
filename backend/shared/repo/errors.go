package repo

import "errors"

var (
	ErrRecordNotFound    = errors.New("The thing you are trying to access wasn't found")
	ErrEmailTaken        = errors.New("User email is already used")
	ErrNotImplementedYet = errors.New("This feature is not implemented yet")
)
