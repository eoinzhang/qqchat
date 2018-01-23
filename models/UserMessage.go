package models

import (
	"gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
	_"fmt"
	"gopkg.in/mgo.v2/bson"
)

type UserMessage struct {
	Homeusername     string `form:homeusername`
	Homeuserpassword string `form:homeuserpassword`
	Homeusercall     string `form:homeusercall`
	Homeuserage      string `form:homeuserage`
	Homeusersex      string `form:Homeusersex`
	Homeuserinterest string `from:Homeuserinterest`
}
type Message struct {
	userMessage []UserMessage
}
//信息初始化
func Ini(username, password string)  {
	session, err1 := mgo.Dial(URL)
	if err1 != nil {
		panic(err1)
	}
	defer session.Close()//这里代码重复太多，需要解决，主要是session的打开与关闭
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("mychat")
	collection := db.C("usermessage")
	usermessage:=&UserMessage{Homeusername:username,Homeuserpassword:password}
	err := collection.Insert(usermessage)
	if err !=nil{
		panic(err)
	}
}
//用户主页信息保存
func MessageSave(usermessage UserMessage)  {
	session, err1 := mgo.Dial(URL)
	if err1 != nil {
		panic(err1)
	}
	defer session.Close()//这里代码重复太多，需要解决，主要是session的打开与关闭
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("mychat")
	collection := db.C("usermessage")
	user := &UserMessage{
		Homeusername:usermessage.Homeusername,
		Homeuserpassword:usermessage.Homeuserpassword,
		Homeusercall:usermessage.Homeusercall,
		Homeuserage:usermessage.Homeuserage,
		Homeusersex:usermessage.Homeusersex,
		Homeuserinterest:usermessage.Homeuserinterest,
	}
	err := collection.Insert(user)
	if err !=nil{
		panic(err)
	}
}
//用户主页信息查看
func MessageLook(homeusername string) UserMessage {
	session, err1 := mgo.Dial(URL)
	if err1 != nil {
		panic(err1)
	}
	defer session.Close()//这里代码重复太多，需要解决，主要是session的打开与关闭
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("mychat")
	collection := db.C("usermessage")
	u := UserMessage{}
	err:=collection.Find(bson.M{"homeusername":homeusername}).One(&u)
	if err !=nil{
		panic(err)
	}
	return u
}