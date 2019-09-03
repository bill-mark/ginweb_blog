package routers

import (
	"blogweb_gin/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter()*gin.Engine{
   router := gin.Default()

   //注册
   router.POST("/register",controllers.RegisterPost)
   //登录
   router.POST("/login",controllers.LoginPost)
   //更新密码
   router.POST("/update_password",controllers.UpdatePassword)


   return router
}