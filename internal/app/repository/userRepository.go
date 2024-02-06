package repository

import (
	"awesomeProject/internal/app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// UserRepository represents the repository for User model.
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new UserRepository instance.
func NewUserRepository(dsn string) (*UserRepository, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	//Auto-migrate the User model to create the table.

	if err = db.Migrator().AutoMigrate(&models.TestUser{}); err != nil {
		return nil, err
	}

	return &UserRepository{db}, nil
}

// CreateUser adds a new user to the database.
func (ur *UserRepository) CreateUser(user *models.TestUser) error {
	return ur.db.Create(user).Error
}

// GetUserByID retrieves a user by ID from the database.
func (ur *UserRepository) GetUserByID(userID uint) (*models.TestUser, error) {
	var user models.TestUser
	err := ur.db.First(&user, userID).Error
	return &user, err
}

// GetAllUsers retrieves all the rows.
func (ur *UserRepository) GetAllUsers() (*[]models.TestUser, error) {
	var users []models.TestUser
	err := ur.db.Find(&users).Error
	return &users, err
}

// UpdateUser updates the information of an existing user in the database.
func (ur *UserRepository) UpdateUser(user *models.TestUser) error {
	return ur.db.Save(user).Error
}

// DeleteUserByID deletes a user by ID from the database.
func (ur *UserRepository) DeleteUserByID(userID uint) error {
	return ur.db.Delete(&models.TestUser{}, userID).Error
}
