package main

import (
	"context"
	userRouter "dsp/src/dsp_user/router"
	"dsp/src/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

/**
 * @author  tianqiang
 * @date  2021/7/20 16:29
 */
func main() {
	util.SysInit()
	Router := gin.Default()
	Router.Use(util.Interceptor())
	rr := GinRouter(Router)

	srv := &http.Server{
		Addr: "127.0.0.1:9000",
		Handler: rr,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n",err)

		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit,os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

}

func GinRouter(r *gin.Engine) *gin.Engine {

	userManageRouter := r.Group("/api/dsp-user")
	userRouter.Routers(userManageRouter)

	resourceManageRouter := r.Group("/api/sf-res")
	resRouter.Routers(resourceManageRouter)



	return r

}