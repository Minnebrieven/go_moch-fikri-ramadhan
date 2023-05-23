package repositories

import (
	"electro-item-management/domains"
	"errors"

	"gorm.io/gorm"
)

type (
	CategoryRepository interface {
		GetCategories(offset, limit int) ([]domains.Category, int, domains.ErrorCode)
		GetCategory(category *domains.Category) domains.ErrorCode
		CreateCategory(categoryData *domains.Category) domains.ErrorCode
		UpdateCategory(categoryData *domains.Category) domains.ErrorCode
		DeleteCategory(categoryData *domains.Category) domains.ErrorCode
	}

	categoryRepository struct {
		db *gorm.DB
	}
)

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db}
}

func (cr *categoryRepository) GetCategories(offset, limit int) ([]domains.Category, int, domains.ErrorCode) {
	categories := []domains.Category{}
	result := cr.db.Limit(limit).Offset(offset).Find(&categories)

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
	return categories, int(result.RowsAffected), err
}

func (cr *categoryRepository) GetCategory(category *domains.Category) domains.ErrorCode {
	err := domains.ErrorCode{}

	err.Err = cr.db.First(category).Error
	if errors.Is(err.Err, gorm.ErrRecordNotFound) {
		err.StatusCode = 200
		err.Err = errors.New("category not found")
		return err
	}
	return err
}

func (cr *categoryRepository) CreateCategory(categoryData *domains.Category) domains.ErrorCode {
	err := domains.ErrorCode{}

	err.Err = cr.db.Create(categoryData).Error
	if errors.As(err.Err, &mysqlErr) && mysqlErr.Number == 1062 {
		err.StatusCode = 409
		err.Err = errors.New("duplicate category found")
		return err
	}

	return err
}

func (cr *categoryRepository) UpdateCategory(categoryData *domains.Category) domains.ErrorCode {
	err := domains.ErrorCode{}
	err.Err = cr.db.Save(categoryData).Error
	if err.Err != nil {
		err.StatusCode = 500
	}
	return err
}

func (cr *categoryRepository) DeleteCategory(categoryData *domains.Category) domains.ErrorCode {
	err := domains.ErrorCode{}
	err.Err = cr.db.First(categoryData).Error
	if errors.Is(err.Err, gorm.ErrRecordNotFound) {
		err.StatusCode = 200
		err.Err = errors.New("category not found")
		return err
	}
	err.Err = cr.db.Delete(categoryData).Error
	if err.Err != nil {
		err.StatusCode = 500
	}
	return err
}
