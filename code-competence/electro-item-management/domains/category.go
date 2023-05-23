package domains

import "errors"

type Category struct {
	ID       uint           `json:"id"`
	Name     string         `json:"name" gorm:"unique"`
	Metadata DomainMetadata `json:"metadata" gorm:"embedded"`
}

func (c *Category) Validate() ErrorCode {
	err := ErrorCode{}
	switch {
	case c.Name == "":
		err.StatusCode = 400
		err.Err = errors.New("name can't be blank")
		return err
	}
	return err
}
