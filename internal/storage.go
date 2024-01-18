package internal

import (
	"github.com/anujk14/go-playground/internal/models"
	"github.com/google/uuid"
)

type UserStore interface {
	GetUser(id uuid.UUID) (models.User, error)
	GetUsers() ([]models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(id uuid.UUID) error
}

type NoteStore interface {
	GetNote(id uuid.UUID) (models.Note, error)
	GetNotes(userIds ...uuid.UUID) ([]models.Note, error)
	CreateNote(note *models.Note) error
	UpdateNote(note *models.Note) error
	DeleteNote(id uuid.UUID) error
}
