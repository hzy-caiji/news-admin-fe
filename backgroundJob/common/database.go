package common

import (
	"common/model"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

//数据库配置
func InitDB() *gorm.DB{
	driverName:= "mysql"
	db, err :=gorm.Open(driverName,"admin:131420@(127.0.0.1:3306)/shcool?charset=utf8&parseTime=True&loc=Local&timeout=3600s")
	if err!=nil {
		panic("failed to connect database,err:"+err.Error())
	}

	//创建数据表
	//db.SetConnMaxLifetime(time.Duration(8*3600) * time.Second)
	db.AutoMigrate(&model.User{},&model.Wirte{})
	DB=db
	return db
}



//获取DB实例
func GetDB() *gorm.DB {
	return DB
}
