package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/dgravesa/useby/pkg/useby"
)

func main() {
	var projectID string
	var username string

	flag.StringVar(&projectID, "projectID", "", "GCP project ID")
	flag.StringVar(&username, "username", "", "Name of user to delete")
	flag.Parse()

	// validate command line arguments
	errs := []string{}
	if projectID == "" {
		errs = append(errs, "projectID is required")
	}
	if username == "" {
		errs = append(errs, "username is required")
	}

	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Println(err)
		}
		os.Exit(1)
	}

	// initialize client
	client, err := useby.NewDatastoreClient(projectID)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := client.DeleteUser(context.Background(), username); err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println("user deleted successfully:", username)
	}
}
