package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 或者其他数据库驱动
	"github.com/jmoiron/sqlx"
)

func main() {
	//连接
	mysqlDB := connectMySQL()
	defer mysqlDB.Close()

	/*
		Exec（）的两个返回值 一个结果集 和一个错误信息 (Result, error)
	*/
	//插入
	//insertSQL := "insert into user(userid,username,password,avatar,create_time,update_time) value(?,?,?,?,?,?)"
	//result, err := mysqlDB.Exec(insertSQL, 2, "jack", "666", "kl.png", "2023-8-8 12:12:12", "2023-8-9 12:12:12")
	//if err != nil {
	//	fmt.Println("数据插入失败", err)
	//	return
	//}
	//id, _ := result.LastInsertId()
	//fmt.Println("数据插入成功 last id: ", id)

	////修改
	//updateSQL := "update user set username='Eric' where id=7"
	//res2, _ := mysqlDB.Exec(updateSQL)
	//rowNum, _ := res2.RowsAffected() //受到影响的行数
	//fmt.Println("更新成功 受影响行数: ", rowNum)
	//

	////修改
	//deleteSQL := "delete from user where id=7"
	//res3, _ := mysqlDB.Exec(deleteSQL)
	//rowNum, _ := res3.RowsAffected() //受到影响的行数
	//fmt.Println("删除成功 受影响行数: ", rowNum)

	//查询
	//querySQL := "select * from user"
	//rows, err := mysqlDB.Query(querySQL)
	//if err != nil {
	//	fmt.Println("数据查询失败", err)
	//	return
	//}
	////fmt.Println(rows)
	//
	////返回的是结果集
	//for rows.Next() {
	//	var id, userid int
	//	var username, password, avatar, create_time, update_time string
	//
	//	rows.Scan(&id, &userid, &username, &password, &avatar, &create_time, &update_time)
	//	fmt.Println(id, userid, username, password, avatar, create_time, update_time)
	//}
	//rows.Close()

	//使用Get Select 查询
	//创建和数据库一一对应的结构体
	type user struct {
		Id         int    `db:"id"`
		UserId     int    `db:"userid"`
		UserName   string `db:"username"`
		Password   string `db:"password"`
		Avatar     string `db:"avatar"`
		CreateTime string `db:"create_time"`
		UpdateTime string `db:"update_time"`
	}

	//查询一条
	userData := new(user)
	mysqlDB.Get(userData, "select * from user where id=8")
	fmt.Println(userData)

	//查询多条
	var userSlice []user
	mysqlDB.Select(&userSlice, "select * from user")

	err := mysqlDB.Select(&userSlice, "select * from user")
	if err != nil {
		fmt.Println("查询失败:", err)
	}
	for i := range userSlice {
		fmt.Println(userSlice[i])
	}

}

// mysql 连接信息
var (
	userName  string = "root"
	passWord  string = "root"
	ipAddress string = "127.0.0.1"
	port      string = "3306"
	dbName    string = "go_orm_db"
	charset   string = "utf8mb4"
)

// 连接mysql
func connectMySQL() *sqlx.DB {
	dbstr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", userName, passWord, ipAddress, port, dbName, charset)

	DB, err := sqlx.Open("mysql", dbstr)

	fmt.Println(err)

	ping(DB)
	return DB

}

// 测试是否连接成功
func ping(DB *sqlx.DB) {
	err := DB.Ping()
	if err != nil {
		fmt.Println("ping failed")
	} else {
		fmt.Println("ping success")
	}
}
