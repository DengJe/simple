package jwt

import (
	"simple/pkg/e"
	"simple/pkg/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = e.SUCCESS
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			code = e.INVALID_PARAMS
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		} else {
			claims, err := util.VerifyAccessTokenInHeader(token)
			log.Println(claims, err)
			if err != nil {
				code = e.ERROR_AUTH
				c.JSON(http.StatusUnauthorized, gin.H{
					"code": code,
					"msg":  err.Error(),
					"data": claims,
				})
				c.Abort()
				return

			} else {
				c.Next()
				return

			}

		}
	}
}
