package util

import "fmt"

//HandleErr simply logs out/prints error and why it happened
func HandleErr(err error, reason ...string) {
	if err != nil {
		fmt.Printf("Error %s happened. This is probably due to: %s", err, reason)
		panic(err)
	}
}
