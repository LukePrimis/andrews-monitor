package webhook

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/LukePrimis/andrews-monitor/config"
)

func SendWehook(availableTimes []string) {
	wh := MakeWebhook(availableTimes)
	payload, err := json.Marshal(wh)
	if err != nil {
		return
	}
	http.Post(config.DISCORD_WEBHOOK_URL, "application/json", bytes.NewBuffer(payload))
}

func SendLoginWehook() {
	wh := LoginWebhook()
	payload, err := json.Marshal(wh)
	if err != nil {
		return
	}
	http.Post(config.DISCORD_WEBHOOK_URL, "application/json", bytes.NewBuffer(payload))
}

func SendOOSWebhook(timeSlot string) {
	wh := OOSWebhook(timeSlot)
	payload, err := json.Marshal(wh)
	if err != nil {
		return
	}
	http.Post(config.DISCORD_WEBHOOK_URL, "application/json", bytes.NewBuffer(payload))
}

func SendAllOOSWebhook() {
	wh := AllOOSWebhook()
	payload, err := json.Marshal(wh)
	if err != nil {
		return
	}
	http.Post(config.DISCORD_WEBHOOK_URL, "application/json", bytes.NewBuffer(payload))
}
