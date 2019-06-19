package jwt

import (
    "net/http"
    "github.com/gin-gonic/gin"
	"github.com/dengje/you-cms/pkg/util"
	"github.com/dengje/you-cms/pkg/e"

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
                "code" : code,
                "msg" : e.GetMsg(code),
                "data" : data,
            })

            c.Abort()
            return
		} else {

			claims, err := util.VerifyAccessTokenInHeader(token)
			if err != nil {
				code = e.ERROR_AUTH
				c.JSON(http.StatusUnauthorized, gin.H{
					"code" : code,
					"msg" : err.Error(),
					"data" : claims,
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