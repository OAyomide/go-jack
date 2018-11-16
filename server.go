package main

import (
	"fmt"
	"net/http"
	"os"
)

func HomeRoute(i http.ResponseWriter, j *http.Request) {
	fmt.Fprint(i, "HELLO THERE!! THIS IS OUR ROUTE!!")
}

func getPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = "3500"
		fmt.Println("PORT NOT DEFINED FOR THE SERVER. USING THE PORT %s AS THE PORT\n", port)
	}
	return port
}
