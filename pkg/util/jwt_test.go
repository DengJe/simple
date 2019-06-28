package util

import (
	"log"
	"testing"
)


func TestGenerRerfreshToken(t *testing.T){
	New()
	//token,err := GenerateAccessToken("jjj")
	//refresh,err := GenerateRefreshToken("refresh")
	//
	//log.Println("access_token",token,err)
	//log.Println("refresh_token",refresh,err)
	res,err := VerifyAccessTokenInHeader("Bearer "+"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjE0NTc3NDMsImlkZW50aWZ5IjoiampqIiwidHlwZSI6IkFDQ0VTUyJ9.JX2iNN7AWvZSu5PjBJ6naDvTXcnv2k_vvDp1IVHRyOA")
	log.Println(res,err)
	//res,err = VerifyRefreshToken(refresh)
	//log.Println(res,err)



}



