package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dgravesa/passdude/pkg/passdude"
)

func main() {
	var projectID string
	var username, password string

	flag.StringVar(&projectID, "projectID", "", "GCP project ID")
	flag.StringVar(&username, "username", "", "Name of user to create")
	flag.StringVar(&password, "password", "", "Password of user to create")
	flag.Parse()

	// validate command line arguments
	errs := []string{}
	if projectID == "" {
		errs = append(errs, "projectID is required")
	}
	if username == "" {
		errs = append(errs, "username is required")
	}
	if password == "" {
		errs = append(errs, "password is required")
	}

	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Println(err)
		}
		os.Exit(1)
	}

	// initialize client
	client, err := passdude.NewDatastoreClient(projectID)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	user, err := client.CreateUser(username, password)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println("new user created successfully:", user.Name)
	}
}
