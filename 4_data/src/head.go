package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Файлын нэр заана уу\n")
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Файлыг нээхэд алдаа гарлаа: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	n := 10
	scanner := bufio.NewScanner(f)
	for n > 0 && scanner.Scan() {
		fmt.Println(scanner.Text())
		n--
	}
}
