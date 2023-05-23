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
	ItemHandler interface{}

	itemHandler struct {
		itemService services.ItemService
	}
)

func NewItemHandler(itemServ services.ItemService) *itemHandler {
	return &itemHandler{itemService: itemServ}
}

func (u *itemHandler) GetItems(c echo.Context) error {
	response := domains.GeneralListResponse{}
	err := domains.ErrorCode{}

	itemParam := domains.ItemParamQuery{Name: c.QueryParam("keyword")}
	itemParam.PageString = c.QueryParam("page")
	err.Err = itemParam.ToPageINT()
	if err.Err != nil {
		response.StatusCode = 400
		response.Message = "invalid page parameter: " + itemParam.PageString
		return c.JSON(response.StatusCode, response)
	}

	items, totalData, err := u.itemService.GetItemsService(&itemParam)
	if err.Err != nil {
		response.StatusCode = err.StatusCode
		response.Message = err.Err.Error()
		return c.JSON(response.StatusCode, response)
	}

	response.StatusCode = 200
	response.Message = "success"
	response.Page = itemParam.Page
	response.DataShown = constants.LIMIT
	response.TotalData = totalData
	response.Data = items

	return c.JSON(http.StatusOK, response)
}

func (u *itemHandler) GetItemsByCategory(c echo.Context) error {
	response := domains.GeneralListResponse{}
	err := domains.ErrorCode{}

	page := domains.Pages{PageString: c.QueryParam("page")}
	err.Err = page.ToPageINT()
	if err.Err != nil {
		response.StatusCode = 400
		response.Message = "invalid page parameter: " + page.PageString
		return c.JSON(response.StatusCode, response)
	}

	categoryIDString := c.Param("category_id")
	var categoryID int
	categoryID, err.Err = strconv.Atoi(categoryIDString)
	if err.Err != nil {
		response.StatusCode = 400
		response.Message = fmt.Sprintf("invalid id parameter %v", categoryIDString)
		return c.JSON(response.StatusCode, response)
	}

	items, totalData, err := u.itemService.GetItemsByCategoryService(categoryID, &page)
	if err.Err != nil {
		response.StatusCode = err.StatusCode
		response.Message = err.Err.Error()
		return c.JSON(response.StatusCode, response)
	}

	response.StatusCode = 200
	response.Message = "success"
	response.DataShown = constants.LIMIT
	response.TotalData = totalData
	response.Data = items

	return c.JSON(http.StatusOK, response)
}

func (u *itemHandler) GetItemByID(c echo.Context) error {
	response := domains.GeneralResponse{}
	err := domains.ErrorCode{}
	item := domains.Item{}

	itemIDString := c.Param("id")
	var itemID int
	itemID, err.Err = strconv.Atoi(itemIDString)
	if err.Err != nil {
		response.StatusCode = 400
		response.Message = fmt.Sprintf("invalid id parameter %v", itemIDString)
		return c.JSON(response.StatusCode, response)
	}
	item.ID = uint(itemID)

	err = u.itemService.GetItemService(&item)
	if err.Err != nil {
		response.StatusCode = err.StatusCode
		response.Message = err.Err.Error()
		return c.JSON(response.StatusCode, response)
	}

	response.StatusCode = 200
	response.Message = "success"
	response.Data = item

	return c.JSON(http.StatusOK, response)
}

func (u *itemHandler) CreateItem(c echo.Context) error {
	response := domains.GeneralResponse{}
	item := domains.Item{}

	if err := c.Bind(&item); err != nil {
		if err != nil {
			response.StatusCode = 400
			response.Message = err.Error()
			return c.JSON(response.StatusCode, response)
		}
	}

	err := u.itemService.CreateItemService(&item)
	if err.Err != nil {
		response.StatusCode = err.StatusCode
		response.Message = err.Err.Error()
		return c.JSON(response.StatusCode, response)
	}

	response.StatusCode = 201
	response.Message = "success create item"
	response.Data = item

	return c.JSON(response.StatusCode, response)
}

func (u *itemHandler) EditItem(c echo.Context) error {
	response := domains.GeneralResponse{}
	err := domains.ErrorCode{}
	var itemID int
	itemID, err.Err = strconv.Atoi(c.Param("id"))
	if err.Err != nil {
		response.StatusCode = 400
		response.Message = "id parameter must be valid"
		return c.JSON(response.StatusCode, response)
	}

	modifiedItemData := domains.Item{}
	if err := c.Bind(&modifiedItemData); err != nil {
		response.StatusCode = 400
		response.Message = err.Error()
		return c.JSON(response.StatusCode, response)
	}

	item, err := u.itemService.EditItemService(itemID, &modifiedItemData)
	if err.Err != nil {
		response.StatusCode = err.StatusCode
		response.Message = err.Err.Error()
		return c.JSON(response.StatusCode, response)
	}

	response.StatusCode = 200
	response.Message = "success edit item data"
	response.Data = item

	return c.JSON(response.StatusCode, response)

}

func (u *itemHandler) DeleteItem(c echo.Context) error {
	response := domains.GeneralResponse{}
	err := domains.ErrorCode{}
	var itemID int
	itemID, err.Err = strconv.Atoi(c.Param("id"))
	if err.Err != nil {
		response.StatusCode = 400
		response.Message = "id parameter must be valid"
		return c.JSON(response.StatusCode, response)
	}

	err = u.itemService.DeleteItemService(itemID)
	if err.Err != nil {
		response.StatusCode = err.StatusCode
		response.Message = err.Err.Error()
		return c.JSON(response.StatusCode, response)
	}

	response.StatusCode = 200
	response.Message = "success delete item"
	response.Data = map[string]int{"item_id": itemID}

	return c.JSON(response.StatusCode, response)
}
