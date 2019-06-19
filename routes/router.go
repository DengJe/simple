package routes

import(
	"github.com/dengje/simple/api/cms"
	"github.com/gin-gonic/gin"
	"github.com/dengje/simple/api"
	"github.com/dengje/simple/middleware/jwt"
)

func RegisterRoutes(router *gin.Engine) {
	
	auth := router.Group("/auth")
	{
		auth.GET("/", api.GetAuth)
		
	}

	// 简单的路由组: v1
	apiv1 := router.Group("/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/", cms.CmsTest)
		
	}

	// 简单的路由组: cms
	rcms := router.Group("/cms")

	{
		rcms.GET("/test", cms.CmsTest)
	}
	
}



