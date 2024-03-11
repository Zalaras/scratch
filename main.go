package main

import "fmt"

func main() {
	fmt.Println("2 World!")
	newFunc()
}

func newFunc() {
	fmt.Println("new test!")
}

type Message struct {
	Header string
	Body   string
	Meta   Meta
}

type Meta struct {
	Data interface{}
}

func NewMessage() Message {
	return Message{
		Header: "Temp2",
		Body:   "Temp2",
	}
}

func UpdateMessage(msg *Message, body string) {
	msg.Body = body
}

func UpdateMsgHeader(msg *Message, header string) {
	msg.Header = header
}

func AddMsgMeta(msg *Message, meta Meta) {
	msg.Meta = meta
}
