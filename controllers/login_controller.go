package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	myjwt "blog_model/middleware"

	"blog_model/models"
	"blog_model/utils"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type User struct {
	Id       int    `json:"userid"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//登录
func LoginPost(c *gin.Context) {
	fmt.Println("LoginPost")
	var user User
	_ = c.BindJSON(&user)
	fmt.Println(user.Id, user.Username, user.Password)

	id := models.QueryUserWithParam(user.Username, utils.MD5(user.Password))
	fmt.Println("id is :", id)
	if id > 0 {
		generateToken(c, id, user.Username)
	} else {
		c.JSON(http.StatusOK, gin.H{"code": -1, "message": "登录失败"})
	}
}

// 生成令牌
func generateToken(c *gin.Context, id int, username string) {
	j := &myjwt.JWT{
		[]byte("dengzixiang"),
	}

	claims := myjwt.CustomClaims{
		ID:   id,
		Name: username,
		StandardClaims: jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),    // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600000), // 过期时间 一千小时
			Issuer:    "dengzixiang",                      //签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "登录成功！",
		"token":   token,
	})
	return
}

//更新密码
func UpdatePassword(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)
	var cc map[string]string
	json.Unmarshal(data, &cc)
	username := cc["username"]
	password := cc["password"]
	newpassword := cc["newpassword"]
	fmt.Println(cc)
	//return
	id := models.QueryUserWithParam(username, utils.MD5(password))
	if id > 0 {
		_, err := models.UpdateUser(id, username, utils.MD5(newpassword))
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": -1, "message": "注册失败"})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": 0, "message": "更新密码成功"})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"code": -1, "message": "账号或密码不正确"})
	}
}
