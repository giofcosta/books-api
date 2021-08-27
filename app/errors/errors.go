package app

import "errors"

var ErrUnexpected = errors.New("Unexpected internal error")
var ErrBookNotFound = errors.New("Video not found")
