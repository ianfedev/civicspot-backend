package types

import "time"

// Auditable represents common fields shared across domain entities
// such as ID and timestamps for creation and updates.
type Auditable struct {
	ID        string    // ID is the unique identifier (UUID or ULID).
	CreatedAt time.Time // CreatedAt marks the entity creation time.
	UpdatedAt time.Time // UpdatedAt marks the last time the entity was modified.
}
