package controllers

import (
	myjwt "blogweb_gin/middleware"
	"blogweb_gin/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

//新增学校
func AddSchoolPost(c *gin.Context){
	data, _ := ioutil.ReadAll(c.Request.Body)
	var cc models.School
	json.Unmarshal(data,&cc)
	//fmt.Printf("%+v",cc)
	_,err := models.InsertSchool(cc)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{"code":-1,"message":"失败"})
	}else{
		c.JSON(http.StatusOK,gin.H{"code":0,"message":"成功"})
	}
}

//获得学校全部
func GetSchoolList(c *gin.Context){
	isPass := c.GetBool("isPass")
	if !isPass {
		return
	}
	claims := c.MustGet("claims").(*myjwt.CustomClaims)
	if claims != nil {
		cc,_ := models.GetAllSchool()
        //fmt.Printf("%+v",cc)
		c.JSON(http.StatusOK,gin.H{
			"code":0,
			"data":cc,
		})
	}else{
		c.JSON(http.StatusOK,gin.H{"code":-100,"message":"token wrong"})
	}
}

func GetSchoolListFenye(c *gin.Context){
	var allnum = models.QuerySchoolRowNum()

	page,_ := strconv.Atoi(c.Query("page"))
	total,_ := strconv.Atoi(c.Query("total"))
    cc,_ := models.FindSchoolWithPage(page,total)

	c.JSON(http.StatusOK,gin.H{
		"code":0,
		"allnumber":allnum,
		"data":cc,
	})
}