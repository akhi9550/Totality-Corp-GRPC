package repository

import (
	interfaces "grpc-user-service/pkg/repository/interface"
	"grpc-user-service/pkg/utils/models"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userRepository{
		DB: DB,
	}
}

func (u *userRepository) GetUserByID(Id int64) (models.Users, error) {
	var user models.Users
	err := u.DB.Raw(`SELECT id, fname, city, phone, height, married FROM users WHERE id=$1`, Id).Scan(&user).Error
	if err != nil {
		return models.Users{}, err
	}
	return user, nil
}
func (u *userRepository) GetUsersByIDs(Ids []int64) ([]models.Users, error) {
	var users []models.Users
	for _, id := range Ids {
		var user models.Users
		err := u.DB.Raw(`SELECT id, fname, city, phone, height, married FROM users WHERE id=$1`, id).Scan(&user).Error
		if err == nil {
			users = append(users, user)
		}
	}
	return users, nil
}

func (u *userRepository) SearchCity(city string) ([]models.Users, error) {
	var users []models.Users
	err := u.DB.Raw(`SELECT id, fname, city, phone, height, married FROM users WHERE city ILIKE '%' || $1 || '%'`, city).Scan(&users).Error
	if err != nil {
		return []models.Users{}, err
	}
	return users, nil
}
func (u *userRepository) SearchPhone(phone string) ([]models.Users, error) {
	var users []models.Users
	err := u.DB.Raw(`SELECT id, fname, city, phone, height, married FROM users WHERE phone=$1`, phone).Scan(&users).Error
	if err != nil {
		return []models.Users{}, err
	}
	return users, nil
}

func (u *userRepository) SearchMarried(married bool) ([]models.Users, error) {
	var users []models.Users
	err := u.DB.Raw(`SELECT id, fname, city, phone, height, married FROM users WHERE married=$1`, married).Scan(&users).Error
	if err != nil {
		return []models.Users{}, err
	}
	return users, nil
}

func (u *userRepository) AddUser(user models.User) error {
	err := u.DB.Exec(`INSERT INTO users (fname, city, phone, height, married) VALUES ($1, $2, $3, $4, $5)`,
		user.Fname, user.City, user.Phone, user.Height, user.Married).Error
	if err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) CheckUserExistsByPhone(phone string) bool {
	var count int
	if err := ur.DB.Raw("SELECT count(*) FROM users WHERE phone = ?", phone).Scan(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (ur *userRepository) CheckUserAvailabilityWithUserID(Id int64) bool {
	var count int
	if err := ur.DB.Raw("SELECT count(*) FROM users WHERE id = ?", Id).Scan(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (u *userRepository) CheckUserAvailabilityWithUserIDs(ids []int64) bool {
	for _, id := range ids {
		var count int
		if err := u.DB.Raw("SELECT count(*) FROM users WHERE id = ?", id).Scan(&count).Error; err != nil {
			return false
		}
		if count == 0 {
			return false
		}
	}
	return true
}
