package main

import (
	"common/common"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func main() {
	db:=common.InitDB()
	defer db.Close()			//延迟关闭

	r := gin.Default()
	r=CollectRoute(r)

	panic(r.Run(":8090")) // listen and serve on 0.0.0.0:8080
}

