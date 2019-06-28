package cms

import(
	"net/http"
	"simple/model"
	"github.com/gin-gonic/gin"
	password2 "simple/pkg/password"
)

func CreateUser(c *gin.Context)  {
	nickname := c.PostForm("nickname")
	password := c.PostForm("password")
	email := c.PostForm("email")
	encryptPassword := password2.CreatePassword(password, 5)
	err := model.CreateUser(nickname,string(encryptPassword),email)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{
			"code" : 200,
			"msg" : err.Error(),
			"data" : 1,
		})
	}else{
		c.JSON(http.StatusOK, gin.H{
			"code" : 200,
			"msg" : err.Error(),
			"data" : 1,
		})
	}


}


