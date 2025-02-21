package handlers

import (
	"demo/pkg/database"
	"demo/pkg/structs"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed. Use POST.", http.StatusMethodNotAllowed)
		return
	}

	var user structs.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println("❌ Error decoding request body:", err)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	// Get database connection
	db := database.GetDB()

	_, err := db.Exec(`
		INSERT INTO users (username, first_name, last_name, age, gender) 
		VALUES (?, ?, ?, ?, ?)`,
		user.Username, user.FirstName, user.LastName, user.Age, user.Gender)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			http.Error(w, "email or nickname already in use" , http.StatusBadRequest)
		}
		log.Printf("❌ Failed to insert user into database: %v", err)
		http.Error(w , "error inserting user" , http.StatusInternalServerError)
	}
	
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}
