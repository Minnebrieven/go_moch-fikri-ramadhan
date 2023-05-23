package repositories

import (
	"electro-item-management/domains"
	"errors"

	"gorm.io/gorm"
)

type (
	ItemRepository interface {
		GetItems(offset, limit int) ([]domains.Item, int, domains.ErrorCode)
		GetItemsByCategory(categoryID, offset, limit int) ([]domains.Item, int, domains.ErrorCode)
		GetItemsByName(itemName string, offset, limit int) ([]domains.Item, int, domains.ErrorCode)
		GetItem(item *domains.Item) domains.ErrorCode
		CreateItem(itemData *domains.Item) domains.ErrorCode
		UpdateItem(itemData *domains.Item) domains.ErrorCode
		DeleteItem(itemData *domains.Item) domains.ErrorCode
	}

	itemRepository struct {
		db *gorm.DB
	}
)

func NewItemRepository(db *gorm.DB) *itemRepository {
	return &itemRepository{db}
}

func (ir *itemRepository) GetItems(offset, limit int) ([]domains.Item, int, domains.ErrorCode) {
	items := []domains.Item{}
	result := ir.db.Limit(limit).Offset(offset).Preload("Category").Find(&items)

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
	return items, int(result.RowsAffected), err
}

func (ir *itemRepository) GetItemsByName(itemName string, offset, limit int) ([]domains.Item, int, domains.ErrorCode) {
	items := []domains.Item{}
	result := ir.db.Where("name LIKE ?", itemName).Limit(limit).Offset(offset).Preload("Category").Find(&items)

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
	return items, int(result.RowsAffected), err
}

func (ir *itemRepository) GetItemsByCategory(categoryID int, offset, limit int) ([]domains.Item, int, domains.ErrorCode) {
	items := []domains.Item{}
	result := ir.db.Where("category_id = ?", categoryID).Limit(limit).Offset(offset).Preload("Category").Find(&items)

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
	return items, int(result.RowsAffected), err
}

func (ir *itemRepository) GetItem(item *domains.Item) domains.ErrorCode {
	err := domains.ErrorCode{}

	err.Err = ir.db.Preload("Category").First(item).Error
	if errors.Is(err.Err, gorm.ErrRecordNotFound) {
		err.StatusCode = 200
		err.Err = errors.New("item not found")
		return err
	}
	return err
}

func (ir *itemRepository) CreateItem(itemData *domains.Item) domains.ErrorCode {
	err := domains.ErrorCode{}

	err.Err = ir.db.Create(itemData).Error
	if errors.As(err.Err, &mysqlErr) && mysqlErr.Number == 1062 {
		err.StatusCode = 409
		err.Err = errors.New("duplicate item found")
		return err
	}

	return err
}

func (ir *itemRepository) UpdateItem(itemData *domains.Item) domains.ErrorCode {
	err := domains.ErrorCode{}
	err.Err = ir.db.Save(itemData).Error
	if err.Err != nil {
		err.StatusCode = 500
	}
	return err
}

func (ir *itemRepository) DeleteItem(itemData *domains.Item) domains.ErrorCode {
	err := domains.ErrorCode{}
	err.Err = ir.db.First(itemData).Error
	if errors.Is(err.Err, gorm.ErrRecordNotFound) {
		err.StatusCode = 200
		err.Err = errors.New("item not found")
		return err
	}
	err.Err = ir.db.Delete(itemData).Error
	if err.Err != nil {
		err.StatusCode = 500
	}
	return err
}
