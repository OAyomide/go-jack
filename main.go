package main

import (
	"go-jack/logger"
)

//Bot struct is the struct our bot property will be made after
//BRAINSTORM: SHouldnt it be an interface? 🤔
type Bot struct {
	name    string
	channel string
}

func init() {
	logger.CreateLogFolder("Error")
	logger.CreateLogFile("Error", "error.critical.log")
}
func main() {
	//for i := 0; i <= 5; i++ {
	logger.Print("info", "HELLO THERE!!")
	//}
}
