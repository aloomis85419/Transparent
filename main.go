package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting http server...")
	cfg := Config{}
	cfg.Configure("57a8172387f671b61ff90cdd614f9a4bf82c071d")
	//http client initialization
	client := newLDAClient(cfg, http.Client{})
	client.listFilings()
	err := http.ListenAndServe("/", nil)
	if err != nil {
		log.Fatal("Failed to start server.")
	}
}

func CheckErr(err error, processMsg string, critical bool) {
	if err != nil {
		if critical {
			log.Fatalf("Fatal: Error occured while %s: %s", processMsg, err.Error())
		} else {
			log.Println(fmt.Sprintf("Error occured while %s: %s", processMsg, err.Error()))
		}
	}
}
