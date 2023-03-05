package main

import (
	"fmt"
	"net/http"

	"github.com/4IDTest/SingOut/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Handle("/logout", handlers.LogoutHandler()).Methods(http.MethodPost)
	port := "8081"
	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	fmt.Println("Staring Login test server on Port " + port)
	// Start Server on defined port/host.
	server.ListenAndServe()
}
