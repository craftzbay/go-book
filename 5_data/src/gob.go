package main

import ("encoding/gob"; "fmt"; "os" )

type Person struct {
 Name  Name
 Email []Email
}
type Name struct {
  First string
  Last string
}
type Email struct {
  Kind string
  Address string
}

func main() {
 person := Person{
  Name: Name{First: "Ууганбаяр", Last: "Сүхбаатар"},
  Email: []Email{Email{Kind: "хувийн", Address: "ubs121@gmail.com"},
         Email{Kind: "ажлын", Address: "ub@hotmail.com"}}}

  saveGob("person.gob", person)
}

func saveGob(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	checkError(err)
	encoder := gob.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
	outFile.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}