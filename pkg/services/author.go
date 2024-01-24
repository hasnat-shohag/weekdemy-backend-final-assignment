package services

import (
	"errors"
	"noob-server/pkg/domain"
	"noob-server/pkg/models"
	"noob-server/pkg/types"
)

type authorService struct {
	authorRepo domain.IAuthorRepo
	bookRepo   domain.IBookRepo
}

func AuthorServiceInstance(authorRepo domain.IAuthorRepo, bookRepo domain.IBookRepo) domain.IAuthorService {
	return &authorService{
		authorRepo: authorRepo,
		bookRepo:   bookRepo,
	}
}

func (service *authorService) GetAllAuthors() ([]types.AuthorRequest, error) {
	var allAuthors []types.AuthorRequest
	author := service.authorRepo.GetAllAuthors()
	if len(author) == 0 {
		return nil, errors.New("No author found")
	}
	for _, val := range author {
		allAuthors = append(allAuthors, types.AuthorRequest{
			AuthorID:          val.AuthorId,
			AuthorName:        val.AuthorName,
			AuthorAddress:     val.AuthorAddress,
			AuthorPhoneNumber: val.AuthorPhoneNumber,
		})
	}
	return allAuthors, nil
}

func (service *authorService) GetAuthor(authorID uint) (types.AuthorRequest, error) {
	authorDetail, err := service.authorRepo.GetAuthor(authorID)
	author := types.AuthorRequest{
		AuthorID:          authorDetail.AuthorId,
		AuthorName:        authorDetail.AuthorName,
		AuthorAddress:     authorDetail.AuthorAddress,
		AuthorPhoneNumber: authorDetail.AuthorPhoneNumber,
	}
	if err != nil {
		return author, errors.New("Author not found")
	}
	return author, nil
}

func (service *authorService) CreateAuthor(author *models.AuthorDetail) error {
	if err := service.authorRepo.CreateAuthor(author); err != nil {
		return errors.New("Author was not created")
	}
	return nil
}

func (service *authorService) UpdateAuthor(updatedAuthor *models.AuthorDetail) error {
	existingAuthor, err := service.GetAuthor(updatedAuthor.AuthorId)
	if err != nil {
		return errors.New("No author found")
	}
	if updatedAuthor.AuthorName == "" {
		updatedAuthor.AuthorName = existingAuthor.AuthorName
	}
	if updatedAuthor.AuthorAddress == "" {
		updatedAuthor.AuthorAddress = existingAuthor.AuthorAddress
	}
	if updatedAuthor.AuthorPhoneNumber == "" {
		updatedAuthor.AuthorPhoneNumber = existingAuthor.AuthorPhoneNumber
	}
	if err := service.authorRepo.UpdateAuthor(updatedAuthor); err != nil {
		return errors.New("Author was not updated")
	}
	return nil
}

func (service *authorService) DeleteAuthor(authorID uint) error {
	if err := service.authorRepo.DeleteAuthor(authorID); err != nil {
		return errors.New("Author was not deleted")
	}
	if err := service.bookRepo.DeleteBooksByAuthorID(authorID); err != nil {
		return errors.New("Author was not deleted")
	}
	return nil
}
