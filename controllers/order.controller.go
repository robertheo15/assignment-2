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

type OrderController struct {
	db *gorm.DB
}

func NewOrderController(db *gorm.DB) *OrderController {
	return &OrderController{
		db: db,
	}
}

func (controller *OrderController) CreateOrder(ctx *gin.Context) {
	var newOrder models.Order

	err := ctx.ShouldBindJSON(&newOrder)
	if err != nil {
		helpers.BadRequestResponse(ctx, err.Error())
		return
	}

	err = controller.db.Debug().Create(&newOrder).Error
	if err != nil {
		helpers.BadRequestResponse(ctx, err.Error())
		return
	}

	helpers.WriteJsonResponse(ctx, http.StatusCreated, gin.H{
		"success": true,
		"data":    newOrder,
	})
}

func (controller *OrderController) FindOrders(ctx *gin.Context) {
	limit := ctx.Query("limit")
	limitInt := 10

	if limit != "" {
		l, err := strconv.Atoi(limit)
		if err == nil {
			limitInt = l
		}
	}

	var orders []models.Order
	var total int64

	err := controller.db.Debug().Limit(limitInt).Preload("Items").Find(&orders).Count(&total).Error
	if err != nil {
		helpers.BadRequestResponse(ctx, err.Error())
		return
	}

	helpers.WriteJsonResponse(ctx, http.StatusOK, gin.H{
		"success": true,
		"data":    orders,
		"query": map[string]interface{}{
			"limit": limitInt,
			"total": total,
		},
	})
}

func (controller *OrderController) FindOrderById(ctx *gin.Context) {
	id := ctx.Param("id")
	var order models.Order

	err := controller.db.Debug().Preload("Items").Where("id = ?", id).First(&order).Error
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
		"data":    order,
	})
}

func (controller *OrderController) UpdateOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	var order models.Order

	err := ctx.ShouldBindJSON(&order)
	if err != nil {
		helpers.BadRequestResponse(ctx, err.Error())
		return
	}

	var newOrder = models.Order{
		CustomerName: order.CustomerName,
		Items:        order.Items,
	}

	err = controller.db.Preload("Items").First(&order, id).Error
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
		Model(&order).
		Association("Items").
		Replace(newOrder.Items)
	//Error
	if err != nil {
		helpers.BadRequestResponse(ctx, err.Error())
		return
	}

	helpers.WriteJsonResponse(ctx, http.StatusOK, gin.H{
		"success": true,
		"result":  order,
	})
}

func (controller *OrderController) DeleteOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	var order models.Order

	err := controller.db.First(&order, id).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helpers.NotFoundResponse(ctx, "Order data not found")
			return
		}
		helpers.BadRequestResponse(ctx, err.Error())
		return
	}

	err = controller.db.Debug().Model(&order).Association("Items").Clear()
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
		"message": fmt.Sprintf("order_id %d has been successfully deleted", order.ID),
	})
}
