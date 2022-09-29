package controllers

import (
	"assignment2/server/models"
	"assignment2/server/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderRepo repositories.OrderRepository
}

func NewOrderController(orderRepo repositories.OrderRepository) *OrderController {
	return &OrderController{
		orderRepo: orderRepo,
	}
}

func (o *OrderController) AddOrder(ctx *gin.Context) {
	var dataReq models.DataReq
	err := ctx.ShouldBindJSON(&dataReq)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var dataOrder = models.Order{
		Customer_name: dataReq.Customer_name,
		Ordered_at:    time.Now(),
	}

	err = o.orderRepo.AddOrder(&dataOrder)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var dataItem []models.Item

	for _, s := range dataReq.Items {

		dataItem = append(dataItem, models.Item{
			Item_code:   s.Item_code,
			Description: s.Description,
			Quantity:    s.Quantity,
			Order_id:    dataOrder.Order_id,
		})
	}

	err = o.orderRepo.AddItems(&dataItem)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "created order success",
	})
}

func (o *OrderController) GetOrders(ctx *gin.Context) {
	orders, err := o.orderRepo.GetOrders()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"payload": orders,
	})
}

func (o *OrderController) EditOrder(ctx *gin.Context) {
	var dataReq models.DataReq
	err := ctx.ShouldBindJSON(&dataReq)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id := ctx.Param("id")
	intVar, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var dataOrder = models.Order{
		Order_id:      intVar,
		Customer_name: dataReq.Customer_name,
		Ordered_at:    time.Now(),
	}

	err = o.orderRepo.EditOrder(&dataOrder)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	for _, s := range dataReq.Items {
		var dataItem = models.Item{
			Item_code:   s.Item_code,
			Description: s.Description,
			Quantity:    s.Quantity,
		}
		err = o.orderRepo.EditItems(&dataItem, intVar, s.Item_id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "update order success",
	})
}

func (o *OrderController) DeleteOrders(ctx *gin.Context) {
	id := ctx.Param("id")
	intVar, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "DeleteOrders 001: " + err.Error(),
		})
		return
	}

	err = o.orderRepo.DeleteOrder(intVar)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "DeleteOrders 002: " + err.Error(),
		})
		return
	}

	err = o.orderRepo.DeleteItem(intVar)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "DeleteOrders 003: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "deleted order success",
	})
}
