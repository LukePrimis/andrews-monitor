package monitor

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/LukePrimis/andrews-monitor/config"
	"github.com/LukePrimis/andrews-monitor/sms"
	"github.com/LukePrimis/andrews-monitor/webhook"
)

func (a *AndrewsTask) getTimes() (string, error) {
	var data = strings.NewReader(`action=get_merchant_schedule&formToken=` + a.FormID)
	req, err := http.NewRequest("POST", "https://get.cbord.com/brown/full/process_order.php", data)
	if err != nil {
		return "", err
	}
	req.Header.Set("Host", "get.cbord.com")
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="90", "Google Chrome";v="90"`)
	req.Header.Set("accept", "*/*")
	req.Header.Set("x-requested-with", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36")
	req.Header.Set("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("origin", "https://get.cbord.com")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("referer", "https://get.cbord.com/brown/full/food_merchant.php?ID=289d47bf-12e6-4d97-9711-24ace6ae67c1")
	req.Header.Set("accept-language", "en-US,en;q=0.9")
	resp, err := a.Client.Do(req)
	if err != nil {
		return "", err
	}
	resp.Body.Close()

	date := strings.Split(time.Now().Format("01-02-2006"), "-")
	data = strings.NewReader(`action=get_order_time_options&dueDate=` + date[0] + `%2F` + date[1] + `%2F2021&formToken=` + a.FormID)
	req, err = http.NewRequest("POST", "https://get.cbord.com/brown/full/process_order.php", data)
	if err != nil {
		return "", err
	}
	req.Header.Set("Host", "get.cbord.com")
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="90", "Google Chrome";v="90"`)
	req.Header.Set("accept", "*/*")
	req.Header.Set("x-requested-with", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36")
	req.Header.Set("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("origin", "https://get.cbord.com")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("referer", "https://get.cbord.com/brown/full/food_merchant.php?ID=289d47bf-12e6-4d97-9711-24ace6ae67c1")
	req.Header.Set("accept-language", "en-US,en;q=0.9")
	resp, err = a.Client.Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		fmt.Println("non 200 response while monitoring:")
		fmt.Println(resp.Status)
		fmt.Println("sleeping and retrying")
		return "got times", nil
	}

	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	timeArray := strings.Split(string(bodyText), "</option><option")
	if len(timeArray) > 1 {
		restock := false
		instock := false
		if len(a.Times) == 0 {
			restock = true
		} else {
			instock = true
		}
		timeArray = timeArray[1:]
		formattedTimes := []string{}
		for ind, item := range timeArray {
			var fmtime string
			if ind == len(timeArray)-1 {
				fmtime = item[strings.Index(item, "> ")+2 : strings.Index(item, "</option>")]
			} else {
				fmtime = item[strings.Index(item, "> ")+2:]
			}
			a.Times[fmtime] = true
			formattedTimes = append(formattedTimes, fmtime)
		}
		if restock {
			webhook.SendWehook(formattedTimes)
			msg := fmt.Sprintf("Andrews Times Available: %s", strings.Join(formattedTimes, ", "))
			sms.SendSMS(config.PHONE_NUMBER, msg)

			fmt.Println("times found and sent to webhook and SMS")
		} else {
			fmt.Println("times still available")
		}
		keysToDelete := []string{}
		for key := range a.Times {
			if !contains(key, formattedTimes) {
				fmt.Printf("time %s out of stock\n", key)
				webhook.SendOOSWebhook(key)
				keysToDelete = append(keysToDelete, key)
			}
		}
		for _, key := range keysToDelete {
			delete(a.Times, key)
		}
		if len(a.Times) == 0 && instock {
			webhook.SendAllOOSWebhook()
		}
	} else {
		fmt.Println("no times available, sleeping")
	}

	return "got times", nil
}

func contains(target string, list []string) bool {
	for _, item := range list {
		if item == target {
			return true
		}
	}
	return false
}
