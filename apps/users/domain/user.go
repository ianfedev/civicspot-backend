package domain

import (
	types "github.com/ianfedev/civicspot-backend/pkg/common/domain"
	"time"
)

// DocumentType represents the type of national document used by the user.
type DocumentType string

const (
	// CC represents "Cédula de Ciudadanía".
	CC DocumentType = "CC"
	// TI represents "Tarjeta de Identidad".
	TI DocumentType = "TI"
	// CE represents "Cédula de Extranjería".
	CE DocumentType = "CE"
)

// User contains personal and geographic information of a system-registered citizen or official.
// It does not manage authentication or authorization concerns.
type User struct {
	types.Auditable
	ID           string       // ID is the unique identifier of the user (UUID or ULID).
	FirstName    string       // FirstName is the user's given name.
	LastName     string       // LastName is the user's family name.
	DocumentType DocumentType // DocumentType specifies the type of identification (CC, TI, etc.).
	DocumentID   string       // DocumentID is the actual identification number (e.g., cédula).
	City         string       // City is the city of residence of the user.
	State        string       // State refers to the broader region or administrative division.
	Address      string       // Address is the detailed location within the city (e.g., street address).
	ProfilePhoto *string      // ProfilePhoto contains the URL to the user's profile picture. It is optional.
	CreatedAt    time.Time    // CreatedAt records the timestamp when the user was first registered.
}
