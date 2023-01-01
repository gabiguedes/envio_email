package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

func sendMail(from string, to []string, body string) {
	auth := smtp.PlainAuth("Thanos", from, PASS, HOST)
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         HOST,
	}

	toHeader := strings.Join(to, ",")
	subject := "From: " + from + "\n" + "To: " + toHeader + "\n" + "Subject: Você ganhou!\n\n" + body

	//conecta com o servidor SMTP
	conn, err := tls.Dial("tcp", SERVERNAME, tlsConfig)
	checkErr(err)

	c, err := smtp.NewClient(conn, HOST)
	checkErr(err)

	err = c.Auth(auth)
	checkErr(err)

	err = c.Mail(from)
	checkErr(err)

	for _, addr := range to {
		err = c.Rcpt(addr)
		checkErr(err)
	}

	w, err := c.Data()
	checkErr(err)

	_, err = w.Write([]byte(subject))
	checkErr(err)

	err = w.Close()
	checkErr(err)

	c.Quit()
}

func checkErr(err error) {
	if err != nil {
		log.Panic("-- ERROR -- : " + err.Error())
	}
}

func main() {
	println("============= INICIANDO ENVIO DE EMAILS =============")

	sendMail("zooiv3rde@gmail.com", []string{"gabriela_rayssa@hotmail.com"}, "Voce ganhou 1 milhao de reais, clique aqui para receber!")
	fmt.Println("Email successfully sent!")
}
