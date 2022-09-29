package controllers

import (
	"assignment2/server/models"
	"assignment2/server/repositories"
	"assignment2/server/views"
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

// AddOrder godoc
// @Summary Add new order
// @Decription Add new order
// @Tags order
// @Accept json
// @Produce json
// @Param data body models.DataReq true "Add New Order"
// @Success 200 {object} views.GetAllPeopleSwagger
// @Router /v1/order [post]
func (o *OrderController) AddOrder(ctx *gin.Context) {
	var dataReq models.DataReq
	err := ctx.ShouldBindJSON(&dataReq)
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "ADD_ORDER_FAILED",
			Error:   err.Error(),
		})
		return
	}

	var dataOrder = models.Order{
		Customer_name: dataReq.Customer_name,
		Ordered_at:    time.Now(),
	}

	err = o.orderRepo.AddOrder(&dataOrder)
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "ADD_ORDER_FAILED",
			Error:   err.Error(),
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
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "ADD_ORDER_FAILED",
			Error:   err.Error(),
		})
		return
	}

	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "ADD_ORDER_SUCCESS",
	})
}

// GetOrders godoc
// @Summary Get all order
// @Decription Get all order
// @Tags order
// @Accept json
// @Produce json
// @Success 200 {object} views.GetAllPeopleSwagger
// @Router /v1/order [get]
func (o *OrderController) GetOrders(ctx *gin.Context) {
	orders, err := o.orderRepo.GetOrders()
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "GET_ORDER_FAILED",
			Error:   err.Error(),
		})
		return
	}
	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "GET_ORDER_SUCCESS",
		Payload: orders,
	})
}

// EditOrder godoc
// @Summary Update data order
// @Decription Update data order
// @Tags order
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Param data body models.DataReq true "Update New Order"
// @Success 200 {object} views.GetAllPeopleSwagger
// @Router /v1/order [put]
func (o *OrderController) EditOrder(ctx *gin.Context) {
	var dataReq models.DataReq
	err := ctx.ShouldBindJSON(&dataReq)
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "UPDATE_ORDER_FAILED",
			Error:   err.Error(),
		})
		return
	}

	id := ctx.Param("id")
	intVar, err := strconv.Atoi(id)
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "UPDATE_ORDER_FAILED",
			Error:   err.Error(),
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
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "UPDATE_ORDER_FAILED",
			Error:   err.Error(),
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
			WriteJsonResponse(ctx, &views.Response{
				Status:  http.StatusInternalServerError,
				Message: "UPDATE_ORDER_FAILED",
				Error:   err.Error(),
			})
			return
		}
	}
	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "UPDATE_ORDER_SUCCESS",
	})
}

// DeleteOrders godoc
// @Summary Delete data order
// @Decription Delete data order
// @Tags order
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} views.GetAllPeopleSwagger
// @Router /v1/order/{id} [delete]
func (o *OrderController) DeleteOrders(ctx *gin.Context) {
	id := ctx.Param("id")
	intVar, err := strconv.Atoi(id)
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "DELETE_ORDER_FAILED",
			Error:   err.Error(),
		})
		return
	}

	err = o.orderRepo.DeleteOrder(intVar)
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "DELETE_ORDER_FAILED",
			Error:   err.Error(),
		})
		return
	}

	err = o.orderRepo.DeleteItem(intVar)
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "DELETE_ORDER_FAILED",
			Error:   err.Error(),
		})
		return
	}
	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "DELETE_ORDER_SUCCESS",
	})

}
