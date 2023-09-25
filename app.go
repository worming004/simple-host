package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusBadRequest)
  fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
  log.Println("Starting server")
  http.HandleFunc("/", greet)
  http.ListenAndServe(":8085", nil)
}
