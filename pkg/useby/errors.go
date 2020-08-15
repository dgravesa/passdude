package useby

import "errors"

// ErrInvalidLogin is returned by Authenticate when a user provides invalid credentials
var ErrInvalidLogin = errors.New("useby: invalid user credentials")
