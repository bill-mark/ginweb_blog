package controllers

import (
	"blogweb_gin/models"
	"blogweb_gin/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type User struct{
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginPost(c *gin.Context){
	fmt.Println("LoginPost")
	var user User
    c.BindJSON(&user)
	fmt.Println(user)

	id := models.QueryUserWithParam(user.Username,utils.MD5(user.Username))
	fmt.Println("id is :",id)
	if id >0{
		c.JSON(http.StatusOK,gin.H{"code":0,"message":"登录成功"})
	}else{
		c.JSON(http.StatusOK,gin.H{"code":-1,"message":"登录失败"})
	}
}

func UpdatePassword(c *gin.Context){
	data, _ := ioutil.ReadAll(c.Request.Body)
	var cc map[string]string
	json.Unmarshal(data,&cc)
	username := cc["username"]
	password := cc["password"]
	newpassword := cc["newpassword"]
    fmt.Println(cc)
	//return
    id := models.QueryUserWithParam(username,utils.MD5(password))
    if id >0{
        _,err := models.UpdateUser(id,username,utils.MD5(newpassword))
		if err != nil{
			c.JSON(http.StatusOK,gin.H{"code":-1,"message":"注册失败"})
		}else{
			c.JSON(http.StatusOK,gin.H{"code":0,"message":"更新密码成功"})
		}
	}else{
		c.JSON(http.StatusOK,gin.H{"code":-1,"message":"账号或密码不正确"})
	}
}