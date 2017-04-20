package main

import (
  "fmt"
  "io"
  "net/http"
  "strings"
)

const Port = 8000

func hello(w http.ResponseWriter, r *http.Request) {
  text := r.URL.Path[1:]
  text = strings.Title(text)

  if text == "" {
    text = "World"
  }
  io.WriteString(w, fmt.Sprintf("Hello, %s", text))
}

func main() {
  http.HandleFunc("/", hello)
  fmt.Println("Server is listing on port ", Port)
  http.ListenAndServe(fmt.Sprintf(":%d", Port), nil)
}
