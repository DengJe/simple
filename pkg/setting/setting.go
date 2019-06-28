package setting


import (
    "errors"
    "log"
    "github.com/go-ini/ini"
    )

var (
    Cfg *ini.File

)

func init() {
    var err error

    Cfg, err = ini.Load("C:/workspace/go/test/simple/conf/app.ini")
    if err != nil {
        log.Printf("Fail to parse 'conf/app.ini': %v", err)
    }



}

func LoadConfig(section string,key string)(string,error){

    sec, err := Cfg.GetSection(section)
    if err != nil {
        return section,errors.New("有误")
    }
    //判断key是否存在
    yes := Cfg.Section(section).HasKey(key)
    if yes == false {
        return key,errors.New("有误")
    }
    config := sec.Key(key).String()

    return config,nil

}

