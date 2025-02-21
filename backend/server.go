package main

import (
	"demo/db/sqlite"
	"demo/pkg/database"
	"demo/pkg/handlers"
	"log"
	"net/http"

	go_handlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	db := database.GetDB()
	defer database.CloseDB()

	err := sqlite.ApplyMigrations(db, "./db/migrations/sqlite")
	if err != nil {
		log.Println("Error applying migrations:", err)
		return
	}

	r := mux.NewRouter()

	r.HandleFunc("/register", handlers.RegisterUser).Methods("POST")

	corsOptions := go_handlers.CORS(
		go_handlers.AllowedOrigins([]string{"http://localhost:5173"}),         // Allow Vue.js frontend
		go_handlers.AllowedMethods([]string{"GET", "POST", "DELETE"}),         // Allowed requests methods
		go_handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}), //Allowed headers
		go_handlers.AllowCredentials(),                                        // This is MANDATORY for cookies/sessions
	)
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("../frontend/src/components"))))

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", corsOptions(r))
}
