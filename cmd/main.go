package main

import (
	"awesomeProject/internal/pkg/app"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func main() {
	//init server
	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}
	if err := a.Run(); err != nil {
		log.Fatal(err)
	}
}

func Handler(c echo.Context) error {
	return c.String(http.StatusOK, "Hi!")
}
