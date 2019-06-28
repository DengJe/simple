package api

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"simple/pkg/util"
	"simple/model"


)



func GetAuth(c *gin.Context) {
	nickname := c.PostForm("nickname")
	password := c.PostForm("password")
	err := model.UserLogin(nickname,password)
	data := make(map[string]interface{})
	if err != nil{
		c.JSON(http.StatusOK, gin.H{
			"code" : 200,
			"msg" : err.Error(),
			"data" : data,
		})
	}else{
		token,refresh, _ := util.GenerateTokens(nickname)
		data["access_token"] = token
		data["refresh_token"] = refresh
		c.JSON(http.StatusOK, gin.H{
			"code" : 200,
			"msg" : "授权成功",
			"data" : data,
		})
	}

}