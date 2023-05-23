package services

import (
	"electro-item-management/domains"
	"electro-item-management/helpers"
	"electro-item-management/middlewares"
	"electro-item-management/repositories"
	"reflect"
)

type (
	UserService interface {
		GetUsersService(offset, limit int) ([]domains.User, int, domains.ErrorCode)
		GetUserService(user *domains.User) domains.ErrorCode
		CreateUserService(userData *domains.User) domains.ErrorCode
		EditUserService(userID int, modifiedUserData *domains.User) (domains.User, domains.ErrorCode)
		DeleteUserService(userID int) domains.ErrorCode
		Login(user *domains.User) (string, domains.ErrorCode)
	}

	userService struct {
		userRepository repositories.UserRepository
	}
)

func NewUserService(userRepo repositories.UserRepository) *userService {
	return &userService{userRepository: userRepo}
}

func (us *userService) GetUsersService(offset, limit int) ([]domains.User, int, domains.ErrorCode) {
	users, totalData, err := us.userRepository.GetUsers(offset, limit)
	if err.Err != nil {
		return nil, totalData, err
	}
	return users, totalData, err
}

func (us *userService) GetUserService(user *domains.User) domains.ErrorCode {
	err := us.userRepository.GetUser(user)
	if err.Err != nil {
		return err
	}
	user.HidePassword()
	return err
}

func (us *userService) CreateUserService(userData *domains.User) domains.ErrorCode {
	if err := userData.Validate(); err.Err != nil {
		return err
	}
	err := us.userRepository.CreateUser(userData)
	if err.Err != nil {
		return err
	}
	userData.HidePassword()
	return err
}

func (us *userService) EditUserService(userID int, modifiedUserData *domains.User) (domains.User, domains.ErrorCode) {
	//find record first if not exists return error
	user := domains.User{ID: uint(userID)}
	err := us.userRepository.GetUser(&user)
	if err.Err != nil {
		return domains.User{}, err
	}

	if err := modifiedUserData.EditValidate(); err.Err != nil {
		return domains.User{}, err
	}
	if modifiedUserData.Password != "" {
		modifiedUserData.Password, err.Err = helpers.HashPassword(modifiedUserData.Password)
		if err.Err != nil {
			return domains.User{}, err
		}
	}

	//replace exist data with new one
	var userPointer *domains.User = &user
	userVal := reflect.ValueOf(userPointer).Elem()
	userType := userVal.Type()

	editVal := reflect.ValueOf(modifiedUserData).Elem()

	for i := 0; i < userVal.NumField(); i++ {
		//skip ID, CreatedAt, UpdatedAt field to be edited
		switch userType.Field(i).Name {
		case "ID":
			continue
		case "CreatedAt":
			continue
		case "UpdatedAt":
			continue
		}

		editField := editVal.Field(i)
		isSet := editField.IsValid() && !editField.IsZero()
		if isSet {
			userVal.Field(i).Set(editVal.Field(i))
		}
	}

	err = us.userRepository.UpdateUser(&user)
	if err.Err != nil {
		return domains.User{}, err
	}

	return user, err

}

func (us *userService) DeleteUserService(userID int) domains.ErrorCode {
	user := domains.User{ID: uint(userID)}
	err := us.userRepository.DeleteUser(&user)
	return err
}

func (us *userService) Login(user *domains.User) (string, domains.ErrorCode) {
	var token string
	err := us.userRepository.Login(user)
	if err.Err != nil {
		return "", err
	}

	user.HidePassword()
	token, err.Err = middlewares.CreateToken(int(user.ID), user.Email)
	if err.Err != nil {
		err.StatusCode = 500
		return "", err
	}
	return token, err
}
