package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/Wishmob/bookingsApp/internal/models"
	mail "github.com/xhit/go-simple-mail/v2"
)

func listenForMail() {
	go func() {
		for {
			msg := <-app.MailChan
			sendMail(msg)
		}
	}()
}

func sendMail(msg models.MailData) {
	server := mail.NewSMTPClient()
	server.KeepAlive = false
	server.SendTimeout = 10 * time.Second
	server.ConnectTimeout = 10 * time.Second
	server.Port = 1025
	server.Host = "localhost"

	client, err := server.Connect()
	if err != nil {
		app.ErrorLog.Println(err)
	}

	email := mail.NewMSG()
	email.SetFrom(msg.From).AddTo(msg.To).SetSubject(msg.Subject)
	if msg.Template == "" {
		email.SetBody(mail.TextHTML, msg.Content)
	} else {
		templateFile, err := ioutil.ReadFile(fmt.Sprintf("./email-templates/%s", msg.Template))
		if err != nil {
			log.Println("failed to load template file.")
		}
		template := string(templateFile)
		email.SetBody(mail.TextHTML, strings.Replace(template, "[%body%]", msg.Content, 1))
	}
	email.Send(client)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("email sent!")
	}
}
