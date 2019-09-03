package routers

import (
	"blogweb_gin/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter()*gin.Engine{
   router := gin.Default()

   //注册
   router.POST("/register",controllers.RegisterPost)

   return router
}