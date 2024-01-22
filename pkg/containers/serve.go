package containers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"noob-server/pkg/config"
	"noob-server/pkg/connection"
	"noob-server/pkg/controllers"
	"noob-server/pkg/repositories"
	"noob-server/pkg/routes"
	"noob-server/pkg/services"
)

func Serve(e *echo.Echo) {
	// config initialization
	config.SetConfig()

	// database connection
	connection.Connect()
	db := connection.GetDB()

	// repository initialization
	bookRepo := repositories.BookDBInstance(db)
	authorRepo := repositories.AuthorDBInstance(db)

	//service initialization
	bookService := services.BookServiceInstance(bookRepo)
	authorService := services.AuthorServiceInstance(authorRepo)

	//controller initialization
	bookCtr := controllers.NewBookController(bookService)
	authorCtr := controllers.NewAuthorController(authorService, bookService)

	//route initialization
	b := routes.BookRoutes(e, bookCtr)
	authorRoutes := routes.AuthorRoutes(e, authorCtr)

	b.InitBookRoute()
	authorRoutes.InitAuthorRoutes()

	// starting server
	log.Fatal(e.Start(fmt.Sprintf(":%s", config.LocalConfig.Port)))
}
