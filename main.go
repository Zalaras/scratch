package main

import "fmt"

func main() {
	fmt.Println("2 World!")
	newFunc()
}

func newFunc() {
	fmt.Println("new test22!")
}

type Message struct {
	Header string `json:"header"`
	Body   string `json:"body"`
	Meta   Meta   `json:"meta"`
}

type Meta struct {
	Data interface{}
}

type NewFeat struct {
	ID   int    `json:"id"`
	New  string `json:"new"`
	Feat string `json:"feat"`
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
