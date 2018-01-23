package controllers

import "qqchat/models"

type UserMessageController struct {
	baseController
}

//用户信息查看
func (um *UserMessageController) Look()  {
	sessionid:=um.GetSession("sessionid")
	sess:=sessionid.(models.User)
	if sessionid ==nil{
		um.Redirect("/",302)
		return
	}
	usermessage:=models.MessageLook(sess.Username)
	um.Data["usermessage1"] = usermessage.Homeusername
	um.Data["usermessage2"] = usermessage.Homeuserpassword
	um.Data["usermessage3"] = usermessage.Homeusercall
	um.Data["usermessage4"] = usermessage.Homeuserage
	um.Data["usermessage5"] = usermessage.Homeusersex
	um.Data["usermessage6"] = usermessage.Homeuserinterest
	um.TplName = "home.html"
}
//用户信息保存
func (um *UserMessageController) Save()  {
	sessionid:=um.GetSession("sessionid")
	if sessionid ==nil{
		um.Redirect("/",302)
		return
	}
	usm:=models.UserMessage{}
    usm.Homeusername=um.GetString("homeusername")
	usm.Homeuserpassword=um.GetString("homeuserpassword")
	usm.Homeusercall=um.GetString("homeusercall")
	usm.Homeuserage=um.GetString("homeuserage")
	usm.Homeusersex=um.GetString("homeusersex")
	usm.Homeuserinterest=um.GetString("homeuserinterest")


	models.MessageSave(usm)
	um.Redirect("loginafter/function/home",302)
}