package monitor

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// these were functsions for andrewsbot but this project is just a monitor

func (a *AndrewsTask) checkFunds() (string, error) {
	req, err := http.NewRequest("GET", "https://get.cbord.com/brown/full/funds_home.php", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("authority", "get.cbord.com")
	req.Header.Set("cache-control", "max-age=0")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36")
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("sec-fetch-site", "cross-site")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-ch-ua", `"Google Chrome";v="87", " Not;A Brand";v="99", "Chromium";v="87"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("referer", "https://sso.brown.edu/")
	req.Header.Set("accept-language", "en-US,en;q=0.9")
	resp, err := a.Client.Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", errors.New("non 200 resp")
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	bString := string(body)

	user := bString[strings.Index(bString, `getOverview("`)+13:]
	endInd := strings.Index(user, `")`)
	user = user[:endInd]

	formToken := bString[strings.Index(bString, `"formToken": "`)+14:]
	endInd = strings.Index(formToken, `"}`)
	formToken = formToken[:endInd]

	fmt.Println(user, formToken)

	var data = strings.NewReader(fmt.Sprintf(`userId=%s&formToken=%s`, user, formToken))
	req, err = http.NewRequest("POST", "https://get.cbord.com/brown/full/funds_overview_partial.php", data)
	if err != nil {
		return "", err
	}
	req.Header.Set("authority", "get.cbord.com")
	req.Header.Set("sec-ch-ua", `"Google Chrome";v="87", " Not;A Brand";v="99", "Chromium";v="87"`)
	req.Header.Set("accept", "*/*")
	req.Header.Set("x-requested-with", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36")
	req.Header.Set("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("origin", "https://get.cbord.com")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("referer", "https://get.cbord.com/brown/full/funds_home.php")
	req.Header.Set("accept-language", "en-US,en;q=0.9")
	resp, err = a.Client.Do(req)
	if err != nil {
		return "", err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)

	bString = string(bodyText)

	swipes := bString[strings.Index(bString, `Meal Credits</td><td class="last-child balance">`)+48:]
	swipes = swipes[:strings.Index(swipes, `</td`)]

	fmt.Printf("%s swipes remaining\n", swipes)

	return "checked", nil
}
