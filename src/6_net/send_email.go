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

	// э-мэйлийн бие хэсэг үүсгэх
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
