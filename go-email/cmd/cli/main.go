package main

import (
	"fmt"
	"log"
	"time"

	"github.com/joefazee/mailer"
)

func main() {

	
	
	config := mailer.MailerConfig{
		Host:    "smtp.mailtrap.io",
		Port:     25,
		Username: "81233be52e47ee",
		Password: "59c8c70f2237d8",
		Timeout:  5 * time.Second,
		Sender:  "erdembsts@gmail.com",
	}

	sender := mailer.New(config)

	err := sender.Send("erdembsts@gmail.com", "welcome.html", nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Email sent!")
}