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
	CategoryHandler interface{}

	categoryHandler struct {
		categoryService services.CategoryService
	}
)

func NewCategoryHandler(categoryServ services.CategoryService) *categoryHandler {
	return &categoryHandler{categoryService: categoryServ}
}

func (u *categoryHandler) GetCategories(c echo.Context) error {
	response := domains.GeneralListResponse{}
	err := domains.ErrorCode{}

	pageString := c.QueryParam("page")
	if pageString != "" {
		response.Page, err.Err = strconv.Atoi(pageString)
		if err.Err != nil {
			response.StatusCode = 400
			response.Message = "invalid page parameter"
			return c.JSON(response.StatusCode, response)
		}
	}
	response.DataShown = constants.LIMIT
	offset := 0
	if response.Page > 1 {
		offset = response.Page * response.DataShown
	}

	categories, totalData, err := u.categoryService.GetCategoriesService(offset, response.DataShown)
	if err.Err != nil {
		response.StatusCode = err.StatusCode
		response.Message = err.Err.Error()
		return c.JSON(response.StatusCode, response)
	}

	response.StatusCode = 200
	response.Message = "success"
	response.TotalData = totalData
	response.Data = categories

	return c.JSON(http.StatusOK, response)
}

func (u *categoryHandler) GetCategoryByID(c echo.Context) error {
	response := domains.GeneralResponse{}
	err := domains.ErrorCode{}
	category := domains.Category{}

	categoryIDString := c.Param("id")
	var categoryID int
	categoryID, err.Err = strconv.Atoi(categoryIDString)
	if err.Err != nil {
		response.StatusCode = 400
		response.Message = fmt.Sprintf("invalid id parameter %v", categoryIDString)
		return c.JSON(response.StatusCode, response)
	}
	category.ID = uint(categoryID)

	err = u.categoryService.GetCategoryService(&category)
	if err.Err != nil {
		response.StatusCode = err.StatusCode
		response.Message = err.Err.Error()
		return c.JSON(response.StatusCode, response)
	}

	response.StatusCode = 200
	response.Message = "success"
	response.Data = category

	return c.JSON(http.StatusOK, response)
}

func (u *categoryHandler) CreateCategory(c echo.Context) error {
	response := domains.GeneralResponse{}
	category := domains.Category{}

	if err := c.Bind(&category); err != nil {
		if err != nil {
			response.StatusCode = 400
			response.Message = err.Error()
			return c.JSON(response.StatusCode, response)
		}
	}

	err := u.categoryService.CreateCategoryService(&category)
	if err.Err != nil {
		response.StatusCode = err.StatusCode
		response.Message = err.Err.Error()
		return c.JSON(response.StatusCode, response)
	}

	response.StatusCode = 201
	response.Message = "success create category"
	response.Data = category

	return c.JSON(response.StatusCode, response)
}

func (u *categoryHandler) EditCategory(c echo.Context) error {
	response := domains.GeneralResponse{}
	err := domains.ErrorCode{}
	var categoryID int
	categoryID, err.Err = strconv.Atoi(c.Param("id"))
	if err.Err != nil {
		response.StatusCode = 400
		response.Message = "id parameter must be valid"
		return c.JSON(response.StatusCode, response)
	}

	modifiedCategoryData := domains.Category{}
	if err := c.Bind(&modifiedCategoryData); err != nil {
		response.StatusCode = 400
		response.Message = err.Error()
		return c.JSON(response.StatusCode, response)
	}

	category, err := u.categoryService.EditCategoryService(categoryID, &modifiedCategoryData)
	if err.Err != nil {
		response.StatusCode = err.StatusCode
		response.Message = err.Err.Error()
		return c.JSON(response.StatusCode, response)
	}

	response.StatusCode = 200
	response.Message = "success edit category data"
	response.Data = category

	return c.JSON(response.StatusCode, response)

}

func (u *categoryHandler) DeleteCategory(c echo.Context) error {
	response := domains.GeneralResponse{}
	err := domains.ErrorCode{}
	var categoryID int
	categoryID, err.Err = strconv.Atoi(c.Param("id"))
	if err.Err != nil {
		response.StatusCode = 400
		response.Message = "id parameter must be valid"
		return c.JSON(response.StatusCode, response)
	}

	err = u.categoryService.DeleteCategoryService(categoryID)
	if err.Err != nil {
		response.StatusCode = err.StatusCode
		response.Message = err.Err.Error()
		return c.JSON(response.StatusCode, response)
	}

	response.StatusCode = 200
	response.Message = "success delete category"
	response.Data = map[string]int{"category_id": categoryID}

	return c.JSON(response.StatusCode, response)
}
