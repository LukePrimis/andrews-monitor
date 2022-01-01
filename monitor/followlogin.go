package monitor

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func (a *AndrewsTask) followLogin() (string, error) {
	req, err := http.NewRequest("GET", "https://sso.brown.edu/idp/profile/SAML2/Redirect/SSO?execution=e1s2", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("sec-ch-ua", `"Google Chrome";v="87", " Not;A Brand";v="99", "Chromium";v="87"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("Referer", a.LoginSession)
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	resp, err := a.Client.Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", errors.New(fmt.Sprintf("%s error", resp.Status))
	}

	body, _ := ioutil.ReadAll(resp.Body)

	bString := string(body)

	txInd := strings.Index(bString, `data-sig-request="`) + 18

	bString = bString[txInd:]

	endInd := strings.Index(bString, `:`)

	txString := bString[:endInd]

	a.Tx = txString

	fmt.Println("got tx", a.Tx)

	appString := bString[endInd+1 : strings.Index(bString, `"`)]

	a.App = appString

	fmt.Println("got app", a.App)

	return "initiated pre 2fa login", nil
}
