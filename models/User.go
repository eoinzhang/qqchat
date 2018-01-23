package models

import (
	_ "gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	URL = "localhost:27017"
)
type User struct {
	Username     string `form:"username"`
	UserPassword string `form:userpassword`
}



type Men struct {
	Users []User
}


//*查询多条数据*
var userAll Men //存放结果

//用户注册
func UserreRister(username, password string) {
	session, err1 := mgo.Dial(URL)
	if err1 != nil {
		panic(err1)
	}
	defer session.Close()//这里代码重复太多，需要解决，主要是session的打开与关闭
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("mychat")
	collection := db.C("user")
	user := &User{Username: username, UserPassword: password}
	err := collection.Insert(user)
	if err !=nil{
		panic(err)
	}


}

//用户登录
func UserLogin(username string)  (User,error){
	session, err1 := mgo.Dial(URL)
	if err1 != nil {
		panic(err1)
	}
	defer session.Close()//这里代码重复太多，需要解决，主要是session的打开与关闭
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("mychat")
	collection := db.C("user")
	u := User{}
	err:=collection.Find(bson.M{"username":username}).One(&u)
	if err !=nil{
		return u ,err
	}
	return u,nil
}
