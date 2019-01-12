package main

import (
	"flag"
	"fmt"
)

//Bot struct is the struct our bot property will be made after
//BRAINSTORM: SHouldnt it be an interface? 🤔
type Bot struct {
	name    string
	channel string
}

// func init() {
// 	logger.CreateLogFolder("error")
// 	logger.CreateLogFile("error", "error.log")
// }

// declare our CLI flags
var (
	help = flag.String("help", "", "Display the help info about the framework")
)

func main() {
	flag.Parse()

	fmt.Printf("User inputs: %s", *help)
}
