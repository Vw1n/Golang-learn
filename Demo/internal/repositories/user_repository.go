package repositories

// 创建用户存储库

import (
	"Demo/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := ur.db.Find(&users)
	return users, result.Error
}

func (ur *UserRepository) GetUserByID(id uint) (models.User, error) {
	var user models.User
	result := ur.db.First(&user, id)
	return user, result.Error
}

func (ur *UserRepository) CreateUser(user *models.User) error {
	result := ur.db.Create(user)
	return result.Error
}

func (ur *UserRepository) UpdateUser(user *models.User) error {
	result := ur.db.Save(user)
	return result.Error
}

func (ur *UserRepository) DeleteUser(id uint) error {
	result := ur.db.Delete(&models.User{}, id)
	return result.Error
}
