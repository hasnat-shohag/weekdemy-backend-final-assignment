package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"noob-server/pkg/domain"
	"noob-server/pkg/models"
	"noob-server/pkg/types"
	"strconv"
)

type IAuthorController interface {
	GetAllAuthors(e echo.Context) error
	GetAuthor(e echo.Context) error
	CreateAuthor(e echo.Context) error
	UpdateAuthor(e echo.Context) error
	DeleteAuthor(e echo.Context) error
	DeleteBookByAuthorID(e echo.Context) error
}

type AuthorController struct {
	AuthorSvc domain.IAuthorService
	BookSvc   domain.IBookService
}

func NewAuthorController(AuthorSvc domain.IAuthorService, BookSvc domain.IBookService) AuthorController {
	return AuthorController{
		AuthorSvc: AuthorSvc,
		BookSvc:   BookSvc,
	}
}

// CreateAuthor implements IAuthorController.
func (controller *AuthorController) CreateAuthor(e echo.Context) error {
	reqAuthor := &types.AuthorRequest{}
	if err := e.Bind(reqAuthor); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	if err := reqAuthor.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	Author := &models.AuthorDetail{
		AuthorName:        reqAuthor.AuthorName,
		AuthorAddress:     reqAuthor.AuthorAddress,
		AuthorPhoneNumber: reqAuthor.AuthorPhoneNumber,
	}
	if err := controller.AuthorSvc.CreateAuthor(Author); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusCreated, "AuthorDetail is created successfully")
}

// GetAuthors implements IAuthorController.
func (controller *AuthorController) GetAllAuthors(e echo.Context) error {
	tempAuthorID := e.QueryParam("authorID")
	AuthorID, err := strconv.ParseInt(tempAuthorID, 0, 0)
	if err != nil && tempAuthorID != "" {
		return e.JSON(http.StatusBadRequest, "Enter a valid Author ID")
	}
	Author, err := controller.AuthorSvc.GetAuthor(uint(AuthorID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, Author)
}

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
func (controller *AuthorController) DeleteAuthor(e echo.Context) error {
	tempAuthorID := e.Param("authorID")
	AuthorID, err := strconv.ParseInt(tempAuthorID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	_, err = controller.AuthorSvc.GetAuthor(uint(AuthorID))
	_, err = controller.AuthorSvc.GetAuthor(uint(AuthorID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	if err := controller.AuthorSvc.DeleteAuthor(uint(AuthorID)); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	if err := controller.BookSvc.DeleteBookByAuthorID(uint(AuthorID)); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, "AuthorDetail is deleted successfully and All Books of the author deleted successfully")
}
