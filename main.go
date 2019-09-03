package main

import (
	"blogweb_gin/database"
	"blogweb_gin/routers"
)

func main(){
	database.InitMysql()
	router := routers.InitRouter()
	router.Run(":8081")
}