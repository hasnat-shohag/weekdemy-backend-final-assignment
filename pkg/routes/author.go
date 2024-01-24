package routes

import (
	"github.com/labstack/echo/v4"
	"noob-server/pkg/controllers"
)

type authorRoutes struct {
	echo             *echo.Echo
	authorController controllers.AuthorController
}

func AuthorRoutes(echo *echo.Echo, authorController controllers.AuthorController) *authorRoutes {
	return &authorRoutes{
		echo:             echo,
		authorController: authorController,
	}
}

func (authorRoutes *authorRoutes) InitAuthorRoutes() {
	e := authorRoutes.echo
	authorRoutes.initAuthorRoutes(e)
}

func (authorRoutes *authorRoutes) initAuthorRoutes(e *echo.Echo) {
	author := e.Group("/bookstore")
	author.POST("/authors", authorRoutes.authorController.CreateAuthor)
	author.GET("/authors", authorRoutes.authorController.GetAllAuthors)
	author.GET("/authors/:authorID", authorRoutes.authorController.GetAuthor)
	author.PUT("/authors/:authorID", authorRoutes.authorController.UpdateAuthor)
	// author.GET("/authors/:authorID/books", authorRoutes.authorController.GetBooksByAuthorID)
	author.DELETE("/authors/:authorID", authorRoutes.authorController.DeleteAuthor)
}
