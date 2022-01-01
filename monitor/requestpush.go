package monitor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func (a *AndrewsTask) requestPush() (string, error) {
	var data = strings.NewReader(`sid=` + url.PathEscape(a.Sid) + `&device=phone1&factor=Duo+Push&dampen_choice=true&out_of_date=False&days_out_of_date=0&days_to_block=None`)
	req, err := http.NewRequest("POST", "https://api-bf620061.duosecurity.com/frame/prompt", data)
	if err != nil {
		return "", nil
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("sec-ch-ua", `"Google Chrome";v="87", " Not;A Brand";v="99", "Chromium";v="87"`)
	req.Header.Set("Accept", "text/plain, */*; q=0.01")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Origin", "https://api-bf620061.duosecurity.com")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", a.LoginSession)
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	resp, err := a.Client.Do(req)
	if err != nil {
		return "", nil
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}

	var pushResp map[string]interface{}

	err = json.Unmarshal(bodyText, &pushResp)

	if err != nil {
		fmt.Println(err)
		return "", nil
	}

	txmap, ok := pushResp["response"].(map[string]interface{})

	if !ok {
		panic("failed")
	}

	a.Txid = txmap["txid"].(string)

	return "pushed login", nil
}
