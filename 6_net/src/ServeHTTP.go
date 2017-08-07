package main

import (
  "fmt"
  "net"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Hello UB!")
}

func main() {
  var h Hello
  http.ListenAndServe("localhost:4000",h)
}
