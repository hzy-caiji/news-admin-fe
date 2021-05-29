package controller

import (
	"common/common"
	"common/model"
	"common/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)


func WirteArt(wr *gin.Context)  {
	DB:=common.GetDB()
	type Article struct {
		Name string `json:"name"`
		Title string `json:"title"`
		News_type string `json:"news_type"`
		Author string `json:"author"`
		Content string `json:"content"`
		Time string `json:"time"`
		Image string `json:"image"`
	}
	json:=Article{}
	wr.BindJSON(&json)
	newWir:=model.Wirte{
		Name:json.Name,
		Title: json.Title,
		News_tpye: json.News_type,
		Author: json.Author,
		Content: json.Content,
		Time: json.Time,
		Image: json.Image,
	}
	DB.Create(&newWir)
	//返回结果
	wr.JSON(200,gin.H{
		"code":200,
		"data":gin.H{
			"data":newWir},
		"msg":"发布成功",
	})
}

func ShowWirArc(swa *gin.Context)  {
	DB:=common.GetDB()
	//获取参数
	swa.Get("/")
	var cat []model.Wirte
	var qcat []model.Wirte
	DB.Scopes(util.Paginate(swa)).Find(&cat)			//分页查询
	DB.Find(&qcat)
	swa.JSON(http.StatusOK,gin.H{
		"code":0,
		"msg":"",
		"len":len(qcat),
		"count":1000,
		"data":cat,
	})
}

func GETid(gid *gin.Context)  {
	DB:=common.GetDB()
	Id, _ := strconv.Atoi(gid.Query("id"))
	var con model.Wirte
	DB.Where("id=?",Id).Find(&con)

	gid.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"",
		"count":1000,
		"data":con,
	})

}

func DeleteWir(de *gin.Context)  {
	DB:=common.GetDB()
	id, _ := strconv.Atoi(de.Query("id"))
	//var con model.Wirte
	DB.Delete(model.Wirte{},id)
	de.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"删除成功",
	})
}

func Serch(se *gin.Context)  {
	db:=common.GetDB()
	type Writer struct {
		Title string `json:"title"`
	}
	json :=Writer{}
	se.BindJSON(&json)
	log.Println("%v",&json)

	var tit []model.Wirte
	var qtit []model.Wirte
	db.Where("title=?",json.Title).Find(&tit)
	se.JSON(http.StatusOK,gin.H{
		"code":0,
		"msg":"",
		"len":len(qtit),
		"count":1000,
		"data":tit,
	})


}
