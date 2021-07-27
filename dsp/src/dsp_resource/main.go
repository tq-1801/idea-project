package dsp_resource

import (
	"dsp/src/util"
	"github.com/gin-gonic/gin"
	"log"
)

/**
 * @author  tianqiang
 * @date  2021/7/20 17:13
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
	router := r.Group("/api/dsp-res")
	resRouter.Routers(router)
	return r
}