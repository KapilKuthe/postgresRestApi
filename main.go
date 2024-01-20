// main.go
package main

import (
	"fmt"
	"net/http"
	"postgresRestApi/database"
	"postgresRestApi/handler"

	"github.com/gorilla/mux"
)

func main() {
	database.InitializeDB()  // Initialize database connection
	defer database.CloseDB() // Close the database connection when the application exits

	// Initialize the JWT secret key
	handler.InitJWTKey()

	router := mux.NewRouter()

	router.HandleFunc("/users", handler.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", handler.GetUser).Methods("GET")
	router.HandleFunc("/users", handler.CreateUser).Methods("POST")

	// Apply AuthMiddleware only to PUT and DELETE requests to paths starting with "/secure"
	secureRouter := router.PathPrefix("/secure").Subrouter()
	secureRouter.Use(handler.AuthMiddleware)
	secureRouter.HandleFunc("/users/{id}", handler.UpdateUser).Methods(http.MethodPut)
	secureRouter.HandleFunc("/users/{id}", handler.DeleteUser).Methods(http.MethodDelete)

	port := 8080
	serverAddress := fmt.Sprintf(":%d", port)
	fmt.Printf("Server is running on %s\n", serverAddress)
	http.ListenAndServe(serverAddress, router)
}
