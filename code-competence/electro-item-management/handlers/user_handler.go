package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"electro-item-management/constants"
	"electro-item-management/domains"
	"electro-item-management/services"

	"github.com/labstack/echo/v4"
)

type (
	UserHandler interface{}

	userHandler struct {
		userService services.UserService
	}
)

func NewUserHandler(userServ services.UserService) *userHandler {
	return &userHandler{userService: userServ}
}

func (u *userHandler) GetUsers(c echo.Context) error {
	response := domains.GeneralListResponse{}
	err := domains.ErrorCode{}

	pageString := c.QueryParam("page")
	if pageString != "" {
		response.Page, err.Err = strconv.Atoi(pageString)
		if err.Err != nil {
			response.StatusCode = 400
			response.Message = "invalid page parameter"
			return c.JSON(response.StatusCode, response)
		}
	}
	response.DataShown = constants.LIMIT
	offset := 0
	if response.Page > 1 {
		offset = response.Page * response.DataShown
	}

	users, totalData, err := u.userService.GetUsersService(offset, response.DataShown)
	if err.Err != nil {
		response.StatusCode = err.StatusCode
		response.Message = err.Err.Error()
		return c.JSON(response.StatusCode, response)
	}

	response.StatusCode = 200
	response.Message = "success"
	response.TotalData = totalData
	response.Data = users

	return c.JSON(http.StatusOK, response)
}

func (u *userHandler) GetUserByID(c echo.Context) error {
	response := domains.GeneralResponse{}
	err := domains.ErrorCode{}
	user := domains.User{}

	userIDString := c.Param("id")
	var userID int
	userID, err.Err = strconv.Atoi(userIDString)
	if err.Err != nil {
		response.StatusCode = 400
		response.Message = fmt.Sprintf("invalid id parameter %v", userIDString)
		return c.JSON(response.StatusCode, response)
	}
	user.ID = uint(userID)

	err = u.userService.GetUserService(&user)
	if err.Err != nil {
		response.StatusCode = err.StatusCode
		response.Message = err.Err.Error()
		return c.JSON(response.StatusCode, response)
	}

	response.StatusCode = 200
	response.Message = "success"
	response.Data = user

	return c.JSON(http.StatusOK, response)
}

func (u *userHandler) CreateUser(c echo.Context) error {
	response := domains.GeneralResponse{}
	user := domains.User{}

	if err := c.Bind(&user); err != nil {
		if err != nil {
			response.StatusCode = 400
			response.Message = err.Error()
			return c.JSON(response.StatusCode, response)
		}
	}

	err := u.userService.CreateUserService(&user)
	if err.Err != nil {
		response.StatusCode = err.StatusCode
		response.Message = err.Err.Error()
		return c.JSON(response.StatusCode, response)
	}

	response.StatusCode = 201
	response.Message = "success create user"
	response.Data = user

	return c.JSON(response.StatusCode, response)
}

func (u *userHandler) EditUser(c echo.Context) error {
	response := domains.GeneralResponse{}
	err := domains.ErrorCode{}
	var userID int
	userID, err.Err = strconv.Atoi(c.Param("id"))
	if err.Err != nil {
		response.StatusCode = 400
		response.Message = "id parameter must be valid"
		return c.JSON(response.StatusCode, response)
	}

	modifiedUserData := domains.User{}
	if err := c.Bind(&modifiedUserData); err != nil {
		response.StatusCode = 400
		response.Message = err.Error()
		return c.JSON(response.StatusCode, response)
	}

	user, err := u.userService.EditUserService(userID, &modifiedUserData)
	if err.Err != nil {
		response.StatusCode = err.StatusCode
		response.Message = err.Err.Error()
		return c.JSON(response.StatusCode, response)
	}

	user.HidePassword()
	response.StatusCode = 200
	response.Message = "success edit user data"
	response.Data = user

	return c.JSON(response.StatusCode, response)

}

func (u *userHandler) DeleteUser(c echo.Context) error {
	response := domains.GeneralResponse{}
	err := domains.ErrorCode{}
	var userID int
	userID, err.Err = strconv.Atoi(c.Param("id"))
	if err.Err != nil {
		response.StatusCode = 400
		response.Message = "id parameter must be valid"
		return c.JSON(response.StatusCode, response)
	}

	err = u.userService.DeleteUserService(userID)
	if err.Err != nil {
		response.StatusCode = err.StatusCode
		response.Message = err.Err.Error()
		return c.JSON(response.StatusCode, response)
	}

	response.StatusCode = 200
	response.Message = "success delete user"
	response.Data = map[string]int{"user_id": userID}

	return c.JSON(response.StatusCode, response)
}

func (u *userHandler) Login(c echo.Context) error {
	loginResponse := domains.LoginResponse{}
	user := domains.User{}
	err := domains.ErrorCode{}

	if err := c.Bind(&user); err != nil {
		if err != nil {
			loginResponse.StatusCode = 400
			loginResponse.Message = err.Error()
			return c.JSON(loginResponse.StatusCode, loginResponse)
		}
	}

	token, err := u.userService.Login(&user)
	if err.Err != nil {
		loginResponse.StatusCode = err.StatusCode
		loginResponse.Message = err.Err.Error()
		return c.JSON(loginResponse.StatusCode, loginResponse)
	}

	loginResponse.StatusCode = 200
	loginResponse.Message = "success login"
	loginResponse.Data = user
	loginResponse.Token = token

	return c.JSON(loginResponse.StatusCode, loginResponse)
}
