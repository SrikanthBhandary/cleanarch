package repository

import (
	"github.com/srikanthbhandary/cleanarch/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (u *UserRepository) Save(user *models.User) (*models.User, error) {
	return user, u.DB.Create(user).Error
}
func (u *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User
	err := u.DB.Find(&users).Error
	return users, err
}
func (u *UserRepository) Delete(user *models.User) error {
	return u.DB.Delete(&user).Error
}
func (u *UserRepository) Migrate() error {
	return u.DB.AutoMigrate(&models.User{})
}

func NewUserRepository(db *gorm.DB) models.UserRepo {
	return &UserRepository{DB: db}
}
