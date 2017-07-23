#Жишээ дасгал

1. Өгөгдсөн хавтас дотор *.html файлуудыг самнаж доторхи холбоосуудыг ( a таагийн href атрибутад заасан URL) хэвлэж харуулах програм бичээрэй. Хавтас самнах үйлдлийг хурдасгахын тулд файл унших үйлдлийг зэрэгцээ гүйцэтгээрэй.

  **Шийдэл**. Програмыг хавтас доторх html файлуудыг самнаж бүртгэх, олсон html файл дотор боловсруулалт хийх гэсэн хоёр дэд хэсэгт хувааж болно.

  Дараах нэртэй функцүүд үүсгэе.
  `crawl` – хавтас, дэд хавтасаар дамжин самнаж html файл хайна
  `findLink` – html файл дотор холбоос хайна

  Дээрх хоёр функцийг зэрэг ажиллуулвал илүү үр дүнтэй байна. Зэрэг ажиллах үедээ хоорондоо мэдээлэл солилцоход нь зориулж `taskChannel` нэртэй суваг үүсгэж болно. Энэ сувгаар боловсруулах шаардлагатай файлын нэрсийг солилцоно.

  Програмыг бүр илүү үр дүнтэй ажиллуулахын тулд файл дотроос холбоос хайх (`findLink`) функцийг 2 буюу түүнээс дээш тоогоор зэрэг ажиллуулж болно.

  Дараах жишээнд 3 зэрэг ажиллуулсан байна.

  ```go
  func main() {
  	// эхлэх хавтас
  	rootFolder := "web_root"

  	// ажлын даалгавар (файлын нэрс) солилцох суваг үүсгэх
  	taskChannel := make(chan string)

  	// самнаж эхлэх (go функц үүсгэж байна)
  	go crawl(rootFolder, taskChannel)

  	// 3 зэрэг findLink функц ажиллуулах
  	for i := 0; i < 3; i++ {
  		go findLink(taskChannel)
  	}

  	// самналт дуусахыг хүлээх
  	<- done
  }
  ```

  Програмын үндсэн шийдэл нь ингээд боллоо. `Crawl` болон `findLink` функцүүдийн бүтэн кодыг доор харуулав.

  ```go
  // хавтасаар самнах
  func crawl(folder string, ch chan string) {
  	files, _ := ioutil.ReadDir(folder)

  	for _, f := range files {

  		if f.IsDir() {
  			// хавтас бол цааш самнах
  			crawl(folder+"/"+f.Name(), ch)
  		} else {
  			//  *.html файл эсэхийг шалгах
  			if strings.HasSuffix(f.Name(), fileExt) {
  				// мөн бол ажлын дараалалд оруулах
  				ch <- folder + "/" + f.Name()
  			}

  		}

  	}
  	done <- true
  }

  // html файл дотроос холбоос хайх функц
  func findLink(ch chan string) {
  	for {
  		select {
  		case fileName := <-ch:
  			fmt.Println(fileName)
  			// TODO: энд холбоос хайх хэсэг байна
  		}
  	}
  }
  ```

2. Go функц ашиглан дэлгэц дээр 7 багананд санамсаргүй тэмдэгтүүдийг санамсаргүй дарааллаар урсган харуулах програм бичээрэй.

  **Шийдэл**: Дэлгэц дээр текст хэвлэхэд курсор байрлаж байгаа байрлалд хэвлэгддэг. Курсорын байрлалыг `\033` тусгай тэмдэгтээр удирдаж болно. Энэ боломжийг ашиглаад `printAt()` функц бичсэн. Энэ функц нь заагдсан байрлалд тэмдэгтийг ногоон өнгөөр хэвлэнэ.

  `matrix()` функц нь санамсаргүй тэмдэгт үүсгэж суваг уруу тасралтгүй илгээнэ, `printer()` функц тэмдэгтүүдийг хүлээж аваад `printAt()` функцийн тусламжтайгаар багануудад хувааж хэвлэнэ.

  ```go
  package main

  import (
    "strconv"
    "fmt"
    "unicode"
    "time"
    "math/rand"
  )

  func printAt(s string, x int, y int) {
   // курсор байрлуулах
   print("\033["+strconv.Itoa(y)+";"+strconv.Itoa(x)+"H")
   print("\033[92m") // ногоон өнгөөр хэвлэх
   print(s+"\r") 	// тэмдэгтийг хэвлэх
  }

  func clearScreen() {
   print("\033[2J")
  }

  func matrix(c chan<- string) {
   for {
    // санамсаргүй тэмдэгт зохиож илгээх
    r:=rune(rand.Intn(65536))
    if unicode.IsPrint(r) {
      c <- strconv.QuoteRune(r)
    }
   }
  }

  func printer(c <-chan string) {
    column:=1
    i:=1
    for {
      msg := <- c
      printAt(msg, column, i)

      // багануудад хуваах
      if column < 80 {
        column=column+10
      } else {
        column = 1
        i=i+1
        println()
      }

      time.Sleep(time.Millisecond * 10)
    }
  }

  func main() {
   var c chan string = make(chan string)

   clearScreen() 	   // дэлгэц цэвэрлэх

   go matrix(c)
   go printer(c)

   var input string
   fmt.Scanln(&input)
  }
  ```
