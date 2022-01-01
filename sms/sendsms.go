package sms

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/LukePrimis/andrews-monitor/config"
)

func SendSMS(number string, message string) {
	client := &http.Client{}
	var data = strings.NewReader(fmt.Sprintf(config.TWIOLIO_REQUEST_BODY, number, message))
	req, err := http.NewRequest("POST", config.TWILIO_API_URL, data)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Authorization", config.TWILIO_API_KEY)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp.Body.Close()
	if resp.StatusCode != 201 {
		fmt.Println("error sending sms")
		return
	}
}
