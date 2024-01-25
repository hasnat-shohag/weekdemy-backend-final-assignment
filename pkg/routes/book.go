package routes

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"noob-server/pkg/controllers"
	"noob-server/pkg/middlewares"
)

type bookRoutes struct {
	echo    *echo.Echo
	bookCtr controllers.BookController
}

func BookRoutes(echo *echo.Echo, bookCtr controllers.BookController) *bookRoutes {
	return &bookRoutes{
		echo:    echo,
		bookCtr: bookCtr,
	}
}

func (bc *bookRoutes) InitBookRoute() {
	e := bc.echo
	bc.initBookRoutes(e)
}

func (bc *bookRoutes) initBookRoutes(e *echo.Echo) {
	//grouping route endpoints
	book := e.Group("/bookstore")

	book.GET("/ping", Pong)

	//initializing http methods - routing endpoints and their handlers
	book.GET("/books", bc.bookCtr.GetAllBooks)
	book.GET("/books/:bookID", bc.bookCtr.GetBook)

	book.Use(middlewares.ValidateToken)

	book.POST("/books", bc.bookCtr.CreateBook)
	book.PUT("/books/:bookID", bc.bookCtr.UpdateBook)
	book.DELETE("/books/:bookID", bc.bookCtr.DeleteBook)
}

func Pong(ctx echo.Context) error {
	fmt.Println("Pong")
	return ctx.JSON(http.StatusOK, "Pong")
}
