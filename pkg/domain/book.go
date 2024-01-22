package domain

import (
	"noob-server/pkg/models"
	"noob-server/pkg/types"
)

// for database Repository operation (call from service)
type IBookRepo interface {
	GetAllBooks() []models.BookDetail
	GetBook(bookID uint) (models.BookDetail, error)
	CreateBook(book *models.BookDetail) error
	UpdateBook(book *models.BookDetail) error
	DeleteBook(bookID uint) error
	DeleteBookByAuthorID(authorID uint) error
}

// for service operation (response to controller || call from controller)
type IBookService interface {
	GetAllBooks() ([]types.BookRequest, error)
	GetBook(bookID uint) (models.BookDetail, error)
	CreateBook(book *models.BookDetail) error
	UpdateBook(book *models.BookDetail) error
	DeleteBook(bookID uint) error
	DeleteBookByAuthorID(authorID uint) error
}
