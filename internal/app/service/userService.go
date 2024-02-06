package service

import (
	"awesomeProject/internal/app/models"
	"github.com/labstack/echo/v4"
	"strconv"
)

type Repository interface {
	CreateUser(user *models.TestUser) error

	GetUserByID(userID uint) (*models.TestUser, error)

	GetAllUsers() (*[]models.TestUser, error)

	UpdateUser(user *models.TestUser) error

	DeleteUserByID(userID uint) error
}

type UserService struct {
	userRepo Repository
}

func NewUserService(repo Repository) *UserService {
	return &UserService{userRepo: repo}
}

func (service *UserService) FindAllUsers(c echo.Context) (*[]models.TestUser, error) {
	users, err := service.userRepo.GetAllUsers()
	return users, err
}

func (service *UserService) FindUser(c echo.Context) (*models.TestUser, error) {
	id, err := strconv.Atoi(c.Param("id"))
	var user *models.TestUser
	user, err = service.userRepo.GetUserByID(uint(id))
	return user, err
}

func (service *UserService) CreateUser(c echo.Context) (*models.TestUser, error) {
	//TODO implement me
	user := new(models.TestUser)
	err := c.Bind(user)
	if err != nil {
		return nil, err
	}
	err = service.userRepo.CreateUser(user)
	return user, err
}

func (service *UserService) UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	var userToUpdate *models.TestUser
	if err != nil {
		return err
	}
	userToUpdate, err = service.userRepo.GetUserByID(uint(id))
	user := new(models.TestUser)
	err = c.Bind(user)
	if err != nil {
		return err
	}
	modifyUser(user, userToUpdate)
	return service.userRepo.UpdateUser(userToUpdate)
}

func (service *UserService) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	return service.userRepo.DeleteUserByID(uint(id))
}

func modifyUser(user *models.TestUser, userToUpdate *models.TestUser) {
	userToUpdate.TgName = user.TgName
	userToUpdate.Username = user.Username
	userToUpdate.TgSurname = user.TgSurname
	userToUpdate.ChatID = user.ChatID

}
