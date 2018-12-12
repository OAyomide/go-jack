package main

import (
	"go-jack/actions"
	"go-jack/logger"
)

//Bot struct is the struct our bot property will be made after
//BRAINSTORM: SHouldnt it be an interface? 🤔
type Bot struct {
	name    string
	channel string
}

func main() {
	logger.Print("warning", "Hello, 🦄", " ")
	actions.Listen("Hello")
}
