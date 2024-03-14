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

type NewFeat struct {
	New  string
	Feat string
}

func NewMessage() Message {
	return Message{
		Header: "Temp23",
		Body:   "Temp23",
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
