package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello world"))
	})

	fmt.Println(fmt.Sprintf("Starting application on port %s - http://localhost%s/", portNumber, portNumber))

	if err := http.ListenAndServe(portNumber, nil); err != nil {
		panic(err)
	}
}
