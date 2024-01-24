package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"noob-server/pkg/domain"
	"noob-server/pkg/models"
	"noob-server/pkg/types"
	"strconv"
)

// IAuthorController is an interface for AuthorController.
type AuthorController struct {
	AuthorSvc domain.IAuthorService
}

// NewAuthorController returns a new AuthorController.
func NewAuthorController(AuthorSvc domain.IAuthorService) AuthorController {
	return AuthorController{
		AuthorSvc: AuthorSvc,
	}
}

// CreateAuthor implements IAuthorController.
func (authorService *AuthorController) CreateAuthor(e echo.Context) error {
	authorRequest := &types.AuthorRequest{}
	if err := e.Bind(authorRequest); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	if err := authorRequest.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	author := &models.AuthorDetail{
		AuthorName:        authorRequest.AuthorName,
		AuthorAddress:     authorRequest.AuthorAddress,
		AuthorPhoneNumber: authorRequest.AuthorPhoneNumber,
	}
	if err := authorService.AuthorSvc.CreateAuthor(author); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusCreated, "Author was created successfully")
}

// GetAllAuthors implements IAuthorController.
func (authorService *AuthorController) GetAllAuthors(e echo.Context) error {
	authors, err := authorService.AuthorSvc.GetAllAuthors()
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, authors)
}

// GetAuthor implements IAuthorController.
func (authorService *AuthorController) GetAuthor(e echo.Context) error {
	tempAuthorID := e.Param("authorID")
	authorID, err := strconv.ParseInt(tempAuthorID, 0, 0)
	if err != nil && tempAuthorID != "" {
		return e.JSON(http.StatusBadRequest, "Enter a valid author ID")
	}
	author, err := authorService.AuthorSvc.GetAuthor(uint(authorID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, author)
}

// UpdateAuthor implements IAuthorController.
func (authorService *AuthorController) UpdateAuthor(e echo.Context) error {
	authorRequest := &types.AuthorRequest{}
	if err := e.Bind(authorRequest); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	tempAuthorID := e.Param("authorID")
	authorID, err := strconv.ParseInt(tempAuthorID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Enter a valid author ID")
	}
	updatedAuthor := &models.AuthorDetail{
		AuthorId:          uint(authorID),
		AuthorName:        authorRequest.AuthorName,
		AuthorAddress:     authorRequest.AuthorAddress,
		AuthorPhoneNumber: authorRequest.AuthorPhoneNumber,
	}
	if err := authorService.AuthorSvc.UpdateAuthor(updatedAuthor); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, "Author was updated successfully")
}

// DeleteAuthor implements IAuthorController.
func (authorService *AuthorController) DeleteAuthor(e echo.Context) error {
	tempAuthorID := e.Param("authorID")
	authorID, err := strconv.ParseInt(tempAuthorID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	_, err = authorService.AuthorSvc.GetAuthor(uint(authorID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	if err := authorService.AuthorSvc.DeleteAuthor(uint(authorID)); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, "Author was deleted successfully")
}
