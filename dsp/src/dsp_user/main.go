package main

import (
	"dsp/src/util"
	"github.com/gin-gonic/gin"
	userRouter "dsp/src/dsp_user/router"
	"log"
)

/**
 * @author  tianqiang
 * @date  2021/7/13 16:24
 */
func main() {
	util.ReadConfig()
	ok :=util.Connect()
	if !ok{
		log.Println("db connect err!")
	}else {
		log.Println("db connected success!")
	}
	Router := gin.Default()
	rr := GinRouter(Router)
	_ = rr.Run(":9801")

}
func GinRouter(r *gin.Engine) *gin.Engine {
	router := r.Group("/api/dsp-user")
	userRouter.Routers(router)
	return r
}
