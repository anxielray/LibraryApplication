package profile

import "net/http"

func HandleProfile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/profile.html")
}