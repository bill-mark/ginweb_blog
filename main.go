package main

import (
	"blog_model/database"
	"blog_model/routers"
)

func main(){
	database.InitMysql()
	router := routers.InitRouter()
	router.Run(":8081")
}