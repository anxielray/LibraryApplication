package main

import (
	auth "Anxiel-Archives/auth"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {

	// Serve the login form
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/login.html")
	})
	http.HandleFunc("/login", auth.HandleLogin)

	//serve the registration form
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/registration.html")
	})
	http.HandleFunc("/register/", auth.HandleRegistration)

	// Handle the dashboard route
	http.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = t.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	fmt.Println("Server is running on localhost:9999")
	if err := http.ListenAndServe(":9999", nil); err != nil {
		log.Fatal(err)
	}

	//defer close the open files
	U, M := auth.FileData()
	defer U.Close()
	defer M.Close()
}
