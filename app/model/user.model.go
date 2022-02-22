package model

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID    `json:"id"`
	Username  string       `json:"username" validate:"required"`
	Email     string       `json:"email" validate:"required,email"`
	Password  string       `json:"password" validate:"required"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}
