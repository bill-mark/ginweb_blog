package routers

import (
	"blog_model/controllers"
	"blog_model/middleware"
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

   //school接口
   taR := router.Group("/school")
   taR.Use(middleware.JWTAuth())
   {
   	//获得学校总列表
   	taR.GET("/list",controllers.GetSchoolList)
   	//新增学校
   	taR.POST("/add",controllers.AddSchoolPost)
   	//分页查询学校
   	taR.GET("/get",controllers.GetSchoolListFenye)
   	//删除学校
   	taR.POST("/delet",controllers.DeletSchoolWithId)
   	//模糊搜索
   	taR.GET("/likesearch",controllers.LikeSearchSchool)
   	//地区搜索
   	taR.GET("/areasearch",controllers.SchoolArea)
   }

   return router
}