package models

import (
	_ "gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
	_"fmt"
	"gopkg.in/mgo.v2"
)

type ChatRecord struct {
	Onlineuser string `form:onlineuser`
	Chatrecord string `form:chatrecord`
}
type Chat struct {
	ChatRecords []User
}
//信息初始化
func In(onlineuser string)  {
	session, err1 := mgo.Dial(URL)
	if err1 != nil {
		panic(err1)
	}
	defer session.Close()//这里代码重复太多，需要解决，主要是session的打开与关闭
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("mychat")
	collection := db.C("chatrecord")
	chatrecord:=&ChatRecord{Onlineuser:onlineuser,Chatrecord:""}
	err := collection.Insert(chatrecord)
	if err !=nil{
		panic(err)
	}
}
//用户聊天记录保存
func RecordSave()  {
	
}
//用户聊天记录查看
func RecordLook()  {

}
