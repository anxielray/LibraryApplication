package progress

import "net/http"

func HandleProgress(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/progress.html")
}