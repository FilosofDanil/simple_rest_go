package app

import (
	c "awesomeProject/configs"
	"awesomeProject/internal/app/endpoint"
	"awesomeProject/internal/app/repository"
	"awesomeProject/internal/app/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.UserService
	r    *repository.UserRepository
	echo *echo.Echo
}

func New() (*App, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	a := &App{}
	repo, err := repository.NewUserRepository(dsn)
	if err != nil {
		log.Fatal(err)
	}
	a.r = repo
	a.s = service.NewUserService(a.r)
	a.e = endpoint.New(a.s)
	a.echo = echo.New()
	a.echo.Use(middleware.Recover())
	a.echo.Use(middleware.Logger())
	a.echo.GET("api/v1/users", a.e.GetAll)
	a.echo.GET("api/v1/users/:id", a.e.GetById)
	a.echo.POST("api/v1/users", a.e.Create)
	a.echo.PUT("api/v1/users/:id", a.e.Update)
	a.echo.DELETE("api/v1/users/:id", a.e.Delete)
	return a, nil
}

func (a *App) Run() error {
	if err := a.echo.Start(c.GetInstance().Server.Port); err != nil {
		log.Fatal(err)
	}
	return nil
}
