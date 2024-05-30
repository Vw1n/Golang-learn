package repositories

// 创建用户存储库

import "Demo/internal/models"

type UserRepository struct {
	users []models.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) GetAllUsers() []models.User {
	// 实现获取所有用户的逻辑
	return ur.users
}
