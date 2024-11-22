package interfaces

import "demo/models"

type IUser interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUserBy(id string) (*models.User, error)
	GetUsers(limit, offset int) ([]models.User, error)
	DeleteUserBy(id string) (int64, error)
}
