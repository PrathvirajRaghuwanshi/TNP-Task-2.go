package main

import (
	"log"
	"net/http"

	"TnpTask2/database"
	"TnpTask2/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize Database
	database.ConnectDB()

	// Initialize router
	r := mux.NewRouter()

	// Middleware
	r.Use(handlers.AuthMiddleware)

	// Routes
	r.HandleFunc("/certificates", handlers.GetCertificates).Methods("GET")
	r.HandleFunc("/certificates/{id}/send", handlers.SendCertificate).Methods("POST")
	r.HandleFunc("/bulk-messages", handlers.SendBulkMessages).Methods("POST")

	log.Println("🚀 Server started on :8080")
	http.ListenAndServe(":8080", r)
}
