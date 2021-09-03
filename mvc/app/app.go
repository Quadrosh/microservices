package app

import (
	"fmt"
	"net/http"

	"github.com/quadrosh/microservices/mvc/controllers"
)

const portNumber = ":8080"

// StartApp ...
func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	fmt.Println(fmt.Sprintf("Starting application on port %s - http://localhost%s/", portNumber, portNumber))

	if err := http.ListenAndServe(portNumber, nil); err != nil {
		panic(err)
	}
}
