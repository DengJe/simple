package cms

import(
	"net/http"
	"github.com/gin-gonic/gin"
)

func CmsTest(c *gin.Context){

	
	c.JSON(http.StatusOK, gin.H{
        "code" : 200,
        "msg" : "this is test",
        "data" : 1,
    })
}