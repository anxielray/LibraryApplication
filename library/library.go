package library

import "net/http"

func HandleLibrary(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/my_books.html")
}
