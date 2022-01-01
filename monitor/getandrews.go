package monitor

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func (a *AndrewsTask) getAndrews() (string, error) {
	req, err := http.NewRequest("GET", "https://get.cbord.com/brown/full/food_merchant.php?ID=289d47bf-12e6-4d97-9711-24ace6ae67c1", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("authority", "get.cbord.com")
	req.Header.Set("cache-control", "max-age=0")
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="90", "Google Chrome";v="90"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36")
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("referer", "https://get.cbord.com/brown/full/food_home.php")
	req.Header.Set("accept-language", "en-US,en;q=0.9")
	resp, err := a.Client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	formToken := parse(string(bodyText), `<input type="hidden" name="formToken" value="`, `"`)
	fmt.Println("FormToken: " + formToken)
	a.FormID = formToken
	return "got formtoken", nil
}
