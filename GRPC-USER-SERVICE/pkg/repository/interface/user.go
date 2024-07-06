package interfaces

import "grpc-user-service/pkg/utils/models"

type UserRepository interface {
	AddUser(user models.User) error
	CheckUserExistsByPhone(phone string) bool
	CheckUserAvailabilityWithUserID(Id int64) bool
	CheckUserAvailabilityWithUserIDs(Id []int64) bool
	GetUserByID(Id int64) (models.Users, error)
	GetUsersByIDs(Ids []int64) ([]models.Users, error)
	SearchCity(city string) ([]models.Users, error)
	SearchPhone(phone string) ([]models.Users, error)
	SearchMarried(married bool) ([]models.Users, error)
}
