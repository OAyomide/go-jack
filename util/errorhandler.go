package util

import "fmt"

//HandleErr simply logs our error and the reason it happened
func HandleErr(err error, reason ...string) {
	if err != nil {
		fmt.Printf("Error %s: Most likely happened due to %s\n", err, reason)
		panic(err)
	}
}
