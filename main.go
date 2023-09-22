package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}

func SendEmail() {
	secretPassword := Config("EMAIL_SECRET_KEY")
	auth := smtp.PlainAuth(
		"",
		"agustfricke@gmail.com",
		secretPassword,
		"smtp.gmail.com",
	)

	tmpl, err := template.ParseFiles("email_template.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	data := struct {
		Name string
	}{
		Name: "Usuario",
	}

	var bodyContent bytes.Buffer
	err = tmpl.Execute(&bodyContent, data)
	if err != nil {
		fmt.Println(err)
		return
	}

	content := fmt.Sprintf("To: agustfricke@gmail.com\r\n"+
		"Subject: Email con plantilla HTML\r\n"+
		"Content-Type: text/html; charset=utf-8\r\n"+
		"\r\n"+
		"%s", bodyContent.String()) 

	err = smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"agustfricke@gmail.com",
		[]string{"agustfricke@gmail.com"},
		[]byte(content),
	)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	SendEmail()
}
