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
	userRepo := repositories.UserDBInstance(db)

	//service initialization
	bookService := services.BookServiceInstance(bookRepo)
	authorService := services.AuthorServiceInstance(authorRepo, bookRepo)
	authService := services.AuthServiceInstance(userRepo)

	//controller initialization
	bookCtr := controllers.NewBookController(bookService)
	authorCtr := controllers.NewAuthorController(authorService)
	authCtr := controllers.NewAuthController(authService)

	//route initialization
	bookRoutes := routes.BookRoutes(e, bookCtr)
	authorRoutes := routes.AuthorRoutes(e, authorCtr)
	authRoutes := routes.NewAuthRoutes(e, authCtr)

	bookRoutes.InitBookRoute()
	authorRoutes.InitAuthorRoutes()
	authRoutes.InitAuthRoutes()

	// starting server
	log.Fatal(e.Start(fmt.Sprintf(":%s", config.LocalConfig.Port)))
}
