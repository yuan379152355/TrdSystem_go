package main

import (
	"fmt"
	//"GO/TrdSystem/public"
	tspublic "TrdSystem/public"
)

func main() {
	// 定义公共句柄
	var sYBase tspublic.SYBase

	// 初始化数据库连接
	if !sYBase.SYDatabase.Init() {
		fmt.Println("open mysql failed!")
	}

	result, err := sYBase.SYDatabase.Db.Exec("INSERT INTO my_use (userid,username,password) VALUES (1002,'yuanshuai','123456')")
	if err != nil {
		fmt.Println("insert failed,error： ", err)
		return
	}
	id, _ := result.LastInsertId()
	fmt.Println("insert id is :", id)
	_, err1 := sYBase.SYDatabase.Db.Exec("update my_use set username = 'yuanshuai01' where userid = 1001")
	if err1 != nil {
		fmt.Println("update failed error:", err1)
	} else {
		fmt.Println("update success!")
	}
	_, err2 := sYBase.SYDatabase.Db.Exec("delete from my_use where userid = 1001 ")
	if err2 != nil {
		fmt.Println("delete error:", err2)
	} else {
		fmt.Println("delete success")
	}

}

//insert id is : 1
//update success!
//delete success
