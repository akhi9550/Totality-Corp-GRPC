package client

import (
	"context"
	"fmt"
	interfaces "grpc-user-api-gateway/pkg/client/interface"
	"grpc-user-api-gateway/pkg/config"
	"grpc-user-api-gateway/pkg/pb"
	"grpc-user-api-gateway/pkg/utils/models"

	"google.golang.org/grpc"
)

type UserClient struct {
	Client pb.UserServiceClient
}

func NewUserClient(cfg config.Config) interfaces.UserClient {
	grpcConnection, err := grpc.Dial(cfg.UserSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect", err)
	}

	grpcClient := pb.NewUserServiceClient(grpcConnection)

	return &UserClient{
		Client: grpcClient,
	}
}

func (u *UserClient) GetUserByID(id int64) (models.Users, error) {
	results, err := u.Client.GetUserByID(context.Background(), &pb.UserIDRequest{
		Id: id,
	})

	if err != nil {
		return models.Users{}, err
	}

	return models.Users{
		ID:      results.User.Id,
		FName:   results.User.Fname,
		City:    results.User.City,
		Phone:   results.User.Phone,
		Height:  results.User.Height,
		Married: results.User.Married,
	}, nil
}

func (u *UserClient) GetUsersByIDs(ids []int64) ([]models.Users, error) {
	users, err := u.Client.GetUsersByIDs(context.Background(), &pb.UserIDsRequest{
		Ids: ids,
	})
	if err != nil {
		return []models.Users{}, err
	}
	var results []models.Users
	for _, v := range users.Users {
		result := models.Users{
			ID:      v.Id,
			FName:   v.Fname,
			City:    v.City,
			Phone:   v.Phone,
			Height:  v.Height,
			Married: v.Married,
		}
		results = append(results, result)
	}
	return results, nil
}

func (u *UserClient) SearchUsers(search models.SearchUser) ([]models.Users, error) {
	users, err := u.Client.SearchUsers(context.Background(), &pb.SearchRequest{
		City:    search.City,
		Phone:   search.Phone,
		Married: search.Married,
	})

	if err != nil {
		return []models.Users{}, err
	}
	var results []models.Users
	for _, v := range users.Users {
		result := models.Users{
			ID:      v.Id,
			FName:   v.Fname,
			City:    v.City,
			Phone:   v.Phone,
			Height:  v.Height,
			Married: v.Married,
		}
		results = append(results, result)
	}
	return results, nil
}

func (u *UserClient) AddUser(user models.User) error {
	users := &pb.Users{Fname: user.FName, City: user.City, Phone: user.Phone, Height: user.Height, Married: user.Married}
	_, err := u.Client.AddUser(context.Background(), &pb.AddUserRequest{
		User: users,
	})
	if err != nil {
		return err
	}
	return nil
}
