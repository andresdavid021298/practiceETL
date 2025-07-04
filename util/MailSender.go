package util

import (
	"log"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendMail(toAddress string, subject string, body string) {
	dialer := configureMailDialer()
	message := buildNewMessage(toAddress, subject, body)
	log.Println("** ENVIANDO CORREO A:", toAddress, "**")
	sendMailError := dialer.DialAndSend(message)
	if sendMailError != nil {
		log.Println("** ERROR AL ENVIAR CORREO:", sendMailError.Error(), "**")
		return
	}
	log.Println("** CORREO ENVIADO EXITOSAMENTE **")
}

func buildNewMessage(toAddress string, subject string, body string) *gomail.Message {
	addressFrom := LoadProperty("from.mail.address")
	newMail := gomail.NewMessage()
	newMail.SetHeader("From", addressFrom)
	newMail.SetHeader("To", toAddress)
	newMail.SetHeader("Subject", subject)
	newMail.SetBody("text/plain", body)
	return newMail
}

func configureMailDialer() *gomail.Dialer {
	host := LoadProperty("host.mail.protocol")
	port, _ := strconv.Atoi(LoadProperty("host.mail.port"))
	username := LoadProperty("mail.username")
	password := LoadProperty("mail.password")
	return gomail.NewDialer(host, port, username, password)
}
