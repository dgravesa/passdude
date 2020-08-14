package passdude

import "errors"

// ErrInvalidLogin is returned by Authenticate when a user provides invalid credentials
var ErrInvalidLogin = errors.New("passdude: invalid user credentials")
