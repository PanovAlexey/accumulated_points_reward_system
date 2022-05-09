package errors

import "errors"

var ErrorOrderAlreadyExists = errors.New("already exists")
var ErrorOrderAlreadySent = errors.New("already sent")
var ErrorOrderNumberInvalid = errors.New("number is invalid")
var ErrorUserAlreadyExists = errors.New("already exists")
