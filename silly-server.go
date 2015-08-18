package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func prefixOutput(output string, req *http.Request) string {
	url := string(req.URL.Path)
	return url + "\n" + output
}

func timeHandler(responseWriter http.ResponseWriter, req *http.Request) {
	timeString := time.Now().Format(time.Stamp)
	io.WriteString(responseWriter, prefixOutput(timeString+"\n", req))
}

func helloHandler(responseWriter http.ResponseWriter, req *http.Request) {
	io.WriteString(responseWriter, prefixOutput("Hello.\n", req))
}

func main() {
	const PORT = ":8080"

	log.Print("listening at http://localhost" + PORT)

	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/hello", helloHandler)

	err := http.ListenAndServe(PORT, nil)

	if err != nil {
		log.Fatal("unable to start server: ", err)
	}
}
