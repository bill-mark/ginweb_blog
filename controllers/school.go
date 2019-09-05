package controllers

import (
	myjwt "blogweb_gin/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetList(c *gin.Context){
	isPass := c.GetBool("isPass")
	if !isPass {
		return
	}
	claims := c.MustGet("claims").(*myjwt.CustomClaims)
	if claims != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "token有效",
			"data":   claims,
		})
	}else{
		c.JSON(http.StatusOK,gin.H{"code":-100,"message":"token wrong"})
	}

}
