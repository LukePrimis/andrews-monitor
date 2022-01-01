package monitor

import (
	"log"
	"net/http"
	"net/url"
	"strings"
)

func (a *AndrewsTask) getAuth() (string, error) {
	var data = strings.NewReader(url.PathEscape(`sid=` + a.Sid + `&akey=` + a.Akey + `&txid=` + a.Txid + `&response_timeout=15&parent=https%3A%2F%2Fsso.brown.edu%2Fidp%2Fprofile%2FSAML2%2FRedirect%2FSSO%3Fexecution%3De1s2&duo_app_url=https%3A%2F%2F127.0.0.1%2Freport&eh_service_url=https%3A%2F%2F1.endpointhealth.duosecurity.com%2Fv1%2Fhealthapp%2Fdevice%2Fhealth%3F_req_trace_group%3Dd78fb730f0ef59a291196e62%252C5b932715ed1c930226cc2a22&eh_download_link=https%3A%2F%2Fdl.duosecurity.com%2FDuoDeviceHealth-latest.msi&is_silent_collection=true`))
	req, err := http.NewRequest("POST", "https://api-bf620061.duosecurity.com/frame/web/v1/auth?tx="+a.Tx+"&parent=https%3A%2F%2Fsso.brown.edu%2Fidp%2Fprofile%2FSAML2%2FRedirect%2FSSO%3Fexecution%3De1s2&v=2.6", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="90", "Google Chrome";v="90"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Origin", "https://api-bf620061.duosecurity.com")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Dest", "iframe")
	req.Header.Set("Referer", "https://api-bf620061.duosecurity.com/frame/web/v1/auth?tx=TX|bHByaW1pczF8RElWUEI2NlRZWVo3R1lLNDFLTkx8MTYyMTE3ODM3NA==|cc221988e6f03cc21512df71c8e9cf646f48daad&parent=https%3A%2F%2Fsso.brown.edu%2Fidp%2Fprofile%2FSAML2%2FRedirect%2FSSO%3Fexecution%3De1s2&v=2.6")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	resp, err := a.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()

	return "got auth", nil
}
