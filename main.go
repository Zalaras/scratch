package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
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
		Header: "Default",
		Body:   "Default",
	}
}

func UpdateMessage(msg *Message, body string) {
	msg.Body = body
}

func UpdateMsgHeader(msg *Message, header string) {
	msg.Header = header
}
