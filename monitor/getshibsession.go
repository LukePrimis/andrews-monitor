package monitor

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
)

func (a *AndrewsTask) getShibSession() (string, error) {
	var data = strings.NewReader(`RelayState=https%3A%2F%2Fget.cbord.com%2Fbrown%2Ffull%2Flogin.php&SAMLResponse=` + url.QueryEscape(a.Saml))
	req, err := http.NewRequest("POST", "https://get.cbord.com/brown/Shibboleth.sso/SAML2/POST", data)
	if err != nil {
		return "", nil
	}
	req.Header.Set("authority", "get.cbord.com")
	req.Header.Set("cache-control", "max-age=0")
	req.Header.Set("sec-ch-ua", `"Google Chrome";v="87", " Not;A Brand";v="99", "Chromium";v="87"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("origin", "https://sso.brown.edu")
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36")
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("sec-fetch-site", "cross-site")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("referer", "https://sso.brown.edu/")
	req.Header.Set("accept-language", "en-US,en;q=0.9")
	resp, err := a.Client.Do(req)
	if err != nil {
		return "", nil
	}

	if resp.StatusCode != 302 {
		return "", errors.New("non 302 response")
	}

	return "got shibsess", nil
}
