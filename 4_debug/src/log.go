package main

import (
	"log"
	"os"
)

func main() {
	Info := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	Info.Println("Hello from logger !")
}
