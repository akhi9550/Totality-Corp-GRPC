package service

import (
	"context"
	"grpc-user-service/pkg/pb"
	interfaces "grpc-user-service/pkg/usecase/interface"
	"grpc-user-service/pkg/utils/models"
)

type UserSever struct {
	userUseCase interfaces.UserUseCase
	pb.UnimplementedUserServiceServer
}

func NewAuthServer(useCaseUser interfaces.UserUseCase) pb.UserServiceServer {
	return &UserSever{
		userUseCase: useCaseUser,
	}
}

func (s *UserSever) GetUserByID(ctx context.Context, req *pb.UserIDRequest) (*pb.UserResponse, error) {
	results, err := s.userUseCase.GetUserByID(req.Id)
	if err != nil {
		return &pb.UserResponse{}, err
	}
	return &pb.UserResponse{User: &pb.User{
		Id:      results.ID,
		Fname:   results.Fname,
		City:    results.City,
		Phone:   results.Phone,
		Height:  results.Height,
		Married: results.Married,
	}}, nil
}

func (s *UserSever) GetUsersByIDs(ctx context.Context, req *pb.UserIDsRequest) (*pb.UsersResponse, error) {
	users, err := s.userUseCase.GetUsersByIDs(req.Ids)
	if err != nil {
		return &pb.UsersResponse{}, err
	}
	var result []*pb.User
	for _, user := range users {
		result = append(result, &pb.User{
			Id:      user.ID,
			Fname:   user.Fname,
			City:    user.City,
			Phone:   user.Phone,
			Height:  user.Height,
			Married: user.Married,
		})
	}
	return &pb.UsersResponse{
		Users: result,
	}, nil
}

func (s *UserSever) SearchUsers(ctx context.Context, req *pb.SearchRequest) (*pb.UsersResponse, error) {
	search := models.SearchUser{
		City:    req.City,
		Phone:   req.Phone,
		Married: req.Married,
	}
	users, err := s.userUseCase.SearchUsers(search)
	if err != nil {
		return &pb.UsersResponse{}, err
	}
	var result []*pb.User
	for _, user := range users {
		result = append(result, &pb.User{
			Id:      user.ID,
			Fname:   user.Fname,
			City:    user.City,
			Phone:   user.Phone,
			Height:  user.Height,
			Married: user.Married,
		})
	}
	return &pb.UsersResponse{
		Users: result,
	}, nil
}

func (s *UserSever) AddUser(ctx context.Context, req *pb.AddUserRequest) (*pb.AddUserResponse, error) {
	newUser := models.User{
		Fname:   req.User.Fname,
		City:    req.User.City,
		Phone:   req.User.Phone,
		Height:  req.User.Height,
		Married: req.User.Married,
	}
	err := s.userUseCase.AddUser(newUser)
	if err != nil {
		return &pb.AddUserResponse{}, err
	}
	return &pb.AddUserResponse{}, nil
}
