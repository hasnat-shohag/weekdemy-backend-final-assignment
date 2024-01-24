package repositories

import (
	"gorm.io/gorm"
	"noob-server/pkg/domain"
	"noob-server/pkg/models"
)

type authorRepo struct {
	db *gorm.DB
}

// interface binding
func AuthorDBInstance(d *gorm.DB) domain.IAuthorRepo {
	return &authorRepo{
		db: d,
	}
}

// all methods of interface are implemented here
func (repo *authorRepo) GetAllAuthors() []models.AuthorDetail {
	var author []models.AuthorDetail

	err := repo.db.Find(&author).Error

	if err != nil {
		return []models.AuthorDetail{}
	}
	return author
}

func (repo *authorRepo) GetAuthor(authorID uint) (models.AuthorDetail, error) {
	var author models.AuthorDetail
	if err := repo.db.Where("author_id = ?", authorID).First(&author).Error; err != nil {
		return author, err
	}
	return author, nil
}

func (repo *authorRepo) CreateAuthor(author *models.AuthorDetail) error {
	err := repo.db.Create(author).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *authorRepo) UpdateAuthor(author *models.AuthorDetail) error {
	err := repo.db.Save(author).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *authorRepo) DeleteAuthor(authorID uint) error {
	var Author models.AuthorDetail
	if err := repo.db.Where("author_id = ?", authorID).Delete(&Author).Error; err != nil {
		return err
	}
	return nil
}
