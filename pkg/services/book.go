package services

import (
	"errors"
	"noob-server/pkg/domain"
	"noob-server/pkg/models"
)

// parent struct to implement interface binding
type bookService struct {
	bookRepo domain.IBookRepo
}

// BookServiceInstance interface binding
func BookServiceInstance(bookRepo domain.IBookRepo) domain.IBookService {
	return &bookService{
		bookRepo: bookRepo,
	}
}

func (service *bookService) GetAllBooks() ([]models.BookDetail, error) {
	book := service.bookRepo.GetAllBooks()
	if len(book) == 0 {
		return nil, errors.New("book not found")
	}
	return book, nil
}

func (service *bookService) GetBook(bookID uint) (models.BookDetail, error) {
	bookDetail, err := service.bookRepo.GetBook(bookID)

	if err != nil {
		return bookDetail, errors.New("book not found")
	}
	return bookDetail, nil
}

func (service *bookService) CreateBook(book *models.BookDetail) error {
	// _, err := service.authorRepo.GetAuthor(book.AuthorId)
	// if err != nil {
	// 	return errors.New("Author ID not found")
	// }
	if err := service.bookRepo.CreateBook(book); err != nil {
		return errors.New("BookDetail was not created")
	}
	return nil
}

func (service *bookService) UpdateBook(updatedBook *models.BookDetail) error {
	existingBook, err := service.GetBook(updatedBook.BookId)
	if err != nil {
		return errors.New("book not found")
	}
	if updatedBook.BookName == "" {
		updatedBook.BookName = existingBook.BookName
	}
	if updatedBook.AuthorId == 0 {
		updatedBook.AuthorId = existingBook.AuthorId
	}
	if updatedBook.Publication == "" {
		updatedBook.Publication = existingBook.Publication
	}
	if err := service.bookRepo.UpdateBook(updatedBook); err != nil {
		return errors.New("BookDetail update was unsuccessful")
	}
	return nil
}

// DeleteBook implements domain.IBookService.
func (service *bookService) DeleteBook(bookID uint) error {
	if err := service.bookRepo.DeleteBook(bookID); err != nil {
		return errors.New("BookDetail is not deleted")
	}
	return nil
}
