package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang-gorm-mysql/ConnectDB"
	"golang-gorm-mysql/model"
	"log"
)

func DbConn(MyUser, Password, Host, Db string, Port int) *gorm.DB {
	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", MyUser,Password, Host, Port, Db )
	db, err := gorm.Open("mysql", connArgs)
	if err != nil {
		log.Fatal(err)
	}
	db.SingularTable(true)
	return db
}


func main() {

	cn := ConnectDB.ConnInfo{
		"root",
		"123456aA@",
		"127.0.0.1",
		3306,
		"test",
	}
	db := DbConn(cn.MyUser,cn.Password,cn.Host,cn.Db,cn.Port)
	defer db. Close ()
	//Ex1: create table
	//db.DropTableIfExists(&model.Classroom{}) // Drops table nếu table có tồn tại
	//// Sau khi kết nối tới database ta tạo bảng bằng lệnh CreateTable()
	//db.CreateTable(&model.Classroom{})// Kiểm tra table có được tạo thành công hay không
	//log.Println("created table ", db.HasTable(&model.Classroom{}))

	//Ex2: Insert
	//var classroom = model.Classroom{Name:"Lớp 10A"}
	//db.Create(&classroom)

	//Ex3: lay ra ban ghi dau tien , order by la primary key
	//var classroom = model.Classroom{}
	//db.First(&classroom)
	//// SELECT * FROM users ORDER BY id LIMIT 1;
	//fmt.Println(classroom.ID)

	//Ex4: truy van dung cu phap where
	//var classroom = model.Classroom{}
	//db.Where("name = ?", "Lớp 10B").First(&classroom)
	//fmt.Println(classroom.ID)

	//Ex5: update
	var classroom = model.Classroom{}
	db.First(&classroom)
	classroom.Name = "Lớp 10A - K31"
	db.Save(classroom)
}