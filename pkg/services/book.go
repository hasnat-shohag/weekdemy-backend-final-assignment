package services

import (
	"errors"
	"noob-server/pkg/domain"
	"noob-server/pkg/models"
	"noob-server/pkg/types"
)

// parent struct to implement interface binding
type bookService struct {
	repo domain.IBookRepo
}

// interface binding
func BookServiceInstance(bookRepo domain.IBookRepo) domain.IBookService {
	return &bookService{
		repo: bookRepo,
	}
}

// all methods of interface are implemented
func (service *bookService) GetAllBooks() ([]types.BookRequest, error) {
	var allBooks []types.BookRequest
	book := service.repo.GetAllBooks()

	if len(book) == 0 {
		return nil, errors.New("no books found")
	}
	for _, val := range book {
		allBooks = append(allBooks, types.BookRequest{
			BookId:      val.BookId,
			BookName:    val.BookName,
			AuthorId:    val.AuthorId,
			Publication: val.Publication,
		})
	}
	return allBooks, nil
}

func (service *bookService) GetBook(bookID uint) (models.BookDetail, error) {
	bookDetail, err := service.repo.GetBook(bookID)
	if err != nil {
		return models.BookDetail{}, errors.New("no book found")
	}
	return bookDetail, nil
}

// CreateBook implements domain.IBookService.
func (service *bookService) CreateBook(book *models.BookDetail) error {
	if err := service.repo.CreateBook(book); err != nil {
		return errors.New("bookdetail is not created")
	}
	return nil
}

// UpdateBook implements domain.IBookService.
func (service *bookService) UpdateBook(book *models.BookDetail) error {
	if err := service.repo.UpdateBook(book); err != nil {
		return errors.New("bookdetail is not updated")
	}
	return nil
}

// DeleteBook implements domain.IBookService.
func (service *bookService) DeleteBook(bookID uint) error {
	if err := service.repo.DeleteBook(bookID); err != nil {
		return errors.New("bookdetail is not deleted")
	}
	return nil
}

// DeleteBookByAuthorID implements domain.IBookService
func (service *bookService) DeleteBookByAuthorID(authorID uint) error {
	if err := service.repo.DeleteBookByAuthorID(authorID); err != nil {
		return errors.New("bookdetail is not deleted")
	}
	return nil
}
