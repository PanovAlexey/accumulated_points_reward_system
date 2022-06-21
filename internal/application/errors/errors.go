package errors

import "errors"

var ErrOrderAlreadyExists = errors.New("already exists")
var ErrOrderAlreadySent = errors.New("already sent")
var ErrOrderNumberInvalid = errors.New("number is invalid")
var ErrUserAlreadyExists = errors.New("already exists")
