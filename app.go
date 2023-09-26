package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusBadRequest)
  t := time.Now()
  log.Printf("request received at %s\n", t.String())
  fmt.Fprintf(w, "Hello World! %s", t)
}

func main() {
  log.Println("Starting server")
  http.HandleFunc("/", greet)
  http.ListenAndServe(":8085", nil)
}
