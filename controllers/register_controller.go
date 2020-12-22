package controllers

import (
	"blog_model/models"
	"blog_model/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
)


//处理注册
func RegisterPost(c *gin.Context){
    fmt.Println("RegisterPost")

	data, _ := ioutil.ReadAll(c.Request.Body)
	var cc map[string]string
	json.Unmarshal(data,&cc)
	username := cc["username"]
	password := cc["password"]

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
    	c.JSON(http.StatusOK,gin.H{"code":-1,"message":"注册失败"})
	}else{
		c.JSON(http.StatusOK,gin.H{"code":0,"message":"注册成功"})
	}
}