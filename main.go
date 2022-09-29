package main

import (
	"assignment2/db"
	"assignment2/server"
	"assignment2/server/controllers"
	"assignment2/server/repositories/postgress"

	"github.com/gin-gonic/gin"
)

// @title Orders API
// @description Assignment 2 Kelas A
// @version v1.0
// @termsOfService http://swagger.io/terms/
// @BasePath /
// @host localhost:4000
// @contact.name Teguh Ridho Afdilla
// @contact.email teguh.afdilla138@gmail.com
func main() {
	db := db.ConnectGorm()

	orderRepo := postgress.NewOrderRepo(db)
	userHandler := controllers.NewOrderController(orderRepo)

	router := gin.Default()
	server.NewRouter(router, userHandler).Start(":4000")
}
