package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
  mux := mux.NewRouter()

  // Load all files in the ./static folder to /static
  fs := http.FileServer(http.Dir("./static"))
  mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

  // // Load the /static/index.html to the / route
  mux.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./static/index.html") 
  })
  http.ListenAndServe(":8000", mux)
}
