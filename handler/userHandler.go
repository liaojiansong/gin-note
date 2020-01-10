package handler

import (
	"database/sql"
	"gin/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const OK int = http.StatusOK
const ROOT string = "D:\\code\\gin\\"
const AVATAR_PATH string = "avatar\\"
// 主页
func Index(context *gin.Context)  {
	context.HTML(OK,"index.tmpl",gin.H{
		"title":"你好呀",
		"method" :"你的请求方式为" + strings.ToLower(context.Request.Method),
	})

}

func UserRegister(ctx *gin.Context)  {
	//email := ctx.PostForm("email")
	//password := ctx.DefaultPostForm("password", "123456")
	//passwordAgain := ctx.DefaultPostForm("password-again", "123456")
	//ctx.String(http.StatusOK,"%s %s %s",email,password,passwordAgain)
	var user model.UserModel
	// 将user模型绑定到表单
	err := ctx.ShouldBind(&user);
	if err != nil {
		log.Println("发生错误",err.Error())
		ctx.String(http.StatusBadRequest,"请求参数不合法\n")
		return
	}
	user.Save()
	log.Printf("%#v\n",user)

	ctx.Redirect(http.StatusMovedPermanently,"/")
}

func UserLogin(ctx *gin.Context)  {
	var user model.UserModel
	err := ctx.Bind(&user)
	if err != nil {
		log.Panicf("登入绑定失败\n")
	}
	u := user.QueryByEmail()
	if u.Password==user.Password {
		log.Printf("登入成功")
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"email": u.Email,
			"id":    u.Id,
		})
	}else {
		goError(err,ctx,"登入失败")
	}


}

// 个人资料
func UserProfile(ctx *gin.Context)  {
	sid := ctx.Query("id")
	// 声明实例用于调用方法
	var user model.UserModel
	// 转化为整型

	id, err := strconv.Atoi(sid)
	if err != nil {
		goError(err,ctx,"转换id失败")
	}

	userModel, err := user.QueryById(id)
	if err != nil {
		goError(err, ctx,"查找失败")
	}
	log.Printf(userModel.Avatar.String)
	// 响应html
	ctx.HTML(OK,"user_profile.tmpl",gin.H{
		"user":   userModel,
		"avatar": `http://localhost:8099/avatar/`+userModel.Avatar.String,
	})

}

// 编辑资料

func UpdateUserProfile(ctx *gin.Context)  {
	var user model.UserModel
	err := ctx.ShouldBind(&user)
	if err != nil {
		goError(err,ctx,"绑定发生了错误")
	}

	//拿取文件
	file, err := ctx.FormFile("avatar-file")
	if err != nil {
		goError(err,ctx,"文件上传失败")
	}
	// 拼接唯一文件名
	fileName := strconv.FormatInt(time.Now().Unix(),10) + file.Filename
	err = ctx.SaveUploadedFile(file, ROOT+AVATAR_PATH+fileName)
	goError(err,ctx,"无法保存文件")
	avatarUrl := fileName
	user.Avatar=sql.NullString{String:avatarUrl}
	err = user.Update(user.Id)
	goError(err,ctx,"无法更新数据")
	ctx.Redirect(http.StatusMovedPermanently,"/user/profile?id=" +"1")

}


// 响应错误页面
func goError(err error,ctx *gin.Context, msg string)  {
	if err != nil {
		ctx.HTML(OK,"error.tmpl",gin.H{
			"error" : err,
		})
		log.Panicf(msg+":%s", err.Error())
	}
}
