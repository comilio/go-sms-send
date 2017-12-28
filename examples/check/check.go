package main

import (
	"fmt"

	sms "../.."
)

func main() {
	var SMS sms.SMS
	SMS.Auth("username", "password")
	responseMessage, success, array := SMS.Check("5DB89598EDC64F11A5FCF11B3FEC063E")
	fmt.Println(responseMessage, success, array)
}
