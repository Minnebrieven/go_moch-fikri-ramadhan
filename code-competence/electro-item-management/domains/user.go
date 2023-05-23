package domains

import "errors"

type User struct {
	ID       uint           `json:"id"`
	Name     string         `json:"name"`
	Email    string         `json:"email" gorm:"unique"`
	Password string         `json:"password"`
	Metadata DomainMetadata `json:"metadata" gorm:"embedded"`
}

func (u *User) Validate() ErrorCode {
	err := ErrorCode{}
	switch {
	case u.Name == "":
		err.StatusCode = 400
		err.Err = errors.New("name can't be blank")
		return err
	case u.Email == "":
		err.StatusCode = 400
		err.Err = errors.New("email can't be blank")
		return err
	case u.Password == "":
		err.StatusCode = 400
		err.Err = errors.New("password can't be blank")
		return err
	}
	return err
}

func (u *User) EditValidate() ErrorCode {
	err := ErrorCode{}
	if u.Name == "" && u.Email == "" && u.Password == "" {
		err.StatusCode = 400
		err.Err = errors.New("nothing to change")
	}
	return err
}

func (u *User) HidePassword() {
	u.Password = "********"
}
