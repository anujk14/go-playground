package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID    uuid.UUID `db:"id"`
	Name  string    `db:"name"`
	Email string    `db:"email"`
}

type Note struct {
	ID        uuid.UUID `db:"id"`
	Title     string    `db:"title"`
	Content   string    `db:"content"`
	CreatedBy uuid.UUID `db:"created_by"`
}
