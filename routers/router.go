package routers

import (
	"qqchat/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//进入首页
    beego.Router("/", &controllers.MainController{})
    //进入登录页面
    beego.Router("/userlogin",&controllers.UserController{},"Get:Login1")
    //登录成功，
	beego.Router("/loginafter",&controllers.UserController{},"Post:Login2")
	//进入功能选择页面
	beego.Router("/loginafter/function",&controllers.UserController{},"*:Func")
	//进入个人主页
	beego.Router("/loginafter/function/home",&controllers.UserMessageController{},"Get:Look")
	//保存个人信息
    beego.Router("/loginafter/function/home/save",&controllers.UserMessageController{},"Post:Save")
	//进入聊天页面
	beego.Router("/loginafter/function/chat",&controllers.ChatRecordController{},"Get:ChatRecord")
	//进入注册页面
	beego.Router("/userlogin/register",&controllers.UserController{},"Get:Register1")
	//注册完成
	beego.Router("/userlogin/register/after",&controllers.UserController{},"Post:Register2")

	//发送信息/lp/fetch
	beego.Router("/lp/post", &controllers.ChatRecordController{})
	beego.Router("/lp/fetch", &controllers.ChatRecordController{},"get:Fetch")
	}
