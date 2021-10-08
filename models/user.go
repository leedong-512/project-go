package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"laji/v1/db"
)

type User struct {
	Id       int64 `gorm:"primary_key"`
	Username string
	Age      int64
}

var (
	userTest *db.Mysql
)

/**
 * @author lidong
 * @description 获取表名
 * @date 17:35 2021/9/9
 * @param
 * @return
 **/
func (User) TableName() string {
	return "user_test"
}

func NewUser() *User {
	return &User{}
}

/**
 * @author lidong
 * @description 创建用户
 * @date 17:35 2021/9/9
 * @param
 * @return
 **/
func (u *User) Create() (*gorm.DB, error) {
	userTest, _ = db.GetMysql()
	defer userTest.DB.Close()
	fmt.Println(u.Username)
	data := userTest.DB.Create(&u)

	return data, nil

}
