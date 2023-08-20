package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"xorm.io/xorm"
)

func main() {
	// mysql 连接信息
	var (
		userName  string = "root"
		passWord  string = "root"
		ipAddress string = "127.0.0.1"
		port      string = "3306"
		dbName    string = "go_orm_db"
		charset   string = "utf8mb4"
	)
	//构建数据库连接信息
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", userName, passWord, ipAddress, port, dbName, charset)

	engine, err := xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		fmt.Println("数据库连接失败： ", err)
	}

	type Customer struct {
		Id      int64
		Name    string
		Age     int
		Passwd  string    `xorm:"varchar(200)"`
		Created time.Time `xorm:"created"`
		Updated time.Time `xorm:"updated"`
	}

	//支持数据库表与结构体之间的灵活映射
	err = engine.Sync2(new(Customer))
	if err != nil {
		fmt.Println("表结构同步失败：", err)
	}

	//插入
	//customer := Customer{Id: 100, Name: "Alice", Age: 18, Passwd: "abc123"}
	//one, err := engine.InsertOne(&customer)
	//fmt.Println("插入", one, "条数据")

	////插入多条数据
	//customer2 := Customer{Id: 1, Name: "jack", Age: 18, Passwd: "abc123"}
	//customer3 := Customer{Id: 2, Name: "eric", Age: 28, Passwd: "abc123"}
	//customer4 := Customer{Id: 4, Name: "chrios", Age: 38, Passwd: "abc123"}
	//insert, err := engine.Insert(customer2, customer3, customer4)
	//if insert >= 1 {
	//	fmt.Println("插入", insert, "条数据")
	//}

	////切片插入多条
	//var customers []Customer
	//customers = append(customers, Customer{Id: 11, Name: "liu", Age: 18, Passwd: "abc123"})
	//customers = append(customers, Customer{Id: 12, Name: "zhang", Age: 18, Passwd: "abc123"})
	//customers = append(customers, Customer{Id: 13, Name: "wu", Age: 18, Passwd: "abc123"})
	//insert, err := engine.Insert(&customers)
	//if insert >= 1 {
	//	fmt.Println("插入", insert, "条数据")
	//}

	//修改
	//cus := Customer{Name: "Mark"}
	//affected, err := engine.ID(1).Update(&cus)
	//fmt.Println("受影响行数: ", affected)

	//cus := Customer{Age: 25}
	//affected, err := engine.Update(&cus, &Customer{Name: "Mark"})
	//// UPDATE user SET ... Where name = ?
	//fmt.Println("受影响行数: ", affected)

	//删除
	//cus := Customer{Name: "wu"}
	//affected, err := engine.ID(100).Delete(&cus)
	//// DELETE FROM user Where id = ?
	//fmt.Println("受影响行数: ", affected)

	// Exec
	//engine.Exec("update customer set age=? where id=?", 24, 4)

	//查询
	//slice, err := engine.Query("select * from customer")
	//fmt.Println(slice)
	//
	//queryString, err := engine.QueryString("select * from customer")
	//fmt.Println(queryString)
	//
	//queryInterface, err := engine.QueryInterface("select * from customer")
	//fmt.Println(queryInterface)

	//Get 获得第一条数据 单条
	//customer := Customer{}
	//engine.Get(&customer)
	//fmt.Println(customer)

	//事务
	session := engine.NewSession()
	defer session.Close()

	session.Begin()
	defer func() {
		err := recover()
		if err != nil {
			//回滚
			fmt.Println(err)
			session.Rollback()
		} else {
			session.Commit()
		}
	}()

	cus := Customer{Id: 99, Name: "lin", Age: 18, Passwd: "aaa"}
	if _, err := session.Insert(&cus); err != nil {
		panic(err)
	}

}
