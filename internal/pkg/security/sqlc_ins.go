package security

import (
	"database/sql"

	"github.com/google/uuid"
)

// ToNullString converts a string to sql.NullString.
func ToNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: s, Valid: true}
}

func ToNullUUID(i uuid.UUID) uuid.NullUUID {
	if i == uuid.Nil {
		return uuid.NullUUID{}
	}
	return uuid.NullUUID{UUID: i, Valid: true}
}
