package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"status-checker/API"
	"status-checker/Checker"
	"time"
)

func DeployService() {

	concurrency := 5
	var sleepTime time.Duration = 60

	router := mux.NewRouter()
	router.HandleFunc("/websites", API.PostWebsiteName).Methods(http.MethodPost)
	router.HandleFunc("/websites", API.GetWebsiteStatus).Methods(http.MethodGet)

	log.Printf("Starting server...")
	var err error
	go func() error {
		err = http.ListenAndServe("localhost:8080", router)
		return err
	}()
	if err == nil {
		log.Printf("Server started at localhost:8080")
	} else {
		log.Printf("%s\n", err)
	}
	go func() {
		Checker.ConcurrentWebsiteCheck(concurrency, sleepTime)
	}()
	return
}

func main() {

	DeployService()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Printf("Stopping server...")
	os.Exit(0)

	return
}
