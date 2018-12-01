package actions

import (
	"fmt"
)

type hearStruct struct {
	text  string
	regex string
	match string
}

//we want to listen to a message that contains the ppts in the struct field
func (h hearStruct) Hear() string {
	if h.regex == "" {
		fmt.Println("Ignoring regex since it isnt present")
	}

	if h.text == "" {
		fmt.Println("Ignoring text cos it isnt passed")
	}

	if h.text == "" && h.regex == "" {
		fmt.Println("Error listening. . .  regex or text not passed")
	}

	if h.text != "" {
		return h.text
	}

	if h.regex != "" {
		return h.regex
	}

	return ""
}
