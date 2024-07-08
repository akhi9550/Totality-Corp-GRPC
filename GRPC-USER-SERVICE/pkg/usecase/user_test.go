package usecase

import (
	"errors"
	mock_repository "grpc-user-service/pkg/repository/mock"
	"grpc-user-service/pkg/utils/models"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_GetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockUserRepository(ctrl)
	useCase := NewUserUseCase(mockRepo)

	testID := int64(1)

	mockRepo.EXPECT().CheckUserAvailabilityWithUserID(testID).Return(true)
	mockRepo.EXPECT().GetUserByID(testID).Return(models.Users{ID: testID, Fname: "Test User"}, nil)

	result, err := useCase.GetUserByID(testID)

	assert.NoError(t, err)
	assert.Equal(t, "Test User", result.Fname)

	mockRepo.EXPECT().CheckUserAvailabilityWithUserID(testID).Return(false)

	_, err = useCase.GetUserByID(testID)
	assert.Error(t, err)
	assert.Equal(t, "user doesn't exist", err.Error())

	mockRepo.EXPECT().CheckUserAvailabilityWithUserID(testID).Return(true)
	mockRepo.EXPECT().GetUserByID(testID).Return(models.Users{}, errors.New("repository error"))

	_, err = useCase.GetUserByID(testID)
	assert.Error(t, err)
	assert.Equal(t, "repository error", err.Error())
}

func Test_GetUsersByIDs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockUserRepository(ctrl)
	useCase := NewUserUseCase(mockRepo)

	testIDs := []int64{1, 2, 3}

	mockRepo.EXPECT().CheckUserAvailabilityWithUserIDs(testIDs).Return(true)
	mockRepo.EXPECT().GetUsersByIDs(testIDs).Return([]models.Users{
		{ID: 1, Fname: "User1", City: "Kannur", Phone: "0123456789", Height: 175.6, Married: true},
		{ID: 2, Fname: "User2", City: "Trissur", Phone: "0123456729", Height: 165.6, Married: false},
		{ID: 3, Fname: "User3", City: "Kozhikode", Phone: "0123456589", Height: 155.6, Married: true},
	}, nil)

	result, err := useCase.GetUsersByIDs(testIDs)

	assert.NoError(t, err)
	assert.Len(t, result, 3)
	assert.Equal(t, "User1", result[0].Fname)

	mockRepo.EXPECT().CheckUserAvailabilityWithUserIDs(testIDs).Return(false)

	_, err = useCase.GetUsersByIDs(testIDs)
	assert.Error(t, err)
	assert.Equal(t, "user doesn't exist", err.Error())

	mockRepo.EXPECT().CheckUserAvailabilityWithUserIDs(testIDs).Return(true)
	mockRepo.EXPECT().GetUsersByIDs(testIDs).Return([]models.Users{}, errors.New("repository error"))

	_, err = useCase.GetUsersByIDs(testIDs)
	assert.Error(t, err)
	assert.Equal(t, "repository error", err.Error())
}

func Test_SearchUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockUserRepository(ctrl)
	useCase := NewUserUseCase(mockRepo)

	testSearch := models.SearchUser{
		City:    "TestCity",
		Phone:   "1234567890",
		Married: true,
	}

	mockRepo.EXPECT().SearchCity(testSearch.City).Return([]models.Users{
		{ID: 1, Fname: "User1", City: testSearch.City},
	}, nil).Times(1)

	mockRepo.EXPECT().SearchPhone(testSearch.Phone).Return([]models.Users{
		{ID: 1, Fname: "User1", Phone: testSearch.Phone},
	}, nil).Times(1)

	mockRepo.EXPECT().SearchMarried(testSearch.Married).Return([]models.Users{
		{ID: 1, Fname: "User1", Married: testSearch.Married},
	}, nil).Times(1)

	result, err := useCase.SearchUsers(testSearch)

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "User1", result[0].Fname)
}

func Test_AddUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockUserRepository(ctrl)
	useCase := NewUserUseCase(mockRepo)

	testUser := models.User{
		Fname:   "Test User",
		Phone:   "1234567890",
		City:    "Test City",
		Height:  170.5,
		Married: false,
	}

	mockRepo.EXPECT().CheckUserExistsByPhone(testUser.Phone).Return(false)
	mockRepo.EXPECT().AddUser(testUser).Return(nil)

	err := useCase.AddUser(testUser)

	assert.NoError(t, err)

	mockRepo.EXPECT().CheckUserExistsByPhone(testUser.Phone).Return(true)

	err = useCase.AddUser(testUser)
	assert.Error(t, err)
	assert.Equal(t, "user with this phone is already exists", err.Error())

	mockRepo.EXPECT().CheckUserExistsByPhone(testUser.Phone).Return(false)
	mockRepo.EXPECT().AddUser(testUser).Return(errors.New("repository error"))

	err = useCase.AddUser(testUser)
	assert.Error(t, err)
	assert.Equal(t, "repository error", err.Error())
}
