package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	controller "github.com/srikanthbhandary/cleanarch/controllers"
	"github.com/srikanthbhandary/cleanarch/repository"
	"github.com/srikanthbhandary/cleanarch/services"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	DB, err := gorm.Open(mysql.Open("root:password@tcp(127.0.0.1:3306)/cleanarch?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Fatal("Unable to connect to the database")
		os.Exit(1)
	}
	userRepository := repository.NewUserRepository(DB)
	userRepository.Migrate()
	userService := services.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	r := gin.Default()
	r.POST("/api/v1/users", userController.AddUser)
	r.GET("/api/v1/users", userController.GetUsers)
	r.Run(":8080")
}
