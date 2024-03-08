package main

import "fmt"

func main() {
	fmt.Println("Hello Universe!")
	newFunc()
}

func newFunc() {
	fmt.Println("new print somethingelseyouknow!")
}

type Message struct {
	Header string
	Body   string
	Meta   string
}

func NewMessage() Message {
	return Message{
		Header: "Temp",
		Body:   "Temp",
	}
}

func UpdateMessage(msg *Message, body string) {
	msg.Body = body
}

func UpdateMsgHeader(msg *Message, header string) {
	msg.Header = header
}
