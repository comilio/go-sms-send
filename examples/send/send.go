package main

import (
	"fmt"

	sms "../.."
)

func main() {
	var SMS sms.SMS
	SMS.Auth("username", "password")
	phones := []string{"+393401234567"}
	responseMessage, success := SMS.Send(phones, "Hello", "Smart")
	fmt.Println(responseMessage, success)
}
