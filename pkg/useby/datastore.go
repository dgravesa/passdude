package useby

import (
	"context"

	"cloud.google.com/go/datastore"
)

type datastoreUserStore struct {
	projectID string
}

type datastoreUser struct {
	User
	Salt     string
	HashPass string
}

const datastoreUserKind = "useby/users"

// NewDatastoreClient creates a new GCP Datastore user store instance
func NewDatastoreClient(projectID string) (Client, error) {
	dus := new(datastoreUserStore)
	dus.projectID = projectID
	return dus, nil
}

// CreateUser creates a user with the username and hashed password and stores in datastore
func (dus *datastoreUserStore) CreateUser(username, password string) (*User, error) {
	client, err := datastore.NewClient(context.Background(), dus.projectID)
	if err != nil {
		return nil, err
	}

	var user datastoreUser
	key := datastore.NameKey(datastoreUserKind, username, nil)

	// verify user does not already exist
	if err := client.Get(context.Background(), key, &user); err != datastore.ErrNoSuchEntity {
		return nil, err
	}

	// create salt and hashed password
	user.Name = username
	user.Salt = makeSalt()
	user.HashPass = applySaltAndHash(password, user.Salt)

	// push new user
	if _, err := client.Put(context.Background(), key, &user); err != nil {
		return nil, err
	}

	return &user.User, nil
}

// DeleteUser deletes a user with the given username
func (dus *datastoreUserStore) DeleteUser(username string) error {
	client, err := datastore.NewClient(context.Background(), dus.projectID)
	if err != nil {
		return err
	}

	key := datastore.NameKey(datastoreUserKind, username, nil)
	return client.Delete(context.Background(), key)
}

func (dus *datastoreUserStore) Authenticate(username, password string) (*User, error) {
	client, err := datastore.NewClient(context.Background(), dus.projectID)
	if err != nil {
		return nil, err
	}

	var user datastoreUser
	key := datastore.NameKey(datastoreUserKind, username, nil)

	// get user from datastore
	err = client.Get(context.Background(), key, &user)
	if err != nil {
		return nil, err
	}

	// validate user password
	reqHashPass := applySaltAndHash(password, user.Salt)
	if reqHashPass != user.HashPass {
		return nil, ErrInvalidLogin
	}

	return &user.User, nil
}
