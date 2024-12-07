package domains

import (
	"database/sql"
	"time"
)

type User struct {
	Id         string       `json:"id"`
	RoleId     string       `json:"role_id"`
	Email      string       `json:"email"`
	Password   string       `json:"password,omitempty"`
	VerifiedAt sql.NullTime `json:"verified_at"`
	DeletedAt  sql.NullTime `json:"deleted_at"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
}
