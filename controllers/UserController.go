package controllers

import "fmt"
import (
	"qqchat/models"
	_"github.com/astaxie/beego"
)

type UserController struct {
	baseController
}

//进入登陆页面
func (u *UserController) Login1() {
	u.TplName = "login.html"
}

//进入登陆成功页面
func (u *UserController) Login2() {
	username := u.GetString("username")
	userpassword := u.GetString("userpassword")
	fmt.Println(username, userpassword)
	if username =="" {
   u.Redirect("/userlogin",302)
   return
	}
    user,err:=models.UserLogin(username)
    if err!=nil{
		u.Redirect("/userlogin",302)
		return
	}
	//beego.URLFor("UserController.Func")动态路径
	u.SetSession("sessionid",user)
	u.Redirect("/loginafter/function",302)
}
//进入功能选择页面
func (u *UserController) Func()  {
	sessionid:=u.GetSession("sessionid")
	if sessionid ==nil{
		u.Redirect("/",302)
		return
	}
	u.TplName = "function.html"
}
func (u *UserController) Register1()  {
	u.TplName = "register.html"
}
//用户注册
func (u *UserController) Register2()  {
	username:=u.GetString("username")
	userpassword1:=u.GetString("userpassword1")
	userpassword2:=u.GetString("userpassword2")
	if userpassword1 !=userpassword2{
		u.Redirect("/userlogin/register",302)
		return
	}
	models.UserreRister(username,userpassword1)
	models.In(username)
	models.Ini(username,userpassword1)
	u.TplName = "login.html"
}