package main
import ("encoding/xml";"fmt";"os")

type Person struct {
  XMLName Name `xml:"person"`
  Name Name `xml:"name"`
  Email []Email `xml:"email"`
}
type Name struct {
  First string `xml:"first"`
  Last string `xml:"last"`
}
type Email struct {
  Type string `xml:"type,attr"`
  Address string `xml:",chardata"`
}

func main() {
  str := `<?xml version="1.0" encoding="utf-8"?>
	<person>
	<name>
	<first>Ууганбаяр</first>
	<last>Сүхбаатар</last>
	</name>
	<email type="хувийн">ubs121@gmail.com</email>
	<email type="ажлын">ub@hotmail.com</email>
	</person>`

  var person Person
  err := xml.Unmarshal([]byte(str), &person)
  checkError(err)

  // обектыг ашиглах
  fmt.Println("Нэр: \"" + person.Name.First + "\"")
  fmt.Println("Э-мэйл 2: \"" + person.Email[1].Address + "\"")
}

func checkError(err error) {
  if err != nil {
    fmt.Println("Fatal error ", err.Error())
    os.Exit(1)
  }
}