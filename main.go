package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting http server...")
	loadDynamicConstants()
	log.Println("Server is running, waiting for connections...")
	err := http.ListenAndServe(":8080", nil)
	CheckErr(err, "starting server", true)
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
