package routes

import (
	"github.com/labstack/echo/v4"
	"noob-server/pkg/controllers"
	"noob-server/pkg/middlewares"
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
	author.GET("/authors", authorRoutes.authorController.GetAllAuthors)
	author.GET("/authors/:authorID", authorRoutes.authorController.GetAuthor)

	author.Use(middlewares.ValidateToken)

	author.POST("/authors", authorRoutes.authorController.CreateAuthor)
	author.PUT("/authors/:authorID", authorRoutes.authorController.UpdateAuthor)
	author.DELETE("/authors/:authorID", authorRoutes.authorController.DeleteAuthor)
}
