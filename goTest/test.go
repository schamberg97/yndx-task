package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getPort() string {
	const defaultPort string = "9000"
	var port string = defaultPort

	args := os.Args[1:]
	if len(args) > 0 {
		portArg := args[0]

		if portArg != "" {
			port = portArg
		}
	} else {
		log.Println("Using default port", defaultPort)
	}

	return port
}

func main() {
	resp, err := http.Get("http://127.0.0.1:" + getPort())
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Time received from microservice:", string(body))
	return
}
