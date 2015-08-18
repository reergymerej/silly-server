package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func timeHandler(responseWriter http.ResponseWriter, req *http.Request) {
	timeString := time.Now().Format(time.Stamp)
	io.WriteString(responseWriter, timeString+"\n")
}

func helloHandler(responseWriter http.ResponseWriter, req *http.Request) {
	io.WriteString(responseWriter, "Hello.\n")
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
