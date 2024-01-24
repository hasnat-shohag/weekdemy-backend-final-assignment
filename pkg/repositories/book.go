package repositories

import (
	"noob-server/pkg/domain"
	"noob-server/pkg/models"

	"gorm.io/gorm"
)

// parent struct to implement interface binding
type bookRepo struct {
	db *gorm.DB
}

// interface binding
func BookDBInstance(d *gorm.DB) domain.IBookRepo {
	return &bookRepo{
		db: d,
	}
}

func (repo *bookRepo) GetAllBooks() []models.BookDetail {
	var book []models.BookDetail
	err := repo.db.Find(&book).Error

	if err != nil {
		return []models.BookDetail{}
	}
	return book
}

func (repo *bookRepo) GetBook(bookID uint) (models.BookDetail, error) {
	var book models.BookDetail
	if err := repo.db.Where("book_id = ?", bookID).First(&book).Error; err != nil {
		return book, err
	}
	return book, nil
}

func (repo *bookRepo) CreateBook(book *models.BookDetail) error {
	err := repo.db.Create(book).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *bookRepo) UpdateBook(book *models.BookDetail) error {
	err := repo.db.Save(book).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *bookRepo) DeleteBook(bookID uint) error {
	var Book models.BookDetail
	if err := repo.db.Where("book_id = ?", bookID).Delete(&Book).Error; err != nil {
		return err
	}
	return nil
}

func (repo *bookRepo) DeleteBooksByAuthorID(authorID uint) error {
	var Book models.BookDetail
	if err := repo.db.Where("author_id = ?", authorID).Delete(&Book).Error; err != nil {
		return err
	}
	return nil
}
