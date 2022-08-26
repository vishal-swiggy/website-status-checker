package API

import (
	"encoding/json"
	"log"
	"net/http"
	"status-checker/WebsiteMap"
	"strings"
)

type postRequest struct {
	WebsiteList []string `json:"websites"`
}

func PostWebsiteName(rw http.ResponseWriter, req *http.Request) {
	log.Printf("API call: %s %s\n", req.Method, req.RequestURI)

	var websiteList postRequest

	if err := json.NewDecoder(req.Body).Decode(&websiteList); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	for _, website := range websiteList.WebsiteList {
		WebsiteMap.AddWebsite(website)
	}
	rw.Write([]byte("Following websites are being monitored: " + strings.Join(WebsiteMap.GetWebsiteList(), ",")))
}

func GetWebsiteStatus(rw http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	log.Printf("API call: %s %s\n", req.Method, req.RequestURI)
	var out []byte
	out, _ = json.Marshal(WebsiteMap.GetWebsiteStatusMap(name))
	rw.Write(out)
}
