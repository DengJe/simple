package util

import (

	"log"
	"testing"


)


func TestGenerRerfreshToken(t *testing.T){
	token,refresh, e := JwtInstance.GenerateTokens("pedro")
	if e != nil {
		log.Println(t, e)
	}
	log.Println(token,refresh)




}



