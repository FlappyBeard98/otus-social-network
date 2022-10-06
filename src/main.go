package main

import (
	"github.com/labstack/echo/v4"
	"social-network/api"
)

func main() {

	e := echo.New()
	api.SetupRoutes(e)
	e.Logger.Fatal(e.Start(":1323"))
}


