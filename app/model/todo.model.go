package model

import (
	"github.com/google/uuid"
	"time"
)

type Todo struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title" validate:"required"`
	Completed bool      `json:"completed" validate:""`
	Date      time.Time `json:"date"`
}
