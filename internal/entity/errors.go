package entutySend

import "errors"

var ErrNotValidationEmail = errors.New("email is empty")
var ErrNotValidationText = errors.New("the letter is empty")
