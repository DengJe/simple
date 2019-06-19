package api

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/dengje/you-cms/pkg/util"


)



func GetAuth(c *gin.Context) {
	util.New()
	token,refresh, _ := util.JwtInstance.GenerateTokens("jjjj")
	data := make(map[string]interface{})
	data["token"] = token
	data["refresh_token"] = refresh
	c.JSON(http.StatusOK, gin.H{
        "code" : 200,
        "msg" : "授权成功",
        "data" : data,
    })
}