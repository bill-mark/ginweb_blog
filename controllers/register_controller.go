package controllers

import (
	"blogweb_gin/models"
	"blogweb_gin/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func RegisterGet(c *gin.Context){
	fmt.Println("get  this is get")
   c.HTML(http.StatusOK,"register.html",gin.H{"title":"注册页"})
}

//处理注册
func RegisterPost(c *gin.Context){
    fmt.Println("RegisterPost")
	username := c.PostForm("username")
	password := c.PostForm("password")
	repassword := c.PostForm("repassword")
	fmt.Println(username,password,repassword)

	id := models.QueryUserWightUsername(username)
	if id > 0{
		c.JSON(http.StatusOK,gin.H{"code":0,"message":"用户已存在"})
		return
	}else{
		fmt.Println("新用户,  下一步")
	}

    //注册用户
    password = utils.MD5(password)
    fmt.Println("md5:",password)

    user := models.User{0,username,password,0,time.Now().Unix()}
    _,err := models.InsetUser(user)
    if err != nil{
    	c.JSON(http.StatusOK,gin.H{"code":0,"message":"注册失败"})
	}else{
		c.JSON(http.StatusOK,gin.H{"code":0,"message":"注册成功"})
	}
}