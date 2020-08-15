package useby

import (
	"context"

	"cloud.google.com/go/datastore"
)

type datastoreUserStore struct {
	projectID string
}

type userCredentials struct {
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

// PutUser creates user credentials and stores in datastorex
func (dus *datastoreUserStore) PutUser(ctx context.Context, username, password string) error {
	client, err := datastore.NewClient(ctx, dus.projectID)
	if err != nil {
		return err
	}

	var credentials userCredentials
	key := datastore.NameKey(datastoreUserKind, username, nil)

	// verify credentials do not already exist for user
	err = client.Get(ctx, key, &credentials)
	if err != datastore.ErrNoSuchEntity {
		return err
	}

	// create salt and hashed password
	credentials.Salt = makeSalt()
	credentials.HashPass = applySaltAndHash(password, credentials.Salt)

	// push new user
	_, err = client.Put(ctx, key, &credentials)
	return err
}

// DeleteUser deletes a user with the given username
func (dus *datastoreUserStore) DeleteUser(ctx context.Context, username string) error {
	client, err := datastore.NewClient(ctx, dus.projectID)
	if err != nil {
		return err
	}

	key := datastore.NameKey(datastoreUserKind, username, nil)
	return client.Delete(ctx, key)
}

// Authenticate validates user credentials against those in datastore
func (dus *datastoreUserStore) Authenticate(ctx context.Context, username, password string) error {
	client, err := datastore.NewClient(ctx, dus.projectID)
	if err != nil {
		return err
	}

	var credentials userCredentials
	key := datastore.NameKey(datastoreUserKind, username, nil)

	// get hashed credentials from datastore
	err = client.Get(ctx, key, &credentials)
	if err != nil {
		return err
	}

	// validate user password
	reqHashPass := applySaltAndHash(password, credentials.Salt)
	if reqHashPass != credentials.HashPass {
		return ErrInvalidLogin
	}

	return nil
}
