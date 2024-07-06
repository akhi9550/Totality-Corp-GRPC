package handler

import (
	"fmt"
	interfaces "grpc-user-api-gateway/pkg/client/interface"
	"grpc-user-api-gateway/pkg/utils/helper"
	"grpc-user-api-gateway/pkg/utils/models"
	"grpc-user-api-gateway/pkg/utils/response"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	GRPC_Client interfaces.UserClient
}

func NewUserHandler(userClient interfaces.UserClient) *UserHandler {
	return &UserHandler{
		GRPC_Client: userClient,
	}
}

func (u *UserHandler) AddUser(c *gin.Context) {
	var AddUser models.User
	if err := c.ShouldBindJSON(&AddUser); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
	}

	pattern := `^\d{10}$`
	regex := regexp.MustCompile(pattern)
	value := regex.MatchString(AddUser.Phone)
	if !value {
		fmt.Printf("%s phone number is not valid", AddUser.Phone)
		return
	}

	err := validator.New().Struct(AddUser)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Constraints not statisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err = u.GRPC_Client.AddUser(AddUser)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusCreated, "User added successfully", nil, nil)
	c.JSON(http.StatusCreated, success)
}

func (au *UserHandler) GetUserByID(c *gin.Context) {
	userID := c.Query("user_id")
	UserID, err := strconv.Atoi(userID)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "UserID not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	user, err := au.GRPC_Client.GetUserByID(int64(UserID))
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusCreated, "Successfully get Userdetails", user, nil)
	c.JSON(http.StatusCreated, success)
}

func (au *UserHandler) GetUsersByIDs(c *gin.Context) {
	user := c.PostFormArray("user_ids")
	users, err := helper.ConvertStringToArray(user)
	if err != nil {
		return
	}
	Users, err := au.GRPC_Client.GetUsersByIDs(users)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusCreated, "Successfully get Userdetails", Users, nil)
	c.JSON(http.StatusCreated, success)
}

func (au *UserHandler) SearchUsers(c *gin.Context) {
	var Search models.SearchUser
	if err := c.ShouldBindJSON(&Search); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
	}

	user, err := au.GRPC_Client.SearchUsers(Search)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusCreated, "Successfully get Userdetails", user, nil)
	c.JSON(http.StatusCreated, success)
}
