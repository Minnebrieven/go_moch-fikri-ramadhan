package services

import (
	"electro-item-management/domains"
	"electro-item-management/repositories"
)

type (
	CategoryService interface {
		GetCategoriesService(offset, limit int) ([]domains.Category, int, domains.ErrorCode)
		GetCategoryService(category *domains.Category) domains.ErrorCode
		CreateCategoryService(categoryData *domains.Category) domains.ErrorCode
		EditCategoryService(categoryID int, modifiedCategoryData *domains.Category) (domains.Category, domains.ErrorCode)
		DeleteCategoryService(categoryID int) domains.ErrorCode
	}

	categoryService struct {
		category repositories.CategoryRepository
	}
)

func NewCategoryService(categoryRepo repositories.CategoryRepository) *categoryService {
	return &categoryService{category: categoryRepo}
}

func (cs *categoryService) GetCategoriesService(offset, limit int) ([]domains.Category, int, domains.ErrorCode) {
	categories, totalData, err := cs.category.GetCategories(offset, limit)
	if err.Err != nil {
		return nil, totalData, err
	}
	return categories, totalData, err
}

func (cs *categoryService) GetCategoryService(category *domains.Category) domains.ErrorCode {
	err := cs.category.GetCategory(category)
	if err.Err != nil {
		return err
	}
	return err
}

func (cs *categoryService) CreateCategoryService(categoryData *domains.Category) domains.ErrorCode {
	if err := categoryData.Validate(); err.Err != nil {
		return err
	}
	err := cs.category.CreateCategory(categoryData)
	if err.Err != nil {
		return err
	}
	return err
}

func (cs *categoryService) EditCategoryService(categoryID int, modifiedCategoryData *domains.Category) (domains.Category, domains.ErrorCode) {
	//find record first if not existsCategoryurn error
	category := domains.Category{ID: uint(categoryID)}
	err := cs.category.GetCategory(&category)
	if err.Err != nil {
		return domains.Category{}, err
	}

	if err := modifiedCategoryData.Validate(); err.Err != nil {
		return domains.Category{}, err
	}

	//replace exist data with new one
	category.Name = modifiedCategoryData.Name

	err = cs.category.UpdateCategory(&category)
	if err.Err != nil {
		return domains.Category{}, err
	}

	return category, err

}

func (cs *categoryService) DeleteCategoryService(categoryID int) domains.ErrorCode {
	category := domains.Category{ID: uint(categoryID)}
	err := cs.category.DeleteCategory(&category)
	return err
}
