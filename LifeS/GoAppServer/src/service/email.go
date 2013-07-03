package service

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/mail"
	"net/smtp"
	"strings"
	"common"
)

func encodeRFC2047(String string) string {
	// use mail's rfc2047 to encode any string
	addr := mail.Address{String, ""}
	return strings.Trim(addr.String(), " <>")
}

func listToString(addresses []mail.Address) string {
	var list []string

	for _, a := range addresses {
		list = append(list, a.String())
	}

	return strings.Join(list, ",")
}

func sendOrderMail(subject string, content string) {
	smtpServer := "smtp.163.com"
	auth := smtp.PlainAuth(
		"",
		"asp_orders@163.com",
		"wqy123",
		smtpServer,
	)

	from := mail.Address{"", "asp_orders@163.com"}
	
	to := make([]mail.Address, len(common.MailTo))
	cc := make([]mail.Address, len(common.MailCc))
	for i, addr := range common.MailTo{
	   to[i] = mail.Address{"", addr}
	}
	for i, addr := range common.MailCc{
	   cc[i] = mail.Address{"", addr}
	}

	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = listToString(to)
	header["Cc"] = listToString(cc)
	header["Subject"] = encodeRFC2047(subject)
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(content))

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		smtpServer+":25",
		auth,
		from.Address,
		common.MailTo,
		[]byte(message),
	)
	if err != nil {
		log.Printf("Failed to send to for reason:%v", err)
	}
	err = smtp.SendMail(
		smtpServer+":25",
		auth,
		from.Address,
		common.MailCc,
		[]byte(message),
	)
	if err != nil {
		log.Printf("Failed to send cc for reason:%v", err)
	}
}
