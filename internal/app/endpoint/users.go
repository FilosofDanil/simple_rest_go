package endpoint

import (
	"awesomeProject/internal/app/models"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type Service interface {
	FindAllUsers(c echo.Context) (*[]models.TestUser, error)

	FindUser(c echo.Context) (*models.TestUser, error)

	CreateUser(c echo.Context) (*models.TestUser, error)

	UpdateUser(c echo.Context) error

	DeleteUser(c echo.Context) error
}

type Endpoint struct {
	s Service
}

func (e *Endpoint) GetAll(ctx echo.Context) error {
	users, err := e.s.FindAllUsers(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return ctx.JSON(http.StatusOK, users)
}

func (e *Endpoint) GetById(ctx echo.Context) error {
	user, err := e.s.FindUser(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, user)
}

func (e *Endpoint) Create(ctx echo.Context) error {
	user, err := e.s.CreateUser(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusCreated, user)
}

func (e *Endpoint) Update(ctx echo.Context) error {
	err := e.s.UpdateUser(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, nil)
}

func (e *Endpoint) Delete(ctx echo.Context) error {
	err := e.s.DeleteUser(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, nil)
}

func New(s Service) (e *Endpoint) {
	return &Endpoint{s: s}
}
