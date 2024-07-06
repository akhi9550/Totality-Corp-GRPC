package usecase

import (
	"errors"
	"fmt"
	interfaces "grpc-user-service/pkg/repository/interface"
	server "grpc-user-service/pkg/usecase/interface"
	"grpc-user-service/pkg/utils/models"
)

type userUseCase struct {
	userRepository interfaces.UserRepository
}

func NewUserUseCase(repository interfaces.UserRepository) server.UserUseCase {
	return &userUseCase{
		userRepository: repository,
	}
}

func (u *userUseCase) GetUserByID(id int64) (models.Users, error) {
	userExist := u.userRepository.CheckUserAvailabilityWithUserID(id)
	if !userExist {
		return models.Users{}, errors.New("user doesn't exist")
	}
	result, err := u.userRepository.GetUserByID(id)
	if err != nil {
		return models.Users{}, err
	}
	return result, nil
}

func (u *userUseCase) GetUsersByIDs(ids []int64) ([]models.Users, error) {
	userExist := u.userRepository.CheckUserAvailabilityWithUserIDs(ids)
	if !userExist {
		return []models.Users{}, errors.New("user doesn't exist")
	}
	result, err := u.userRepository.GetUsersByIDs(ids)
	if err != nil {
		return []models.Users{}, err
	}
	return result, nil
}

func (u *userUseCase) SearchUsers(search models.SearchUser) ([]models.Users, error) {
	fmt.Println("serarch", search.Married)
	var result []models.Users
	foundUsers := make(map[int64]bool)

	if search.City != "" {
		res, err := u.userRepository.SearchCity(search.City)
		if err != nil {
			return nil, err
		}
		for _, user := range res {
			if !foundUsers[user.ID] {
				result = append(result, user)
				foundUsers[user.ID] = true
			}
		}
	}

	if search.Phone != "" {
		res, err := u.userRepository.SearchPhone(search.Phone)
		if err != nil {
			return nil, err
		}
		for _, user := range res {
			if !foundUsers[user.ID] {
				result = append(result, user)
				foundUsers[user.ID] = true
			}
		}
	}

	if search.Married || !search.Married {
		res, err := u.userRepository.SearchMarried(search.Married)
		if err != nil {
			return nil, err
		}
		for _, user := range res {
			if !foundUsers[user.ID] {
				result = append(result, user)
				foundUsers[user.ID] = true
			}
		}
	}
	return result, nil
}

func (u *userUseCase) AddUser(user models.User) error {
	phone := u.userRepository.CheckUserExistsByPhone(user.Phone)
	if phone {
		return errors.New("user with this phone is already exists")
	}
	err := u.userRepository.AddUser(user)
	if err != nil {
		return err
	}
	return nil
}
