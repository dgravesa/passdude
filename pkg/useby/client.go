package useby

// Client is used to manage and authenticate users
type Client interface {
	CreateUser(username, password string) (*User, error)
	Authenticate(username, password string) (*User, error)
	DeleteUser(username string) error
}
