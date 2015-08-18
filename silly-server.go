package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func server(responseWriter http.ResponseWriter, req *http.Request) {
	timeString := time.Now().Format(time.Stamp)
	io.WriteString(responseWriter, timeString+"\n")
}

func main() {
	const PORT = ":8080"

	log.Print("listening at http://localhost" + PORT)

	http.HandleFunc("/time", server)
	err := http.ListenAndServe(PORT, nil)

	if err != nil {
		log.Fatal("unable to start server: ", err)
	}
}
