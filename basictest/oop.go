package main

import "fmt"

type Messageer interface{
	Show()
}

type Message struct{
	message string
}

func (x Message) Show(){
	fmt.Println(x.message)
}

func Msg(msg string) Message{
	return Message{msg}
}

func main(){
	msg := Msg("hello")
	msg.Show()

	msg2 := Msg("helloworld")
	msg2.Show()
}
