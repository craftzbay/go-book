# Э-мэйл илгээх

Э-мэйл үйлчилгээний SMTP, POP3, IMAP зэрэг протоколууд бий. Эдгээрээс SMTP  протоколын тухай үзэе.

Go хэлний `net/smtp` пакет нь SMTP протоколоор э-мэйл илгээх, хүлээн авах, задлах функцүүдийг агуулна. Энэ сангийн Client төрөл нь SMTP сервертэй холбогдон ажилладаг.

Э-мэйл илгээхийн тулд дараах алхамуудыг гүйцэтгэнэ:

* `Dial()` функцийг дуудаж SMTP сервертэй харилцах шинэ `Client` обект үүсгэнэ
* Э-мэйлийн илгээгч, хүлээн авагч хаягуудыг тохируулна
* `Data()` функцийг дуудаж э-мэйлийн бие хэсгийг үүсгэнэ.

```go
package main

import (
  "bytes"
  "log"
  "net/smtp"
)

func main() {
    // SMTP сервер рүү холбогдох
    client, err := smtp.Dial("mail.example.com:25")
    if err != nil {
        log.Fatal(err)
    }
    // илгээгч, хүлээн авагчийг тохируулах
    client.Mail("sender@example.org")
    client.Rcpt("recipient@example.net")

    // э-мэйлийн бие хэсгийг үүсгэх
    wc, err := client.Data()
    if err != nil {
        log.Fatal(err)
    }
    defer wc.Close()

    buf := bytes.NewBufferString("Э-мэйлийн бие.")
    if _, err = buf.WriteTo(wc); err != nil {
        log.Fatal(err)
    }
}
```

Хэрэв тухайн SMTP серверээр э-мэйл илгээх үед нэвтрэх нэр, нууц үг шаарддаг бол `smtp.SendMail()` функцийг ашиглах хэрэгтэй. Энэ функц нь хэд хэдэн үйлдлийг зэрэг гүйцэтгэнэ: эхлээд заасан хаяг дээрх сервер уруу холбогдоно, хэрэв шаардлагатай бол TLS холболт үүсгэнэ, дараа нь хүлээн авагч уруу э-мэйлийг илгээнэ.

Дараах жишээ програмд `SendMail()` функцийг хэрхэн ашиглахыг харуулсан байна.

```go
package main

import (
  "log"
  "smtp"
)

func main() {
    // нэвтрэх эрхийг тохируулах
    auth := smtp.PlainAuth(
        "",
        "user@example.com",
        "password",
        "mail.example.com",
    )

    // э-мэйл илгээх
    err := smtp.SendMail(
        "mail.example.com:25",
        auth,
        "sender@example.org",
        []string{"recipient@example.net"},
        []byte("Э-мэйлийн бие."),
    )

    if err != nil {
        log.Fatal(err)
    }
}
```



