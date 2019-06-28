package cms

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"simple/pkg/util"
)

func CmsTest(c *gin.Context){
	e,err :=util.VerifyAccessTokenInHeader("Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjE0NTM1NzgsImlkZW50aWZ5Ijoic3VwZXIiLCJ0eXBlIjoiQUNDRVNTIn0.UP6IRREGDE3XC4cq9rfq3WnuzPq4sDsYYedHIbV4TZw")
	log.Println(e,err)
	c.JSON(http.StatusOK, gin.H{
        "code" : 200,
        "msg" : "this is test",
        "data" : 1,
    })
}

func CmsTest2(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"code" : 200,
		"msg" : "this is test",
		"data" : 1,
	})
}