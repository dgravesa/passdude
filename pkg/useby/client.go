package useby

import "context"

// Client is used to manage and authenticate users
type Client interface {
	PutUser(ctx context.Context, username, password string) error
	Authenticate(ctx context.Context, username, password string) error
	DeleteUser(ctx context.Context, username string) error
}
