package controllers

import (
	"assignment-2/helpers"
	"assignment-2/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type ItemController struct {
	db *gorm.DB
}

func NewItemController(db *gorm.DB) *ItemController {
	return &ItemController{
		db: db,
	}
}

type ItemResponse struct {
	Success bool        `json:"success"`
	Data    models.Item `json:"data"`
}

type ItemsResponse struct {
	Success bool                   `json:"success"`
	Data    []models.Item          `json:"data"`
	Query   map[string]interface{} `json:"query"`
}

// FindItems godoc
// @Summary get all items
// @Description get items
// @Tags items
// @Accept json
// @Produce json
// @Success 200 {object} ItemsResponse
// @Router /items [get]
func (controller *ItemController) FindItems(ctx *gin.Context) {
	limit := ctx.Query("limit")
	limitInt := 10

	if limit != "" {
		l, err := strconv.Atoi(limit)
		if err == nil {
			limitInt = l
		}
	}

	var items []models.Item
	var total int64

	err := controller.db.Debug().Limit(limitInt).Find(&items).Count(&total).Error
	if err != nil {
		helpers.BadRequestResponse(ctx, err.Error())
		return
	}

	helpers.WriteJsonResponse(ctx, http.StatusOK, ItemsResponse{
		Success: true,
		Data:    items,
		Query: map[string]interface{}{
			"limit": limitInt,
			"total": total,
		},
	})
}

// FindItemById godoc
// @Summary get item by id
// @Description get item by id
// @Tags items
// @Accept json
// @Produce json
// @Param 		id path string true "id"
// @Success 200 {object} ItemResponse
// @Router /items/{id} [get]
func (controller *ItemController) FindItemById(ctx *gin.Context) {
	id := ctx.Param("id")
	var item models.Item

	err := controller.db.Debug().Where("id = ?", id).First(&item).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helpers.NotFoundResponse(ctx, "Item data not found")
			return
		}
		helpers.BadRequestResponse(ctx, err.Error())
		return
	}

	helpers.WriteJsonResponse(ctx, http.StatusOK, ItemResponse{
		Success: true,
		Data:    item,
	})
}

// UpdateItem godoc
// @Summary update item by id
// @Description update item by id
// @Tags items
// @Accept json
// @Produce json
// @Param 		id path string true "id"
// @Router /items/{id} [put]
func (controller *ItemController) UpdateItem(ctx *gin.Context) {
	id := ctx.Param("id")
	var newItem models.Item

	err := ctx.ShouldBindJSON(&newItem)
	if err != nil {
		helpers.BadRequestResponse(ctx, err.Error())
		return
	}

	err = controller.db.First(&newItem, id).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helpers.NotFoundResponse(ctx, "Item data not found")
			return
		}
		helpers.BadRequestResponse(ctx, err.Error())
		return
	}

	err = controller.db.
		Debug().
		Model(&newItem).
		Where("id = ?", id).Updates(newItem).Error
	//Error
	if err != nil {
		helpers.BadRequestResponse(ctx, err.Error())
		return
	}

	helpers.WriteJsonResponse(ctx, http.StatusOK, gin.H{
		"success": true,
		"result":  newItem,
	})
}

// DeleteItem godoc
// @Summary delete item by id
// @Description delete item by id
// @Tags items
// @Accept json
// @Produce json
// @Param 		id path string true "id"
// @Router /items/{id} [delete]
func (controller *ItemController) DeleteItem(ctx *gin.Context) {
	id := ctx.Param("id")
	var item models.Item

	err := controller.db.First(&item, id).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helpers.NotFoundResponse(ctx, "Item data not found")
			return
		}
		helpers.BadRequestResponse(ctx, err.Error())
		return
	}

	err = controller.db.Debug().Delete(&item).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helpers.NotFoundResponse(ctx, err.Error())
			return
		}
		helpers.BadRequestResponse(ctx, err.Error())
		return
	}

	helpers.WriteJsonResponse(ctx, http.StatusOK, gin.H{
		"success": true,
		"message": fmt.Sprintf("item_id %d has been successfully deleted", item.ID),
	})
}
