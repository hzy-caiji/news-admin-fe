package main

import (
	"common/controller"
	"common/cros"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	//r.Use(middleware.CORSMiddleware())
	r.Use(cros.Cors())			//解决跨域
	r.POST("/api/usr/register", controller.Register)
	r.POST("/api/usr/login", controller.Login)
	//r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)
	r.POST("/api/wirte/wirteart",controller.WirteArt)
	r.GET("/api/wirte/showwir",controller.ShowWirArc)
	r.GET("/api/wirte/getid",controller.GETid)
	r.GET("api/wirte/delwir",controller.DeleteWir)
	r.POST("api/wirte/serch",controller.Serch)

	return r
}
