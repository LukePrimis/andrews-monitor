package monitor

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (a *AndrewsTask) initiateDuo() (string, error) {
	tx := strings.ReplaceAll(strings.ReplaceAll(a.Tx, "|", "%7C"), "=", "%3D")

	payload := `tx=` + tx + `&parent=https%3A%2F%2Fsso.brown.edu%2Fidp%2Fprofile%2FSAML2%2FRedirect%2FSSO%3Fexecution%3De1s2&java_version=&flash_version=&screen_resolution_width=1536&screen_resolution_height=864&color_depth=24&is_cef_browser=false&is_ipad_os=false&react_support=true`

	var data = strings.NewReader(payload)
	req, err := http.NewRequest("POST", "https://api-bf620061.duosecurity.com/frame/web/v1/auth?tx="+a.Tx+"&parent=https%3A%2F%2Fsso.brown.edu%2Fidp%2Fprofile%2FSAML2%2FRedirect%2FSSO%3Fexecution%3De1s2&v=2.6", data)
	if err != nil {
		return "", err
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("sec-ch-ua", `Google Chrome";v="87", " Not;A Brand";v="99", "Chromium";v="87"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Origin", "https://api-bf620061.duosecurity.com")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Dest", "iframe")
	req.Header.Set("Referer", a.LoginSession)
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	resp, err := a.Client.Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		fmt.Println("non 200 response")
		return "", errors.New(fmt.Sprintf("%s error", resp.Status))
	}

	body, _ := io.ReadAll(resp.Body)

	sid := parse(string(body), `<input type="hidden" name="sid" value="`, `">`)
	sid = strings.ReplaceAll(sid, `&#x3d;`, `=`)
	sid = strings.ReplaceAll(sid, `&#x7c;`, `|`)

	ukey := parse(string(body), `<input type="hidden" name="ukey" value="`, `">`)
	akey := parse(string(body), `<input type="hidden" name="akey" value=`, `>`)
	txid := parse(string(body), `<input type="hidden" name="txid" value=`, `>`)

	a.Sid = sid
	a.Ukey = ukey
	a.Akey = akey
	a.Txid = txid

	fmt.Println(sid)
	fmt.Println(ukey)
	fmt.Println(akey)
	fmt.Println(txid)

	return "posted to duo", nil
}
