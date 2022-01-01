package monitor

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"time"

	"github.com/LukePrimis/andrews-monitor/webhook"
)

func Monitor() {
	jar, _ := cookiejar.New(nil)
	c := &http.Client{}

	c.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	c.Jar = jar
	c.Timeout = 15 * time.Second
	t := &AndrewsTask{
		Client: c,
		Times:  make(map[string]bool),
	}

	res, _ := t.initialize()

	fmt.Println(res)

	res, _ = t.getLogin()

	fmt.Println(res)

	res, _ = t.getSession()

	fmt.Println(res)

	res, _ = t.prepareLogin()

	fmt.Println(res)

	res, _ = t.login()

	fmt.Println(res)

	res, _ = t.followLogin()

	fmt.Println(res)

	res, _ = t.getDuo()

	fmt.Println(res)

	res, err := t.initiateDuo()

	if err != nil {
		panic(err)
	}

	res, err = t.getAuth()
	fmt.Println(res)
	if err != nil {
		panic(err)
	}

	res, _ = t.getPrompt()
	fmt.Println(res)

	t.requestPush()

	t.getStatus()

	t.getStatus()

	webhook.SendLoginWehook()

	t.getPushResponse()

	t.getSamlResponse()

	t.getShibSession()

	t.followRedirect()

	t.followRedirect2()

	t.checkFunds()

	t.getAndrews()

	for {
		t.getTimes()
		time.Sleep(time.Minute)
	}
}
