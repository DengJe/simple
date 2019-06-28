package database

import (
	"testing"
	"log"

)

func TestDb(t *testing.T){
	t.Error("dddd")
	err := connectDb()
	log.Println(err)
}

