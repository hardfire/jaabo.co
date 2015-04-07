package main

import (
	"jaabo/controllers"
	"log"
	"net/http"
	"os"
	"path"
	//"github.com/justinas/alice"
)

func main() {

	// set deployment port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/jobs", controllers.JobsIndexHandler)
	mux.HandleFunc("/jobs/", controllers.JobShowHandler)
	mux.HandleFunc("/auth/handle/google", controllers.GoogleAuthHandler)

	log.Println("Listening on port", port)
	http.ListenAndServe(":"+port, mux)
}

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	fp := path.Join("Public", "index.html")
	http.ServeFile(rw, r, fp)
	//fmt.Fprintln(rw, "Home")
}
