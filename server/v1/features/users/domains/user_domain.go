package domains

import (
	"database/sql"
	"time"
)

type User struct {
	Id         string       `json:"id"`
	RoleId     string       `json:"roleId"`
	Email      string       `json:"email"`
	Password   string       `json:"password,omitempty"`
	VerifiedAt sql.NullTime `json:"verifiedAt"`
	DeletedAt  sql.NullTime `json:"deletedAt"`
	CreatedAt  time.Time    `json:"createdAt"`
	UpdatedAt  time.Time    `json:"updatedAt"`
}
