package server

import (
	"assignment2/server/controllers"

	"github.com/gin-gonic/gin"
)

type Router struct {
	router *gin.Engine
	order  *controllers.OrderController
}

func NewRouter(router *gin.Engine, order *controllers.OrderController) *Router {
	return &Router{
		router: router,
		order:  order,
	}
}

func (r *Router) Start(port string) {

	r.router.GET("/v1/order", r.order.GetOrders)
	r.router.POST("/v1/order", r.order.AddOrder)
	r.router.PUT("/v1/order/:id", r.order.EditOrder)
	r.router.DELETE("/v1/order/:id", r.order.DeleteOrders)
	r.router.Run(port)
}
