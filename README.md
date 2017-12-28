# Comilio GoLang SMS Send

GoLang client module to send SMS messages using Comilio SMS Gateway.

To use this library, you must have a valid account on https://www.comilio.it.

**Please note** SMS messages sent with this library will be deducted by your Comilio account credits.

For any questions, please contact us at tech@comilio.it

# Installation

```bash
$ go get github.com/comilio/go-sms-send
```

# Send a message with Classic type

```go
package main

import (
	"fmt"

	sms "github.com/comilio/go-sms-send"
)

func main() {
	var SMS sms.SMS
	SMS.Auth("username", "password")
	phones := []string{"+393401234567", "+393498765432"}
	responseMessage, success := SMS.Send(phones, "Hello World!", "Classic") // could be Smart and SmartPro. Return responseMessage String and success Bool
	fmt.Println(responseMessage, success)
}

```

# Check status of message

```go
package main

import (
	"fmt"

	sms "github.com/comilio/go-sms-send"
)

func main() {
	var SMS sms.SMS
	SMS.Auth("username", "password")
	responseMessage, success, array := SMS.Check("5DB89598EDC64F11A5FCF11B3FEC063E") // returns responseMessage String, success Bool, array Interface
	fmt.Println(responseMessage, success, array)
}
```

# More info

You can check out our website https://www.comilio.it
