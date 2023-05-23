package domains

import "errors"

type Item struct {
	ID          uint           `json:"id"`
	Name        string         `json:"name"`
	CategoryID  uint           `json:"-"`
	Category    Category       `gorm:"constraint:OnUpdate:CASCADE;" json:"category"`
	Description string         `json:"description"`
	Stock       int            `json:"stock"`
	Price       int            `json:"price"`
	Metadata    DomainMetadata `gorm:"embedded" json:"metadata"`
}

func (i *Item) Validate() ErrorCode {
	err := ErrorCode{}
	switch {
	case i.Name == "":
		err.StatusCode = 400
		err.Err = errors.New("name can't be blank")
		return err
	case i.Category.ID == 0 && i.Category.Name == "":
		err.StatusCode = 400
		err.Err = errors.New("category id or category name can't be blank")
		return err
	case i.Description == "":
		err.StatusCode = 400
		err.Err = errors.New("description can't be blank")
		return err
	case i.Stock == 0:
		err.StatusCode = 400
		err.Err = errors.New("stock can't be 0")
		return err
	case i.Price == 0:
		err.StatusCode = 400
		err.Err = errors.New("price can't be 0")
		return err
	}
	return err
}

func (i *Item) EditValidate() ErrorCode {
	err := ErrorCode{}
	if i.Name == "" && i.Category.ID == 0 && i.Description == "" && i.Stock == 0 && i.Price == 0 {
		err.StatusCode = 400
		err.Err = errors.New("nothing to change")
	}
	return err
}
