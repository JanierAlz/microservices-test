package main

import (
	"fmt"
	"net/http"

	"github.com/4IDTest/SingIn/handlers"
	"github.com/4IDTest/SingIn/history"
	"github.com/4IDTest/SingIn/initializers"
	"github.com/4IDTest/SingIn/user"
	"github.com/gorilla/mux"
)

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		fmt.Println("failed to load config")
	} else {
		initializers.ConnectDB(&config)
	}
	router := mux.NewRouter()
	router.Handle("/login", handlers.AuthHandler()).Methods(http.MethodPost)
	router.Handle("/updateLog", handlers.UpdateLogRecordHandler()).Methods(http.MethodPost)
	router.Handle("/register", handlers.RegisterUserHandler()).Methods(http.MethodPost)
	port := "8080"
	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	fmt.Println("Staring Login test server on Port " + port)
	initializers.DB.AutoMigrate(&user.User{}, &history.History{})
	// Start Server on defined port/host.
	server.ListenAndServe()
}
