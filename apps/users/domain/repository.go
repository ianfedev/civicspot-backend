package domain

import "context"

// UserRepository defines access methods for storing and retrieving users.
type UserRepository interface {

	// GetByID returns the user with the given ID, or an error if not found.
	GetByID(ctx context.Context, id string) (*User, error)

	// GetByDocument returns the user with document type and ID, or nil if not found.
	GetByDocument(ctx context.Context, docType DocumentType, docID string) (*User, error)

	// Create persists a new user in the system.
	Create(ctx context.Context, user *User) error

	// Deactivate marks the user as inactive or soft-deleted.
	Deactivate(ctx context.Context, id string) error
}
