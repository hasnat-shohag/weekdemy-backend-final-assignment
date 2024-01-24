package controllers

import (
	"net/http"
	"noob-server/pkg/domain"
	"noob-server/pkg/models"
	"noob-server/pkg/types"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookController struct {
	bookSvc domain.IBookService
}

func NewBookController(bookSvc domain.IBookService) BookController {
	return BookController{
		bookSvc: bookSvc,
	}
}

// CreateBook implements IBookController.
func (bs *BookController) CreateBook(e echo.Context) error {
	reqBook := &types.BookRequest{}
	if err := e.Bind(reqBook); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	if err := reqBook.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	book := &models.BookDetail{
		BookName:    reqBook.BookName,
		AuthorId:    reqBook.AuthorId,
		Publication: reqBook.Publication,
	}

	if err := bs.bookSvc.CreateBook(book); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusCreated, "BookDetail was created successfully")
}

func (bs *BookController) GetAllBooks(e echo.Context) error {
	books, err := bs.bookSvc.GetAllBooks()
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, books)
}

// GetBook implements IBookController.
func (bs *BookController) GetBook(e echo.Context) error {
	tempBookID := e.Param("bookID")
	bookID, err := strconv.ParseInt(tempBookID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Enter a valid book ID")
	}
	//fmt.Println(bookID)
	book, err := bs.bookSvc.GetBook(uint(bookID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, book)
}

// UpdateBook implements IBookController.
func (bs *BookController) UpdateBook(e echo.Context) error {
	reqBook := &types.BookRequest{}
	if err := e.Bind(reqBook); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	tempBookID := e.Param("bookID")
	bookID, err := strconv.ParseInt(tempBookID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Enter a valid book ID")
	}
	updatedBook := &models.BookDetail{
		BookId:      uint(bookID),
		BookName:    reqBook.BookName,
		AuthorId:    reqBook.AuthorId,
		Publication: reqBook.Publication,
	}
	if err := bs.bookSvc.UpdateBook(updatedBook); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusCreated, "BookDetail was updated successfully")
}

// DeleteBook implements IBookController.
func (controller *BookController) DeleteBook(e echo.Context) error {
	tempBookID := e.Param("bookID")
	bookID, err := strconv.ParseInt(tempBookID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}
	_, err = controller.bookSvc.GetBook(uint(bookID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	if err := controller.bookSvc.DeleteBook(uint(bookID)); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, "BookDetail is deleted successfully")
}
