package main

import (
	"fmt"
	//"net/http"
	//"server/routes"
)

var (
	port = "8080"
)

func main() {
	// Application initialization code

	// Set up routes and handlers
	//router := routes.SetupTaskRoutes()

	// Start the HTTP server
	fmt.Printf("Server is running on port %s...\n", port)
	//http.ListenAndServe(":"+port, router)
}
