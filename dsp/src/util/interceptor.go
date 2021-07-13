package util

import (
	"dsp/src/dsp_user/model"
	"github.com/gin-gonic/gin"
	"net/http"

)

/*
拦截器
*/
func Interceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := c.Request.Header.Get("User")
		print(uid)
		if uid != "" {
			user := model.User{}
			var count int64 = 0
			DbConn.Where("uid = ?", uid).Find(&user).Count(&count)
			if int(count) < 1 {
				c.Abort()
				c.JSON(http.StatusOK, gin.H{
					"code": -1,
					"msg":  "当前用户不存在",
				})
				return
			} else {
				c.Next()
			}

		}

	}
}
