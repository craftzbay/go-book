# TLS

Интернэтэд мэдээллийг энкриптлэж явуулах хамгийн түгээмэл механизмT бол TLS \(Transport Layer Security\) юм. Өмнө нь үүнийг SSL \(Secure Sockets Layer\) гэж нэрлэдэг байсан.

TLS-д клиент болон сервер хоорондоо X.509 сертификат солилцож бие биенээ танина. Үүний дараа хувийн түлхүүр үүсгэнэ, ингээд энэ түлхүүрээр бүх энкрипт/декрипт явагдана. Эхний танилцах алхам нь бага зэрэг удаан, гэхдээ үүнээс хойших дамжуулалт хурдан байдаг.

Сервер нь:

```go
// TLSEchoServer.go
package main

import (
        "crypto/rand"
        "crypto/tls"
        "fmt"
        "net"
        "os"
        "time"
)

func main() {

        cert, err := tls.LoadX509KeyPair("my.server.pem", "private.pem")
        checkError(err)
        config := tls.Config{Certificates: []tls.Certificate{cert}}

        now := time.Now()
        config.Time = func() time.Time { return now }
        config.Rand = rand.Reader

        service := "0.0.0.0:1200"

        listener, err := tls.Listen("tcp", service, &config)
        checkError(err)
        fmt.Println("Listening")
        for {
                conn, err := listener.Accept()
                if err != nil {
                        fmt.Println(err.Error())
                        continue
                }
                fmt.Println("Accepted")
                go handleClient(conn)
        }
}

func handleClient(conn net.Conn) {
        defer conn.Close()

        var buf [512]byte
        for {
                fmt.Println("Trying to read")
                n, err := conn.Read(buf[0:])
                if err != nil {
                        fmt.Println(err)
                }
                _, err2 := conn.Write(buf[0:n])
                if err2 != nil {
                        return
                }
        }
}

// func checkError(err error)
```

Энэ сервер нь дараах клиенттэй ажиллана:

```go
// TLSEchoClient.go
package main

import (
        "fmt"
        "os"
        "crypto/tls"
)

func main() {
        if len(os.Args) != 2 {
                fmt.Println("Usage: ", os.Args[0], "host:port")
                os.Exit(1)
        }
        service := os.Args[1]

        conn, err := tls.Dial("tcp", service, nil)
        checkError(err)

        for n := 0; n < 10; n++ {
                fmt.Println("Writing...")
                conn.Write([]byte("Hello " + string(n+48)))

                var buf [512]byte
                n, err := conn.Read(buf[0:])
                checkError(err)

                fmt.Println(string(buf[0:n]))
        }
        os.Exit(0)
}

//func checkError(err error)
```



