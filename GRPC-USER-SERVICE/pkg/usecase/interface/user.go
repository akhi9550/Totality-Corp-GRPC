package interfaces

import "grpc-user-service/pkg/utils/models"

type UserUseCase interface {
	GetUserByID(id int64) (models.Users, error)
	GetUsersByIDs(ids []int64) ([]models.Users, error)
	SearchUsers(search models.SearchUser) ([]models.Users, error)
	AddUser(user models.User) error
}
