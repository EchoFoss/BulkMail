package campaign

import "errors"

var (
	nameIsRequiredErr    = errors.New("name is required and must not be empty")
	contentIsRequiredErr = errors.New("content is required and must not be empty")
	contactIsRequiredErr = errors.New("contact must be higher than 0")
)
