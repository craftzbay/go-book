package main

import (
	"log"
	"net/smtp"
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
