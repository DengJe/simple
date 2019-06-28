package routes

import (
	"github.com/gin-gonic/gin"
	"simple/api"
	"simple/api/cms"
	"simple/middleware/jwt"
)

func RegisterRoutes(router *gin.Engine) {

	//登入授权
	router.POST("/login", api.GetAuth)

	//路由分组：cms
	rcms := router.Group("/cms")
	{
		user := rcms.Group("/user")
		{
			//刷新令牌
			user.POST("/refresh",api.GetAuth)
			user.POST("/create",cms.CreateUser)
		}

	}



	// 简单的路由组: v1
	apiv1 := router.Group("/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/c", cms.CmsTest)
		
	}


	
}



