package common

import (
	"fmt"
	"sync"

	"gopkg.in/gomail.v2"
)

func SendMail(email string, productName string, quantity int, wg *sync.WaitGroup) {
	message :=  fmt.Sprintf("Dear customer,\n\nWe regret to inform you that your order for the product %s with quantity %d could not be processed as we currently do not have sufficient stock. We will notify you once the product is available again, and we will process your order at that time.\n\nThank you for your understanding and patience!", productName, quantity)
	if wg != nil {
		defer wg.Done()
		message = fmt.Sprintf("Dear customer,\n\nYour order for the product %s with quantity %d has been successfully processed.\n\nThank you for your order!", productName, quantity)
	}
	mail := gomail.NewMessage()
	mail.SetHeader("From", <mail address>)
	mail.SetHeader("To", email)
	mail.SetHeader("Subject", "Order successful")
	body := message
	mail.SetBody("text/plain", body)

	config := gomail.NewDialer("smtp.gmail.com", 587, <mail address>, <password>)
	if err := config.DialAndSend(mail); err != nil {
		panic(err)
	}

	fmt.Println("Email sent successfully to", email)
}