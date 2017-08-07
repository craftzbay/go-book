#  Жишээ дасгал

1. `os.Stat()` функц нь файлын тухай мэдээлэл буцаана. Энэ функцийг ашиглан файлын хэмжээг олох програм бичээрэй.

  ```go
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
  ```

2. Өгөгдсөн текст дотор орсон э-мэйл хаягуудыг ялгаж харуулах програм бичээрэй. Хялбар э-мэйл хаягийг `\w+@\w+\.\w+` хэвээр хайж болно.

  ```go
  package main

  import (
    "fmt"
    "regexp"
  )

  func main() {
  	txt := "Та ubs121@gmail.com, ub121@hotmail.com хаягуудаар холбоо барьж болно."

  	re, _ := regexp.Compile(`\w+@\w+\.\w+`)
  	all := re.FindAllString(txt, -1)

  	for _, m := range all {
  		fmt.Printf("%s\n", m)
  	}
  }
  ```

3. Файлаас эхний 10 мөрийг таслан дэлгэцэнд харуулах `Head` нэртэй програм бичээрэй.

  ```go
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
  ```
