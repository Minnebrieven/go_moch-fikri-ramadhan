package services

import (
	"electro-item-management/domains"
	"electro-item-management/repositories"
	"reflect"
)

type (
	ItemService interface {
		GetItemsService(params *domains.ItemParamQuery) ([]domains.Item, int, domains.ErrorCode)
		GetItemsByCategoryService(categoryID int, page *domains.Pages) ([]domains.Item, int, domains.ErrorCode)
		GetItemService(item *domains.Item) domains.ErrorCode
		CreateItemService(itemData *domains.Item) domains.ErrorCode
		EditItemService(itemID string, modifiedItemData *domains.Item) (domains.Item, domains.ErrorCode)
		DeleteItemService(itemID string) domains.ErrorCode
	}

	itemService struct {
		itemRepository     repositories.ItemRepository
		categoryRepository repositories.CategoryRepository
	}
)

func NewItemService(itemRepo repositories.ItemRepository, categoryRepo repositories.CategoryRepository) *itemService {
	return &itemService{itemRepository: itemRepo, categoryRepository: categoryRepo}
}

func (is *itemService) GetItemsService(params *domains.ItemParamQuery) ([]domains.Item, int, domains.ErrorCode) {
	offset, limit := params.Pages.CalcOffsetLimit()

	switch {
	case params.IsNameSet():
		params.SetNameForDBQuery()
		items, totalData, err := is.itemRepository.GetItemsByName(params.Name, offset, limit)
		if err.Err != nil {
			return nil, totalData, err
		}
		return items, totalData, err
	default:
		items, totalData, err := is.itemRepository.GetItems(offset, limit)
		return items, totalData, err
	}
}

func (is *itemService) GetItemsByCategoryService(categoryID int, page *domains.Pages) ([]domains.Item, int, domains.ErrorCode) {
	offset, limit := page.CalcOffsetLimit()
	items, totalData, err := is.itemRepository.GetItemsByCategory(categoryID, offset, limit)
	if err.Err != nil {
		return nil, totalData, err
	}
	return items, totalData, err
}

func (is *itemService) GetItemService(item *domains.Item) domains.ErrorCode {
	err := is.itemRepository.GetItem(item)
	if err.Err != nil {
		return err
	}
	return err
}

func (is *itemService) CreateItemService(itemData *domains.Item) domains.ErrorCode {
	if err := itemData.Validate(); err.Err != nil {
		return err
	}

	if itemData.Category.ID != 0 {
		category := domains.Category{ID: itemData.Category.ID}
		err := is.categoryRepository.GetCategory(&category)
		if err.Err != nil {
			return err
		}
		itemData.Category.ID = category.ID
		itemData.Category.Name = category.Name
		itemData.Category.Metadata.CreatedAt = category.Metadata.CreatedAt
		itemData.Category.Metadata.UpdatedAt = category.Metadata.CreatedAt
	}

	err := is.itemRepository.CreateItem(itemData)
	if err.Err != nil {
		return err
	}
	return err
}

func (is *itemService) EditItemService(itemID string, modifiedItemData *domains.Item) (domains.Item, domains.ErrorCode) {
	//find record first if not exists return error
	item := domains.Item{}

	err := item.InsertID(itemID)
	if err.Err != nil {
		return domains.Item{}, err
	}

	err = is.itemRepository.GetItem(&item)
	if err.Err != nil {
		return domains.Item{}, err
	}

	if err := modifiedItemData.EditValidate(); err.Err != nil {
		return domains.Item{}, err
	}

	if modifiedItemData.Category.ID != 0 {
		category := domains.Category{ID: modifiedItemData.Category.ID}
		err = is.categoryRepository.GetCategory(&category)
		if err.Err != nil {
			return domains.Item{}, err
		}
		item.Category.ID = category.ID
		item.Category.Name = category.Name
		item.Category.Metadata.CreatedAt = category.Metadata.CreatedAt
		item.Category.Metadata.UpdatedAt = category.Metadata.CreatedAt
	}

	//replace exist data with new one
	var itemPointer *domains.Item = &item
	itemVal := reflect.ValueOf(itemPointer).Elem()
	itemType := itemVal.Type()

	editVal := reflect.ValueOf(modifiedItemData).Elem()

	for i := 0; i < itemVal.NumField(); i++ {
		//skip ID, CreatedAt, UpdatedAt field to be edited
		switch itemType.Field(i).Name {
		case "ID":
			continue
		case "CreatedAt":
			continue
		case "UpdatedAt":
			continue
		case "CategoryID":
			continue
		case "Category":
			continue
		}

		editField := editVal.Field(i)
		isSet := editField.IsValid() && !editField.IsZero()
		if isSet {
			itemVal.Field(i).Set(editVal.Field(i))
		}
	}

	err = is.itemRepository.UpdateItem(&item)
	if err.Err != nil {
		return domains.Item{}, err
	}

	return item, err

}

func (is *itemService) DeleteItemService(itemID string) domains.ErrorCode {
	item := domains.Item{}

	err := item.InsertID(itemID)
	if err.Err != nil {
		return err
	}
	err = is.itemRepository.DeleteItem(&item)
	return err
}
