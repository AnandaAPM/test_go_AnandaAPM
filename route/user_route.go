package route

import (
	"go_test/controller"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine,userController *controller.UserController){
	
	public :=router.Group("/auth")
	{
		public.POST("/register",userController.RegisterUser)
		public.POST("/login",userController.LoginUser)
	}
		
}