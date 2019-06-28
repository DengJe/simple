package model

import (
	"fmt"
	"testing"
)

func TestDb(t *testing.T){

	//Ctable()
	//AddUser("super","123456","123456@qq.com")
	err := UserLogin("super","123456")
	err = CheckUser("1234560@qq.com")
	fmt.Println(err)
}

