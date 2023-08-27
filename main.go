package main

import (
	// "time"
	// "fmt"
	"encoding/json"
	"net/http"
	// "github.com/gorilla/mux"
  // "template/html"
	// "errors"
	// "io"
)

type User struct {
  Name string `json:"name"`
  // Age int `json:"age"`
}

var users []User

func getUsers(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  // w.Header().Set("Content-Type", "text/html")
	json.NewEncoder(w).Encode(users)
}


func main() {

  // users = append(users, User{Name: "malachi"})

  
  mux := http.NewServeMux()
  
  // Staticly load your /static files
  dir := http.Dir("./static")
  fs := http.FileServer(dir)
  mux.Handle("/static/", http.StripPrefix("/static/", fs))

  // Load the /static/index.html to the / route
  mux.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./static/index.html") 
  })

  http.ListenAndServe(":8000", mux)

}
