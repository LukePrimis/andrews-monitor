package monitor

import (
	"errors"
	"fmt"
	"net/http"
)

func (a *AndrewsTask) initialize() (string, error) {
	req, err := http.NewRequest("GET", "https://get.cbord.com/brown/full/prelogin.php", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("authority", "get.cbord.com")
	req.Header.Set("sec-ch-ua", `Google Chrome";v="87", " Not;A Brand";v="99", "Chromium";v="87"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36")
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("sec-fetch-site", "none")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("accept-language", "en-US,en;q=0.9")
	resp, err := a.Client.Do(req)
	if err != nil {
		return "", err
	}
	if resp.StatusCode == 200 {
		return "initialized", nil
	} else {
		return fmt.Sprintf("%s error", resp.Status), errors.New(fmt.Sprintf("%s error", resp.Status))
	}
}
