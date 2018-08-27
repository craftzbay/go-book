package main

import (
 "os"
 "fmt"
)

func main() {
  fi, err := os.Stat("file_size.go")
  if err != nil {
    // файлын мэдээллийг авч чадсангүй
    os.Exit(1)
  }

  fmt.Printf("%s файлын хэмжээ %d байт\n", fi.Name(), fi.Size())
}