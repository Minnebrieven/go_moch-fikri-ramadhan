package repositories

import (
	"electro-item-management/domains"
	"electro-item-management/helpers"
	"errors"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var mysqlErr *mysql.MySQLError

type (
	UserRepository interface {
		GetUsers(offset, limit int) ([]domains.User, int, domains.ErrorCode)
		GetUser(user *domains.User) domains.ErrorCode
		CreateUser(userData *domains.User) domains.ErrorCode
		UpdateUser(userData *domains.User) domains.ErrorCode
		DeleteUser(userData *domains.User) domains.ErrorCode
		Login(user *domains.User) domains.ErrorCode
	}

	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUsers(offset, limit int) ([]domains.User, int, domains.ErrorCode) {
	users := []domains.User{}
	result := ur.db.Limit(limit).Offset(offset).Omit("password").Find(&users)

	err := domains.ErrorCode{}
	if result.Error != nil {
		err.StatusCode = 500
		err.Err = result.Error
		return nil, 0, err
	}

	// return no data
	if result.RowsAffected == 0 {
		err.StatusCode = 200
		err.Err = errors.New("no data")
	}
	return users, int(result.RowsAffected), err
}

func (ur *userRepository) GetUser(user *domains.User) domains.ErrorCode {
	err := domains.ErrorCode{}

	err.Err = ur.db.First(user).Error
	if errors.Is(err.Err, gorm.ErrRecordNotFound) {
		err.StatusCode = 200
		err.Err = errors.New("user not found")
		return err
	}
	return err
}

func (ur *userRepository) CreateUser(userData *domains.User) domains.ErrorCode {
	err := domains.ErrorCode{}
	//hashing password
	var hashedPassword string
	hashedPassword, err.Err = helpers.HashPassword(userData.Password)
	if err.Err != nil {
		err.StatusCode = 500
		err.Err = errors.New("something went wrong: cant hashing password")
		return err
	}
	userData.Password = hashedPassword

	err.Err = ur.db.Create(userData).Error
	if errors.As(err.Err, &mysqlErr) && mysqlErr.Number == 1062 {
		err.StatusCode = 409
		err.Err = errors.New("duplicate user found")
		return err
	}

	return err
}

func (ur *userRepository) UpdateUser(userData *domains.User) domains.ErrorCode {
	err := domains.ErrorCode{}
	err.Err = ur.db.Save(userData).Error
	if err.Err != nil {
		err.StatusCode = 500
	}
	return err
}

func (ur *userRepository) DeleteUser(userData *domains.User) domains.ErrorCode {
	err := domains.ErrorCode{}
	err.Err = ur.db.First(userData).Error
	if errors.Is(err.Err, gorm.ErrRecordNotFound) {
		err.StatusCode = 200
		err.Err = errors.New("user not found")
		return err
	}
	err.Err = ur.db.Delete(userData).Error
	if err.Err != nil {
		err.StatusCode = 500
	}
	return err
}

func (ur *userRepository) Login(user *domains.User) domains.ErrorCode {
	err := domains.ErrorCode{}
	userPassword := user.Password
	err.Err = ur.db.Where("email = ?", user.Email).First(user).Error
	if errors.Is(err.Err, gorm.ErrRecordNotFound) {
		err.StatusCode = 401
		err.Err = errors.New("login failed: wrong email")
		return err
	}

	isMatch := helpers.MatchingPasswordHash(userPassword, user.Password)
	if !isMatch {
		err.StatusCode = 401
		err.Err = errors.New("login failed: wrong password ")
		return err
	}
	return err
}
