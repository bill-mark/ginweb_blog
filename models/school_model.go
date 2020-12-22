package models

import (
	"blog_model/database"
	"fmt"
)

type School struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Area string `json:"area"`
}

func InsertSchool(school School) (int64, error) {
	return database.ModifyDB("insert into school(name,area) values (?,?)", school.Name, school.Area)
}

func GetAllSchool() ([]School, error) {
	rows, err := database.QueryDB("select id,name,area from school")
	if err != nil {
		return nil, err
	}
	var alllist []School
	for rows.Next() {
		var school School
		_ = rows.Scan(&school.Id, &school.Name, &school.Area)
		alllist = append(alllist, school)
	}
	return alllist, nil
}

//按页数查询
func FindSchoolWithPage(page int, total int) ([]School, error) {

	page--
	fmt.Println("---------->page", page)
	return QuerySchoolWithPage(page, total)
}

//获取总条数
func QuerySchoolRowNum() int {
	row := database.QueryRowDB("select count(id) from school")
	num := 0
	row.Scan(&num)
	return num
}

/**
limit分页查询语句，
	语法：limit m，n
	m代表从多少位开始获取，与id值无关
	n代表获取多少条数据
*/
func QuerySchoolWithPage(page, num int) ([]School, error) {

	sql := fmt.Sprintf("limit %d,%d", page*num, num)

	return QuerySchoolWithCon(sql)
}

func QuerySchoolWithCon(sql string) ([]School, error) {
	sql = "select id,name,area from school " + sql
	rows, err := database.QueryDB(sql)
	if err != nil {
		return nil, err
	}

	var schoolList []School
	for rows.Next() {
		var school School
		_ = rows.Scan(&school.Id, &school.Name, &school.Area)
		schoolList = append(schoolList, school)
	}
	return schoolList, nil
}

//删除
func DeletSchoolWithId(ID int) (int64, error) {
	return database.ModifyDB("delete from school where id=?", ID)
}

//模糊搜索
func SearchSchoolLikeKey(key string) ([]School, error) {
	rows, err := database.QueryDBLike(key)
	if err != nil {
		return nil, err
	}
	var schoolList []School
	for rows.Next() {
		var school School
		_ = rows.Scan(&school.Id, &school.Name, &school.Area)
		schoolList = append(schoolList, school)
	}
	return schoolList, nil
}

//地区查询
func SearchSchoolArea(key string) ([]School, error) {
	sql := fmt.Sprintf("select * from school where area='%s' ", key)
	fmt.Println(sql)
	rows, err := database.QueryDB(sql)
	if err != nil {
		return nil, err
	}
	var schoolList []School
	for rows.Next() {
		var school School
		_ = rows.Scan(&school.Id, &school.Name, &school.Area)
		schoolList = append(schoolList, school)
	}
	return schoolList, nil
}
