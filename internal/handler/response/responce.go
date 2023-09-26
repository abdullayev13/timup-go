package response

import (
	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, res interface{}) {
	c.JSON(200, map[string]interface{}{
		"res":    res,
		"status": true,
	})
}
func Fail(c *gin.Context, msg string) {
	c.JSON(200, map[string]interface{}{
		"status": false,
		"msg":    msg,
	})
}

func FailErr(c *gin.Context, err error) {
	Fail(c, err.Error())
}
func FailErrOrMsg(c *gin.Context, err error, msg string) {
	if err != nil {
		Fail(c, err.Error())
	} else {
		Fail(c, msg)
	}
}
