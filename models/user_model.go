package models

import (
	"blogweb_gin/database"
	"fmt"
)

type User struct {
	Id int
	Username string
	Password string
	Status int   //0正常  1删除
	Createtime int64
}

//插入
func InsetUser(user User)(int64,error){
	return database.ModifyDB("insert into users(username,password,status,createtime)values(?,?,?,?)",user.Username,user.Password,user.Status,user.Createtime)
}

//条件查询
func QueryUserWightCon(con string)int{
	sql :=fmt.Sprintf("select id from users %s",con)
	fmt.Println(sql)
	row := database.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	return id
}

//用户名查询ID
func QueryUserWightUsername(username string)int{
	sql := fmt.Sprintf("where username='%s'",username)
	return QueryUserWightCon(sql)
}