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
				continue
			}
			us = parts[0][2:]
			psw = parts[1][2:]

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

			// If authentication fails, show an error message
			fmt.Fprintf(w, "Invalid username or password.\n")
		}
	})

	//Handle registration
	http.HandleFunc("/registration", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../templates/registration.html")
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")
		if credentialExists(username) {
			fmt.Println("User Exists!")
		} else {
			file, err := os.Open("../data/members.txt")
			if err != nil {
				fmt.Println(err)
			}
			defer file.Close()
			userData := fmt.Sprintf("u-%s e-%s p-%s\n", username, email, password)
			_, err = file.WriteString(userData)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("A user was added")
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
		}
	})

	fmt.Println("Server is running on localhost:9999")
	if err := http.ListenAndServe(":9999", nil); err != nil {
		log.Fatal(err)
	}
}

func credentialExists(credential string) bool {
	file, err := os.Open("../data/users.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		u := strings.Split(line, " ")[0][2:]
		//			p := strings.Split(line, " ")[1][2:]
		if credential == u {
			return true
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}
	return false
}
