package main

import (
	"github.com/labstack/echo/v4"
	"noob-server/pkg/containers"
)

func main() {
	// New Echo Instance
	e := echo.New()
	containers.Serve(e)
}
