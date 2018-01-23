package models

import (
	"fmt"
	"time"
)

type UsersMessage struct {
	Uname string
	Content string
}
var UM UsersMessage
var IJ UsersMessage
var i=UsersMessage{}
var IJJ=make([]UsersMessage,0)
var ch1 =make(chan UsersMessage,10)
func Post2(Uname,Content string)  {
	UM=UsersMessage{Uname:Uname,Content:Content}
	fmt.Println(UM)
	Post1(ch1)
}
func Post1(ch chan UsersMessage)  {


		if UM!=i{
			ch <- UM
		}



}
func init() {



	go Fetch1(ch1)

}
func Fetch1(ch chan UsersMessage)  {
	fmt.Println(ch)
	for{
		IJ=<-ch
		fmt.Println(ch)
		IJJ=make([]UsersMessage,0)
  if IJ!=i{
	  IJJ=append(IJJ,IJ)
  }
   time.Sleep(time.Second*1)
	}
}
func Fetch2()  []UsersMessage{

	return IJJ

}