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

	helpers.WriteJsonResponse(ctx, http.StatusOK, gin.H{
		"success": true,
		"data":    items,
		"query": map[string]interface{}{
			"limit": limitInt,
			"total": total,
		},
	})
}

func (controller *ItemController) FindItemById(ctx *gin.Context) {
	id := ctx.Param("id")
	var item models.Item

	err := controller.db.Debug().Where("id = ?", id).First(&item).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helpers.NotFoundResponse(ctx, "Order data not found")
			return
		}
		helpers.BadRequestResponse(ctx, err.Error())
		return
	}

	helpers.WriteJsonResponse(ctx, http.StatusOK, gin.H{
		"success": true,
		"data":    item,
	})
}

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
			helpers.NotFoundResponse(ctx, "Order data not found")
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
