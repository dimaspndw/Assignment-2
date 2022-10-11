package controllers

import (
	"assigment-2/database"
	"assigment-2/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateOrderInput struct {
	CustormerName string        `json:"CustomerName"`
	OrderedAt     time.Time     `json:"OrderedAt"`
	Items         []models.Item `json:"Items"`
}

type CreateItemInput struct {
	ItemCode    uint   `json:"ItemCode"`
	Description string `json:"Description"`
	Quantity    uint   `json:"Quantity"`
}

func CreateOrder(ctx *gin.Context) {
	db := database.GetDB()
	var input CreateOrderInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	order := models.Order{CustormerName: input.CustormerName, OrderedAt: input.OrderedAt, Items: input.Items}

	err := db.Create(&order).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_message": err,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"orders": order,
	})
}

func GetOrders(ctx *gin.Context) {
	db := database.GetDB()
	var orders []models.Order
	err := db.Preload("Items").Find(&orders).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_message": err,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"datas": orders,
	})
}

func UpdateOrder(ctx *gin.Context) {
	db := database.GetDB()
	var input CreateOrderInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_1": err,
		})
		return
	}
	order := models.Order{CustormerName: input.CustormerName, OrderedAt: input.OrderedAt, Items: input.Items}

	err := db.Where("id = ?", ctx.Param("orderId")).Updates(&order).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_2": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"new data orders": order,
	})
}

func DeleteOrder(ctx *gin.Context) {
	db := database.GetDB()

	order := models.Order{}
	items := models.Item{}

	errDeleteItems := db.Where("order_id = ?", ctx.Param("orderId")).Delete(&items).Error
	err := db.Where("id = ?", ctx.Param("orderId")).Delete(&order).Error

	if errDeleteItems != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error message": err,
		})
		return
	}

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error message": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"delete data": order,
	})
}
