package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	// Serve the login form
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../templates/login.html")
	})

	// Handle login submissions
	http.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}

		// Get username and password
		username := r.FormValue("username")
		password := r.FormValue("password")

		//Validate the credentials
		file, err := os.Open("../data/users.txt")
		if err != nil {
			log.Fatalf("failed to open file: %s", err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			us := strings.Split(line, " ")[0][2:]
			psw := strings.Split(line, " ")[1][2:]
			if us == username && psw == password {
				// Redirect to the new path
        http.Redirect(w, r, "/", http.StatusMovedPermanently)
    })
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatalf("error reading file: %s", err)
		}
		fmt.Fprintf(w, "Logged in as: %s\n", username)
		// You can add authentication logic here
	})

	// Start the server
	fmt.Println("Starting server on :9090")
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal(err)
	}
}
