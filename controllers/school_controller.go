package controllers

import (
	myjwt "blog_model/middleware"
	"blog_model/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
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

//分页查询学校
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

//删除学校
func DeletSchoolWithId(c *gin.Context){
	data,_ := ioutil.ReadAll(c.Request.Body)
	fmt.Println(data)
	var cc  map[string]int
	json.Unmarshal(data,&cc)
	//fmt.Println(cc)
	_,err := models.DeletSchoolWithId(cc["id"])
	if err !=nil{
		log.Println(err)
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"message":err,
		})
	}else{
		c.JSON(http.StatusOK,gin.H{
			"code":0,
		})
	}
}

//模糊搜索
func LikeSearchSchool(c *gin.Context){
	key := c.Query("key")

	result,err := models.SearchSchoolLikeKey(key)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(result)
		c.JSON(http.StatusOK,gin.H{
			"code":0,
			"data":result,
		})
	}
}

//地区查询
func SchoolArea(c *gin.Context){
	key := c.Query("key")
	result,err := models.SearchSchoolArea(key)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(result)
		c.JSON(http.StatusOK,gin.H{
			"code":0,
			"data":result,
		})
	}
}