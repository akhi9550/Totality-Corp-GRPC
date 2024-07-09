package service_test

import (
	"context"
	"errors"
	"testing"

	"grpc-user-service/pkg/api/service"
	"grpc-user-service/pkg/pb"
	mock_usecase "grpc-user-service/pkg/usecase/mock"
	"grpc-user-service/pkg/utils/models"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserServer_GetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mock_usecase.NewMockUserUseCase(ctrl)
	userServer := service.NewAuthServer(mockUseCase)

	testCases := []struct {
		name          string
		userID        int64
		buildStubs    func()
		checkResponse func(t *testing.T, resp *pb.UserResponse, err error)
	}{
		{
			name:   "Success",
			userID: 1,
			buildStubs: func() {
				mockUseCase.EXPECT().
					GetUserByID(int64(1)).
					Times(1).
					Return(models.Users{ID: 1, Fname: "John", City: "New York", Phone: "1234567890", Height: 180, Married: false}, nil)
			},
			checkResponse: func(t *testing.T, resp *pb.UserResponse, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, "John", resp.User.Fname)
			},
		},
		{
			name:   "User Not Found",
			userID: 2,
			buildStubs: func() {
				mockUseCase.EXPECT().
					GetUserByID(int64(2)).
					Times(1).
					Return(models.Users{}, errors.New("user not found"))
			},
			checkResponse: func(t *testing.T, resp *pb.UserResponse, err error) {
				assert.Error(t, err)
				assert.NotNil(t, resp)
				assert.Nil(t, resp.User)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.buildStubs()
			req := &pb.UserIDRequest{Id: tc.userID}
			resp, err := userServer.GetUserByID(context.Background(), req)
			tc.checkResponse(t, resp, err)
		})
	}
}



func TestUserServer_GetUsersByIDs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mock_usecase.NewMockUserUseCase(ctrl)
	userServer := service.NewAuthServer(mockUseCase)

	testCases := []struct {
		name          string
		userIDs       []int64
		buildStubs    func()
		checkResponse func(t *testing.T, resp *pb.UsersResponse, err error)
	}{
		{
			name:    "Success",
			userIDs: []int64{1, 2},
			buildStubs: func() {
				mockUseCase.EXPECT().
					GetUsersByIDs([]int64{1, 2}).
					Times(1).
					Return([]models.Users{
						{ID: 1, Fname: "John", City: "New York", Phone: "1234567890", Height: 180, Married: false},
						{ID: 2, Fname: "Doe", City: "Los Angeles", Phone: "0987654321", Height: 170, Married: true},
					}, nil)
			},
			checkResponse: func(t *testing.T, resp *pb.UsersResponse, err error) {
				assert.NoError(t, err)
				assert.Len(t, resp.Users, 2)
				assert.Equal(t, "John", resp.Users[0].Fname)
				assert.Equal(t, "Doe", resp.Users[1].Fname)
			},
		},
		{
			name:    "Error",
			userIDs: []int64{1, 2},
			buildStubs: func() {
				mockUseCase.EXPECT().
					GetUsersByIDs([]int64{1, 2}).
					Times(1).
					Return(nil, errors.New("internal error"))
			},
			checkResponse: func(t *testing.T, resp *pb.UsersResponse, err error) {
				assert.Error(t, err)
				assert.Nil(t, resp.Users)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.buildStubs()
			req := &pb.UserIDsRequest{Ids: tc.userIDs}
			resp, err := userServer.GetUsersByIDs(context.Background(), req)
			tc.checkResponse(t, resp, err)
		})
	}
}

func TestUserServer_SearchUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mock_usecase.NewMockUserUseCase(ctrl)
	userServer := service.NewAuthServer(mockUseCase)

	testCases := []struct {
		name          string
		searchRequest *pb.SearchRequest
		buildStubs    func()
		checkResponse func(t *testing.T, resp *pb.UsersResponse, err error)
	}{
		{
			name: "Success",
			searchRequest: &pb.SearchRequest{
				City:    "New York",
				Phone:   "1234567890",
				Married: false,
			},
			buildStubs: func() {
				mockUseCase.EXPECT().
					SearchUsers(models.SearchUser{City: "New York", Phone: "1234567890", Married: false}).
					Times(1).
					Return([]models.Users{
						{ID: 1, Fname: "John", City: "New York", Phone: "1234567890", Height: 180, Married: false},
					}, nil)
			},
			checkResponse: func(t *testing.T, resp *pb.UsersResponse, err error) {
				assert.NoError(t, err)
				assert.Len(t, resp.Users, 1)
				assert.Equal(t, "John", resp.Users[0].Fname)
			},
		},
		{
			name: "Error",
			searchRequest: &pb.SearchRequest{
				City:    "New York",
				Phone:   "1234567890",
				Married: false,
			},
			buildStubs: func() {
				mockUseCase.EXPECT().
					SearchUsers(models.SearchUser{City: "New York", Phone: "1234567890", Married: false}).
					Times(1).
					Return(nil, errors.New("search error"))
			},
			checkResponse: func(t *testing.T, resp *pb.UsersResponse, err error) {
				assert.Error(t, err)
				assert.Nil(t, resp.Users)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.buildStubs()
			resp, err := userServer.SearchUsers(context.Background(), tc.searchRequest)
			tc.checkResponse(t, resp, err)
		})
	}
}

func TestUserServer_AddUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUseCase := mock_usecase.NewMockUserUseCase(ctrl)
	userServer := service.NewAuthServer(mockUseCase)

	testCases := []struct {
		name           string
		addUserRequest *pb.AddUserRequest
		buildStubs     func()
		checkResponse  func(t *testing.T, resp *pb.AddUserResponse, err error)
	}{
		{
			name: "Success",
			addUserRequest: &pb.AddUserRequest{
				User: &pb.Users{
					Fname:   "John",
					City:    "New York",
					Phone:   "1234567890",
					Height:  180,
					Married: false,
				},
			},
			buildStubs: func() {
				mockUseCase.EXPECT().
					AddUser(models.User{Fname: "John", City: "New York", Phone: "1234567890", Height: 180, Married: false}).
					Times(1).
					Return(nil)
			},
			checkResponse: func(t *testing.T, resp *pb.AddUserResponse, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
			},
		},
		{
			name: "Error",
			addUserRequest: &pb.AddUserRequest{
				User: &pb.Users{
					Fname:   "John",
					City:    "New York",
					Phone:   "1234567890",
					Height:  180,
					Married: false,
				},
			},
			buildStubs: func() {
				mockUseCase.EXPECT().
					AddUser(models.User{Fname: "John", City: "New York", Phone: "1234567890", Height: 180, Married: false}).
					Times(1).
					Return(errors.New("add user error"))
			},
			checkResponse: func(t *testing.T, resp *pb.AddUserResponse, err error) {
				assert.Error(t, err)
				assert.NotNil(t, resp)
				assert.Empty(t, resp)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.buildStubs()
			resp, err := userServer.AddUser(context.Background(), tc.addUserRequest)
			tc.checkResponse(t, resp, err)
		})
	}
}