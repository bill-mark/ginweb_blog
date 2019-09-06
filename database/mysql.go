package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_"github.com/jmoiron/sqlx"
	_ "github.com/jmoiron/sqlx"

	"fmt"
	"log"
)

var db *sql.DB

func InitMysql(){
	fmt.Println("InitMysql")
	if db == nil{
		// db,_ = sqlx.Open("mysql","root:FF558234ttpk@tcp(127.0.0.1:3306)/blogweb_gin?charset=utf8")
		db, _ = sql.Open("mysql", "root:FF558234ttpk@tcp(127.0.0.1:3306)/blogweb_gin")
		CreateTableWithUser()
		CreateTableWithSchool()
	}
}

func CreateTableWithUser(){
	sql := `CREATE TABLE IF NOT EXISTS users(
        id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
        username VARCHAR(64),
        password VARCHAR(64),
        status INT(4),
        createtime INT(10)
    );`
	ModifyDB(sql)
}

//操作数据库
func ModifyDB(sql string,args ...interface{})(int64,error){
	result,err := db.Exec(sql,args...)
	if err != nil{
		log.Println(err)
		return 0,err  //0表示正常状态，1表示删除
	}
	count,err := result.RowsAffected()
	if err != nil{
		log.Println(err)
		return 0,err
	}
	return count,nil
}

//精确查询
func QueryRowDB(sql string) *sql.Row{
	return db.QueryRow(sql)
}

//查询
func QueryDB(sql string)(*sql.Rows,error){
	return db.Query(sql)
}

func CreateTableWithSchool(){
	sql := `create table if not exists school(
            id int(4) primary key auto_increment not null,
            name  varchar(64),
            area  varchar(64)
           );`
	ModifyDB(sql)
}