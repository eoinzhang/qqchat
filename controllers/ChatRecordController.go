package controllers

import (
	"qqchat/models"
	"fmt"
)

type ChatRecordController struct {
	baseController
}
//进入聊天室
func (cr *ChatRecordController) ChatRecord()  {
	sessionid:=cr.GetSession("sessionid")
	if sessionid ==nil{
		cr.Redirect("/",302)
		return
	}
	//加入聊天的方法

	sess:=sessionid.(models.User)
	cr.Data["Username"]= sess.Username
	cr.TplName = "chat.html"
}

func (cr *ChatRecordController) Post()  {
	cr.TplName = "chat.html"
	uname:=cr.GetString("uname")
	content:=cr.GetString("content")
	fmt.Println(uname,content)
	if len(uname)==0||len(content)==0{
		return
	}
	models.Post2(uname,content)



}
func (this *ChatRecordController) Fetch()  {

	us:=models.Fetch2()
	if us==nil{
		this.Data["json"] = nil
		this.ServeJSON()
		return
	}
	this.Data["json"] = us
	this.ServeJSON()
	return
}