package model

import (
	"errors"
	"fmt"
	"simple/database"
	password2 "simple/pkg/password"
	"time"
)

type User struct {
	ID         uint `gorm:"primary_key"`
	Nickname   string
	Password   string
	Email      string
	CreateTime time.Time
	UpateTime  time.Time
	DeleteTime time.Time
	Avatar     string
	Group_id   int
}

func Ctable() {

	e := database.Db.CreateTable(&User{})
	fmt.Println(e)
}

//添加用户
func CreateUser(nickname string, password string, email string) error{
	encryptPassword := password2.CreatePassword(password, 5)
	if nickname == ""{
		return errors.New("用户名不能为空")
	}else if password == ""{
		return errors.New("密码不能为空")
	}else if email == ""{
		return errors.New("邮箱不能为空")
	}
	user := User{
		Nickname:   nickname,
		Password:   string(encryptPassword),
		Email:      email,
		CreateTime: time.Now(),
	}
	db := database.Db.Create(&user)
	if db.Error != nil {

		return errors.New("失败")

	}else{
		return nil
	}

}

//用户登入
func UserLogin(nickname string,password string)error{
	var user User
	if nickname == ""{
		return errors.New("用户名不能为空")
	}
	if password == ""{
		return errors.New("密码不能为空")
	}
	db := database.Db.Where("nickname = ?", nickname).First(&user)
    if db.Error !=nil{
    	return errors.New("用户不存在")
	}else{
		ok := password2.ComparePassword([]byte(user.Password), []byte(password))
		if !ok{
			return errors.New("密码错误")
		}
	}
	return nil


}

func CheckUser(email string)error{
	var user User
	return database.Db.Where("email= ?", email).First(&user).Error




}

