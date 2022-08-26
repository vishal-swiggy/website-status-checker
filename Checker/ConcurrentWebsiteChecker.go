package Checker

import (
	"context"
	"status-checker/WebsiteMap"
	"time"
)

func routine(channel chan string, routineNo int, checker StatusChecker) {
	for {

		select {
		case website := <-channel:
			ctx := context.WithValue(context.Background(), "routineNo", routineNo)
			if status, _ := checker.Check(ctx, website); status {
				WebsiteMap.UpdateWebsiteStatus(website, "UP")
			} else {
				WebsiteMap.UpdateWebsiteStatus(website, "DOWN")
			}
		}
	}
}

func ConcurrentWebsiteCheck(concurrency int, delay time.Duration) {

	channel := make(chan string, concurrency)
	var checker StatusChecker = WebsiteChecker{}

	for i := 1; i <= concurrency; i++ {
		go routine(channel, i, checker)
	}

	for {
		websiteMap := WebsiteMap.GetWebsiteList()
		for _, website := range websiteMap {
			channel <- website
		}
		time.Sleep(time.Second * delay)
	}
}
