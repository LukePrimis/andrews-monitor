package monitor

import (
	"errors"
	"fmt"
	"net/http"
)

func (a *AndrewsTask) getDuo() (string, error) {
	a.LoginSession = `https://api-bf620061.duosecurity.com/frame/web/v1/auth?tx=` + a.Tx + `&parent=https%3A%2F%2Fsso.brown.edu%2Fidp%2Fprofile%2FSAML2%2FRedirect%2FSSO%3Fexecution%3De1s2&v=2.6`
	req, err := http.NewRequest("GET", "https://api-bf620061.duosecurity.com/frame/web/v1/auth?tx="+a.Tx+"&parent=https%3A%2F%2Fsso.brown.edu%2Fidp%2Fprofile%2FSAML2%2FRedirect%2FSSO%3Fexecution%3De1s2&v=2.6", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("sec-ch-ua", `"Google Chrome";v="87", " Not;A Brand";v="99", "Chromium";v="87"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Dest", "iframe")
	req.Header.Set("Referer", "https://sso.brown.edu/")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	resp, err := a.Client.Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", errors.New(fmt.Sprintf("%s error", resp.Status))
	}

	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(resp.Header)
	// fmt.Println(string(body))

	return "got duo", nil
}
