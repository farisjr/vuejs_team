package main

import (
	"fmt"
	"vuejs/collabs/config"
	"vuejs/collabs/routes"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	config.InitDb()
	config.InitPort()
	routes.New(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.HTTP_PORT)))
}
