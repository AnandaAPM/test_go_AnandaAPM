package main

import (
	"go_test/config"
	"go_test/controller"
	"go_test/models"
	"go_test/repository"
	"go_test/route"
	"go_test/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	
	config.Init()
	if config.DB == nil {
		log.Fatal("Database not initialized")
	}

	err := config.DB.AutoMigrate(&models.User{},&models.Book{},&models.Author{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	userRepo := repository.UserRepo(config.DB)
	userService := service.UserServices(userRepo)
	userController := controller.UserControllerDI(userService)

	bookRepo := repository.BookRepo(config.DB)
	bookService := service.BookServices(bookRepo)
	bookController := controller.BookControllerDI(bookService)

	authorRepo := repository.AuthorRepo(config.DB)
	authorService := service.AuthorServices(authorRepo)
	authorController := controller.AuthorControllerDI(authorService)

	
	r := gin.Default()
	route.UserRoutes(r,userController)
	route.BookRoutes(r,bookController)
	route.AuthorRoutes(r,authorController)
	r.GET("/",func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,gin.H{"message":"Hello"})
	})


	r.Run(":3000")
}