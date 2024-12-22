package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

func main() {
	// Serve the login form
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../templates/login.html")
	})

	// Handle the dashboard route
	http.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("../templates/index.html")
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

	// Handle login submissions
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}

		// Get username and password
		username := r.FormValue("username")
		password := r.FormValue("password")
		var us, psw string

		// Validate the credentials
		file, err := os.Open("../data/users.txt")
		if err != nil {
			log.Fatalf("failed to open file: %s", err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		authenticated := false

		for scanner.Scan() {
			line := scanner.Text()
			parts := strings.Split(line, " ")
			if len(parts) < 2 {
				continue // Skip lines that don't have enough parts
			}
			us = parts[0][2:]  // Assuming username starts after two characters
			psw = parts[1][2:] // Assuming password starts after two characters

			if us == username && psw == password {
				authenticated = true
				break
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatalf("error reading file: %s", err)
		}

		if authenticated {
			// Redirect to the dashboard on successful login
			http.Redirect(w, r, "/dashboard", http.StatusMovedPermanently)
		} else {
			fmt.Printf("%s is correct but is != %s\n", us, username)
			fmt.Printf("%s is correct but is != %s\n", psw, password)
			// If authentication fails, show an error message
			fmt.Fprintf(w, "Invalid username or password.\n")
		}
	})

	fmt.Println("Server is running on localhost:9999")
	if err := http.ListenAndServe(":9999", nil); err != nil {
		log.Fatal(err)
	}
}
