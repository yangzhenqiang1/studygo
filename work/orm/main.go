package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//新增修改的model
type User struct {
	gorm.Model
	ID       uint `gorm:"primaryKey"`
	Name     string
	Age      int       `gorm:"default:0"`
	Birthday time.Time `gorm:"default:NULL"`
}

//查询的model
type Users struct {
	// gorm.Model
	ID   uint `gorm:"primaryKey"`
	Name string
	Age  int `gorm:"default:0"`
}

func main() {
	// db, err := gorm.Open("mysql", "root:yangzhenqiang@(192.168.0.113:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	dsn := "root:yangzhenqiang@tcp(192.168.0.113:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	/*Create*/
	// 迁移 在数据库中创建Users表
	// db.AutoMigrate(&User{})
	// return
	/*使用struct创建数据*/
	// user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	// result := db.Create(&user) // 通过数据的指针来创建
	// db.Select("Name", "Age", "CreatedAt").Create(&user)
	// db.Omit("Name", "Age", "CreatedAt").Create(&user)
	// fmt.Printf("id:%v,error:%v,rowAffected:%v", user.ID, result.Error, result.RowsAffected)
	// user.ID             // 返回插入数据的主键
	// result.Error        // 返回 error
	// result.RowsAffected // 返回插入记录的条数
	/*使用map创建数据*/
	//不如struct好，区别类似php的model
	// user1 := map[string]interface{}{"Name": "Yzq", "Age": 18, "Birthday": time.Now()}
	// res := db.Table("users").Create(user1)
	// fmt.Printf("%v", res)
	/*批量插入*/
	// var users = []User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
	// db.Create(&users)
	// for _, user2 := range users {
	// 	fmt.Printf("%v", user2.ID) // 1,2,3
	// }

	/*Read*/
	// var user User
	// db.First(&user)//select * from users order by id limit 1
	// db.Last(&user) //select * from users order by id desc limit 1
	// result := db.Take(&user) //select * from users limit 1
	// fmt.Printf("user:%v\nresult.Error:%v\nresult.RowsAffected :%v\n", user, result.Error, result.RowsAffected)
	//定义一个值为User结构体类型的切片
	// users := []User{}//定义切片方式1
	users := make([]Users, 1, 15) //定义切片方式2
	//获取users表中全部数据内容
	result := db.Find(&users) //select * from users
	fmt.Printf("全部数据：%v\n", users)
	fmt.Printf("找到的条数：%v\n错误信息：%v\n", result.RowsAffected, result.Error)
	fmt.Printf("len:%v\ncap:%v\n", len(users), cap(users))
	//有条件的查询
	// res := db.Where("name = ?", "Yzq").Find(&users)//获取全部符合条件数据
	// fmt.Printf("name=Yzq1的用户数据信息：%v\n找到的条数：%v\n错误信息：%v\n", users, res.RowsAffected, res.Error)
	// var user Users
	// res := db.Where("name = ?", "Yzq1").First(&user) //获取符合条件数据id正序第一条
	// fmt.Printf("name=Yzq1的用户数据信息：%v\n找到的条数：%v\n错误信息：%v\n", user, res.RowsAffected, res.Error)
	result = db.Where("age in ?", []int{120}).Find(&users)
	fmt.Printf("IN 数据信息：%v\n找到的条数：%v\n错误信息：%v\n", users, result.RowsAffected, result.Error)
	fmt.Printf("%v\n", lastWeek)

}
