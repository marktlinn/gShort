package bite

import "errors"

var (
	ErrInternal       = errors.New("internal server error: try again")
	ErrExists         = errors.New("already exists")
	ErrInvalidRequest = errors.New("invali request")
	ErrNotExists      = errors.New("does not exist")
)
