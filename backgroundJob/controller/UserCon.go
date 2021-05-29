package controller

import (
	"common/common"
	"common/model"
	"common/util"
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"strings"
)

func Register(c *gin.Context) {

	DB:=common.GetDB()

	//获取参数
	name:=c.PostForm("name")
	number:=c.PostForm("number")
	password:=c.PostForm("password")

	//数据验证
	if len(number)!=10{
		c.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"msg":"学号必须为10位"})
		return
	}
	if len(password)<=6{
		c.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"msg":"密码不能少于6位"})
		return
	}
	//如果人名为空，就随机一个10位字符串
	if len(name)==0{
		name=util.Random(10)
	}

	//验证是否用户存在
	if innumber(DB,number){
		c.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"msg":"用户已经存在"})
		return
	}

	//创建用户
	//加密密码
	/*hasedPassword,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"code":500,"msg":"加密错误"})
		return
	}*/
	md:=Encode(password)
	newUser:=model.User{
		Name: name,
		Number: number,
		Password: md,
	}
	DB.Create(&newUser)


	//返回结果
	c.JSON(200,gin.H{
		"code":200,
		"msg":"注册成功",
	})
}

//验证学号是否存在
func innumber(db *gorm.DB, number string) bool{
	var user model.User
	db.Where("number=?",number).Find(&user)
	if user.ID!=0{
		return true
	}

	return false
}

func Login(lg *gin.Context) {

	DB:=common.GetDB()
	//获取参数
	/*number:=lg.PostForm("number")
	password:=lg.PostForm("password")*/
	type Username struct {
		Number string `json:"number"`
		Password string `json:"password"`
	}
	json :=Username{}
	lg.BindJSON(&json)
	log.Println("%v",&json)

	//数据验证
	if len(json.Number)!=10{
		lg.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"msg":"学号必须为10位"})
		return
	}
	if len(json.Password)<=6{
		lg.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"msg":"密码不能少于6位"})
		return
	}

	//判断是否存在
	var user model.User
	DB.Where("number=?",json.Number).Find(&user)
	if user.ID==0{
		lg.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"msg":"用户不存在"})
		return
	}

	if Check(json.Password,user.Password)!=true{
		lg.JSON(http.StatusBadRequest,gin.H{"code":400,"msg":"密码错误"})
		return
	}


	//发放token
	token:="11"

	//返回结果
	lg.JSON(200,gin.H{
		"code":200,
		"data":gin.H{"token":token,
			"data":user},
		"msg":"登录成功",
	})
}

func FinUser(c *gin.Context)  {
	DB:=common.GetDB()
	//获取参数
	c.Get("/")
	var cat []model.User
	var qcat []model.User
	DB.Scopes(util.Paginate(c)).Find(&cat)			//分页查询
	DB.Find(&qcat)
	c.JSON(http.StatusOK,gin.H{
		"code":0,
		"msg":"",
		"len":len(qcat),
		"count":1000,
		"data":cat,
	})
}



//MD5加密
//Check判断是否相等
func Check(content, encrypted string) bool {
	return strings.EqualFold(Encode(content), encrypted)
}
//加密
func Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
