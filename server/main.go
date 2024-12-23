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

	//open the files in the directory data
	Userfile, err := os.Open("../data/users.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	Membersfile, err := os.OpenFile("../data/members.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	// Serve the login form
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../templates/login.html")
	})

	//serve the registration page
	// Serve the login form
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../templates/registration.html")
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
		scanner := bufio.NewScanner(Userfile)
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
			fmt.Println("Authenticated")
			// Redirect to the dashboard on successful login
			http.Redirect(w, r, "/dashboard", http.StatusMovedPermanently)
		} else {
			fmt.Println(username)

			// If authentication fails
			fmt.Fprintf(w, "Invalid username or password. Try loggin in.\n")
		}
	})

	//Handle registration
	http.HandleFunc("/registration", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			// Serve the registration page
			http.ServeFile(w, r, "../templates/registration.html")
			return
		}

		if r.Method == http.MethodPost {
			// Parse form data
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Failed to parse form data", http.StatusBadRequest)
				return
			}

			username := r.FormValue("username")
			email := r.FormValue("email")
			password := r.FormValue("password")

			// Validate form inputs
			if username == "" || email == "" || password == "" {
				http.Error(w, "All fields are required", http.StatusBadRequest)
				return
			}

			// Check if the username already exists
			if CredentialExists(username) {
				http.Error(w, fmt.Sprintf("Username %s is already taken. Please choose another.", username), http.StatusConflict)
				return
			}

			// Add user data to the members file
			userData := fmt.Sprintf("u-%s e-%s p-%s\n", username, email, password)
			fmt.Println(userData)

			_, err := Membersfile.WriteString(userData)
			if err != nil {
				http.Error(w, "Failed to save user data", http.StatusInternalServerError)
				return
			}

			// Add user to the system
			if AddUser(username, password) {
				
				// Redirect to the home page
				http.Redirect(w, r, "/", http.StatusSeeOther)
			} else {
				http.Error(w, "Failed to add user to the system", http.StatusInternalServerError)
			}
			return
		}

		// Handle unsupported methods
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})

	fmt.Println("Server is running on localhost:9999")
	if err := http.ListenAndServe(":9999", nil); err != nil {
		log.Fatal(err)
	}

	defer Userfile.Close()
	defer Membersfile.Close()
}

func CredentialExists(credential string) bool {
	file, err := os.Open("../data/users.txt")
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		u := strings.Split(line, " ")[0][2:]
		if credential == u {
			return true
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}
	return false
}

func AddUser(username, password string) bool {
	file, err := os.OpenFile("../data/users.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("u-%s p-%s\n", username, password))
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
