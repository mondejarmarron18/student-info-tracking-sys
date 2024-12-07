package domains

import (
	"database/sql"
	"time"
)

type Permission struct {
	Id          string       `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	DeletedAt   sql.NullTime `json:"deleted_at"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}
