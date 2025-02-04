package main

import (
	"fmt"
	"log"
	"net/smtp"
)

const (
	smtpServer = "smtp.gmail.com"
	smtpPort   = "587"
	from       = "ilai81395@gmail.com" //от кого
	password   = "psanrwrprzeupwmz"
	to         = "ilai81395@gmail.com" //кому
)

func sendEmailCallback(name, email, message string) error {
	// Формируем тело письма
	body := fmt.Sprintf(
		"Имя: %s\nEmail: %s\nСообщение:\n%s",
		name, email, message,
	)

	// Формируем заголовок письма
	header := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Новый отзыв\n" +
		"Content-Type: text/plain; charset=utf-8\n\n" +
		body

	// Аутентификация
	auth := smtp.PlainAuth("", from, password, smtpServer)

	// Отправляем письмо
	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, from, []string{to}, []byte(header))
	if err != nil {
		log.Printf("Ошибка отправки письма: %v", err)
		return err
	}

	log.Println("Письмо успешно отправлено!")
	return nil
}

func sendEmailAppointment(name, email, service, date, time, message string) error {

	// Формирование тела письма
	body := fmt.Sprintf(
		"Имя: %s\nEmail: %s\nУслуга: %s\nДата: %s\nВремя: %s\nСообщение:\n%s",
		name, email, service, date, time, message)

	// Заголовки письма
	header := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Новая запись\n" +
		"Content-Type: text/plain; charset=utf-8\n\n" +
		body

	// Аутентификация
	auth := smtp.PlainAuth("", from, password, smtpServer)

	// Отправка письма
	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, from, []string{to}, []byte(header))
	if err != nil {
		log.Fatalf("Ошибка отправки письма: %v", err)
	}
	log.Println("Письмо успешно отправлено!")
	return nil
}
