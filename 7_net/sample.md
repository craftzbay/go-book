# Жишээ дасгал

1. TCP/IP протокол ашиглан сүлжээгээр харилцах клиент болон сервер програмууд бичээрэй.

  **Шийдэл**: Сервер талд `8088` порт дээр ажиллах програм бичие.  Сервер талд `Listener` обект үүсгэсний дараа `Accept()` методыг дуудаж клиент холбогдохыг хүлээнэ. Клиент холбогдоход түүнд мэдээлэл (жнь: өөрийнх нь IP хаяг) дамжуулаад холболтыг хаана.

  ```go
  // tcp_server.go
  package main

  import (
    "fmt"
    "net"
    "os"
  )

  func main() {
    service := ":8088"
    listener, err := net.Listen("tcp", service)
    checkError(err)

    for {
  	conn, err := listener.Accept()
  	if err != nil {
  	 continue
  	}
  	// клиентэд хариу илгээх
  	conn.Write([]byte("Амжилттай холбогдлоо. Таны хаяг: " +
  		 conn.RemoteAddr().String() + "\n" ))
  	conn.Close()
    }
  }

  //  checkError(err error) функцийг энд оруулна
  ```

  Сервер програмын ажиллагааг хурдан туршиж үзэх бол серверээ асаагаад nc програмаар хүсэлт илгээж болно:

  ```sh
  $ echo -n "Hello" | nc localhost 8088
  Амжилттай холбогдлоо. Таны хаяг: 127.0.0.1:48653
  ```

  Одоо дээрх сервер програмтай холбогдох клиент програмыг бичие.

  ```go
  // tcp_client.go
  package main

  import (
    "fmt"
    "io/ioutil"
    "net"
    "os"
  )

  func main() {
    tcpAddr, err := net.ResolveTCPAddr("tcp4", "localhost:8088")
    checkError(err)

    // холболт тогтоох
    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    checkError(err)

    // серверээс мэдээлэл унших
    result, err := ioutil.ReadAll(conn)
    checkError(err)

    // авсан хариуг дэлгэцэнд хэвлэж харуулах
    fmt.Println(string(result))
  }

  //  checkError() функцийг энд оруулна
  ```

2. Вэбээс сервер рүү файл ачаалах програм бичээрэй.

  **Шийдэл**: Вэбээс ирсэн хүсэлтийг боловсруулахын тулд вэб функц үүсгэх хэрэгтэй. `Upload` нэртэй дараах вэб функц үүсгэе.

  ```go
  func Upload(w http.ResponseWriter, r *http.Request) {
  // 1. илгээсэн файлын агуулгыг унших

  // 2. уншсан өгөгдлөө файл руу бичих
  }
  ```

  a) Илгээсэн файлын агуулгыг http.Request обектын Body талбараас уншиж болно.
  ```go
  body, err = ioutil.ReadAll(r.Body)
  ```
  b) Уншсан өгөгдлөө файл руу бичихэд ioutil.WriteFile() функцийг ашиглаж болно.

  ```go
  err = ioutil.WriteFile("upload.dat", body, 0700)
  ```

  Ингээд `Upload` функцээ http үйлчилгээнд бүртгэх хэрэгтэй.

  ```go
  http.HandleFunc("/", Upload)
  ```
