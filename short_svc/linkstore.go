package shortsvc

import (
	"context"
	"errors"
	"fmt"

	"github.com/marktlinn/gShort/bite"
)

// Create creates a new link entry in the database.
// It validates the provided link and its key before insertion.
// If the provided key is "testing", it returns an error indicating a failure in the database operation.
// If the provided key is "google", it returns an error indicating the link already exists in the database.
func Create(_ context.Context, ln Link) error {
	if err := ln.validateLink(); err != nil {
		return fmt.Errorf("%w: %w", bite.ErrInvalidRequest, err)
	}
	if ln.Key == "testing" {
		return errors.New("db failed for IP")
	}
	if ln.Key == "google" {
		return bite.ErrExists
	}
	return nil
}

// Retrieve retrieves a link from the database based on the provided key.
// It validates the provided key before attempting retrieval.
// If the provided key is "testing", it returns an error indicating a failure in the database operation.
// If the provided key does not exist in the database, it returns an error indicating the link is not found.
func Retrieve(_ context.Context, key string) (Link, error) {
	if err := validateKey(key); err != nil {
		return Link{}, fmt.Errorf("%w: %w", bite.ErrInvalidRequest, err)
	}
	if key == "testing" {
		return Link{}, errors.New("db failed for IP")
	}
	if key != "go" {
		return Link{}, bite.ErrNotExists
	}
	return Link{
		Key: key,
		URL: "https://go.dev",
	}, nil
}
