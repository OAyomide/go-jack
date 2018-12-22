package main

import (
	"fmt"

	"github.com/go-jack/parser"
)

type Bot struct {
	name string
}

var conf parser.Config

func main() {
	fmt.Println("Hello there! Shiny new framework")
	ps := conf.ReadYml()
	token := ps.AccessToken

	fmt.Println("Here is the access token. . .should fail though", token)
}
