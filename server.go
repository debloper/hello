package main

import (
  "fmt"
  "io"
  "net/http"
)

const Port = 8000

func hello(w http.ResponseWriter, r *http.Request) {
  io.WriteString(w, "Hello world!")
}

func main() {
  http.HandleFunc("/", hello)
  fmt.Println("Server is listing on port ", Port)
  http.ListenAndServe(fmt.Sprintf(":%d", Port), nil)
}
