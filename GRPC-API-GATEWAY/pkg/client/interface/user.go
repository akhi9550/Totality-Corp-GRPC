package interfaces

import "grpc-user-api-gateway/pkg/utils/models"

type UserClient interface {
	GetUserByID(id int64) (models.Users, error)
	GetUsersByIDs(ids []int64) ([]models.Users, error)
	SearchUsers(search models.SearchUser) ([]models.Users, error)
	AddUser(user models.User) error
}
