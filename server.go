package main

import (
  "io"
  "fmt"
  "net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
  io.WriteString(w, "Hello world!")
}

func main() {
  http.HandleFunc("/", hello)
  fmt.Println("Server is listing on port 8000.")
  http.ListenAndServe(":8000", nil)
}
