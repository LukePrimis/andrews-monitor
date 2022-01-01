package monitor

import "strings"

func parse(body, start, end string) string {
	startInd := strings.Index(body, start) + len(start)
	if startInd == -1 {
		return ""
	}
	sub := body[startInd:]
	endInd := strings.Index(sub, end)
	if endInd == -1 {
		return ""
	}
	return sub[:endInd]
}
