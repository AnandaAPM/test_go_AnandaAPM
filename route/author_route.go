package route

import (
	"go_test/controller"
	"go_test/middleware"

	"github.com/gin-gonic/gin"
)
func AuthorRoutes(router *gin.Engine,AuthorController *controller.AuthorController){
	
	protected :=router.Group("/authors")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/",AuthorController.GetAllAuthors)
		protected.POST("/",AuthorController.RegisterAuthor)
		protected.GET("/:id",AuthorController.FindAuthorById)
		protected.PUT("/:id",AuthorController.UpdateAuthor)
		protected.DELETE("/:id",AuthorController.DeleteAuthor)
		
	}
		
}