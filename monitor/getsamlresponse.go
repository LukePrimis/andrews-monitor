package monitor

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func (a *AndrewsTask) getSamlResponse() (string, error) {
	payload := url.PathEscape(a.Auth + ":" + a.App)
	var data = strings.NewReader(`_eventId=proceed&sig_response=` + payload)
	req, err := http.NewRequest("POST", "https://sso.brown.edu/idp/profile/SAML2/Redirect/SSO?execution=e1s2", data)
	if err != nil {
		return "", err
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("sec-ch-ua", `"Google Chrome";v="87", " Not;A Brand";v="99", "Chromium";v="87"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Origin", "https://sso.brown.edu")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Referer", "https://sso.brown.edu/idp/profile/SAML2/Redirect/SSO?execution=e1s2")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	resp, err := a.Client.Do(req)
	if err != nil {
		return "", err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	startInd := strings.Index(string(bodyText), `<input type="hidden" name="SAMLResponse" value="`) + 48

	sub := string(bodyText)[startInd:]

	endInd := strings.Index(sub, "/>") - 1
	sub = sub[:endInd]

	// fmt.Println("got saml resp", sub)

	a.Saml = sub

	return "", nil
}
