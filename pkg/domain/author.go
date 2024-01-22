package domain

import (
	"noob-server/pkg/models"
	"noob-server/pkg/types"
)

// IAuthorRepo for database Repository operation (call from service)
type IAuthorRepo interface {
	GetAllAuthors() []models.AuthorDetail
	GetAuthor(authorID uint) (models.AuthorDetail, error)
	CreateAuthor(author *models.AuthorDetail) error
	UpdateAuthor(author *models.AuthorDetail) error
	DeleteAuthor(authorID uint) error
}

// IAuthorService for service operation (response to controller || call from controller)
type IAuthorService interface {
	GetAllAuthors() ([]types.AuthorRequest, error)
	GetAuthor(authorID uint) (types.AuthorRequest, error)
	CreateAuthor(author *models.AuthorDetail) error
	UpdateAuthor(author *models.AuthorDetail) error
	DeleteAuthor(authorID uint) error
}
