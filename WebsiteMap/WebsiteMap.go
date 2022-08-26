package WebsiteMap

import (
	"log"
	"net/http"
)

type Status string

type WebsiteStatus struct {
	status Status `json:"status"`
}

var websiteMap map[string]WebsiteStatus

func init() {
	websiteMap = newWebsiteMap()
}

// Singleton Pattern
func newWebsiteMap() map[string]WebsiteStatus {
	if websiteMap == nil {
		return make(map[string]WebsiteStatus)
	} else {
		return nil
	}
}

func AddWebsite(website string) (websiteExists bool) {
	_, websiteExists = websiteMap[website]
	var status string
	if !websiteExists {
		log.Printf("`%s` added to map", website)
		if response, err := http.Get("https://" + website); err == nil {
			if response.StatusCode == 200 {
				log.Printf("Website: `%s` UP  Status code: `%d`\n", website, response.StatusCode)
				status = "UP"
			} else {
				log.Printf("Website: `%s` DOWN  Status code: `%d`\n", website, response.StatusCode)
				status = "DOWN"
			}
		} else {
			log.Printf("Error: `%s`\n", err.Error())
			status = "DOWN"
		}
		websiteMap[website] = WebsiteStatus{
			status: Status(status),
		}
	} else {
		log.Printf("`%s` already exists", website)
	}
	return
}

func GetWebsiteList() (websites []string) {
	for website := range websiteMap {
		websites = append(websites, website)
	}
	return
}

func GetWebsiteStatusMap(website string) (websiteStatusMap map[string]string) {
	websiteStatusMap = make(map[string]string)
	if website == "" {
		for website := range websiteMap {
			websiteStatusMap[website] = string(websiteMap[website].status)
		}
	} else {
		websiteProperties, isPresent := websiteMap[website]
		if isPresent {
			websiteStatusMap[website] = string(websiteProperties.status)
		} else {
			websiteStatusMap[website] = "NOT REGISTERED"
		}
	}
	return
}

func UpdateWebsiteStatus(website string, status Status) {
	if websiteStatus, isPresent := websiteMap[website]; isPresent {
		websiteStatus.status = status
		websiteMap[website] = websiteStatus
	}
}
