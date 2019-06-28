package database

import(
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"errors"
	"simple/pkg/setting"

)

var Db *gorm.DB

func init(){
	connectDb()
}

func connectDb()error{
	user,_ := setting.LoadConfig("database","USER")
	password,_ := setting.LoadConfig("database","PASSWORD")
	dbname,_ := setting.LoadConfig("database","NAME")
	config := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",user,password,dbname)
	var err error
	Db, err = gorm.Open("mysql", config)
	if err != nil {
		return errors.New("数据库连接失败")
	}
	//设置默认表名前缀
	gorm.DefaultTableNameHandler = func(Db *gorm.DB, defaultTableName string) string {
		return "je_" + defaultTableName
	}
	return errors.New("连接成功")

}

func CloseDb(){
	defer Db.Close()
}