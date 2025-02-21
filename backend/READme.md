# *if you are using VS code... to view this file in a proper way press: ctrl+shift+V*

Hello

in this directory you will find the backend of the server

---
this is the backend directory
```
backend
├───db
│   ├───migrations
│   │   └───sqlite
│   └───sqlite
└───pkg
    ├───database
    ├───handlers
    └───structs
```
---

### /db/migrations/sqlite

contains migration files for the database {create users table}

---

### /db/sqlite

contains a `sqlite.go` file that contains the code that applies migrations on the database

---

### /pkg

contains directories that are packages in our source code. check them out!

---

### /server.go

contains the code that will run the whole backend.

this will initialize the databse then apply migrations on it.

```go
	db := database.GetDB()
	defer database.CloseDB()

	err := sqlite.ApplyMigrations(db, "./db/migrations/sqlite")
	if err != nil {
		log.Println("Error applying migrations:", err)
		return
	}
```

---

this snippet will allow the frontend port `http://localhost:5173` to send requests to our backend port  `http://localhost:8080`

```go
	corsOptions := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:5173"}),         // Allow Vue.js frontend
		handlers.AllowedMethods([]string{"GET", "POST", "DELETE"}),         // Allowed request methods
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}), //Allowed headers
		handlers.AllowCredentials(),                                        // This is MANDATORY for cookies/sessions
	)
```

and it's important to run the backend with the allowed cors 

`http.ListenAndServe(":8080", corsOptions(r)) `

# PLEASE reach out for further explanation