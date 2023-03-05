package utils

import "errors"

var ErrInvalidCredentials = errors.New("invalid credentials")
var ErrInsertData = errors.New("error inserting data")
var ErrLoggedOut = errors.New("user logged out")
var ErrNotLogged = errors.New("invalid request")
var ErrAlreadyExist = errors.New("user registered")
