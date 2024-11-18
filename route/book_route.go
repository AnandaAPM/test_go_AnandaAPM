package route

import (
	"go_test/controller"
	"go_test/middleware"

	"github.com/gin-gonic/gin"
)
func BookRoutes(router *gin.Engine,bookController *controller.BookController){
	
	protected :=router.Group("/books")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/",bookController.GetAllBooks)
		protected.POST("/",bookController.RegisterBook)
		protected.GET("/:id",bookController.FindBookById)
		protected.PUT("/:id",bookController.UpdateBook)
		protected.DELETE("/:id",bookController.DeleteBook)
	}
		
}