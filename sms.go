package sms

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// SMS : Struct username, password, authToken
type SMS struct {
	username, password string
}

// Auth : Authorize for SMS service Comilio
func (s *SMS) Auth(username, password string) {
	s.username = username
	s.password = password
}

type responseErrorStruct struct {
	Error string `json:"error"`
}

type requestSendStruct struct {
	PhoneNumbers []string `json:"phone_numbers"`
	MessageType  string   `json:"message_type"`
	Text         string   `json:"text"`
}

type responseSendStruct struct {
	MessageID string `json:"message_id"`
}

type statusStruct struct {
	PhoneNumber string `json:"phone_number"`
	Status      string `json:"status"`
}

type responseCheckStruct struct {
	Collection []statusStruct
}

// Send : Send sms request
func (s SMS) Send(phones []string, text string, messageType string) (string, bool) {
	valid := map[string]bool{"Classic": true, "Smart": true, "SmartPro": true}
	if !valid[messageType] {
		return "Message type must be Classic, Smart or SmartPro", false
	}
	structForJSON := &requestSendStruct{
		PhoneNumbers: phones,
		MessageType:  messageType,
		Text:         text,
	}
	jsonFromStruct, _ := json.Marshal(structForJSON)
	jsonStr := []byte(string(jsonFromStruct))
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.comilio.it/rest/v1/message/", bytes.NewBuffer(jsonStr))
	req.SetBasicAuth(s.username, s.password)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var message string
	var success bool
	if resp.StatusCode != 200 {
		errorMessage := responseErrorStruct{}
		json.Unmarshal(body, &errorMessage)
		message = errorMessage.Error
		success = false
	} else {
		successMessage := responseSendStruct{}
		json.Unmarshal(body, &successMessage)
		message = successMessage.MessageID
		success = true
	}
	return message, success
}

// Check : Check the sms status
func (s SMS) Check(messageID string) (string, bool, interface{}) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.comilio.it/rest/v1/message/"+messageID, nil)
	req.SetBasicAuth(s.username, s.password)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var message string
	var success bool
	var object interface{}
	if resp.StatusCode != 200 {
		errorMessage := responseErrorStruct{}
		json.Unmarshal(body, &errorMessage)
		message = errorMessage.Error
		success = false
		object = nil
	} else {
		successMessage := make([]statusStruct, 0)
		json.Unmarshal(body, &successMessage)
		message = "Success"
		success = true
		object = successMessage
	}
	return message, success, object
}
