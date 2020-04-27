package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// http server for echo hostname
// maybe should ref for more function： https://gist.github.com/enricofoltran/10b4a980cd07cb02836f70a4ab3e72d7

func main() {
	var listenAddr string
	flag.StringVar(&listenAddr, "port", ":8080", "http server listen address ip:port")
	flag.Parse()
	http.HandleFunc("/", mainHandler)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}

func mainHandler(w http.ResponseWriter, req *http.Request) {
	hostname, err := os.Hostname()
	t := time.Now()
	// https://stackoverflow.com/questions/20234104/how-to-format-current-time-using-a-yyyymmddhhmmss-format
	tString := t.Format("2006-01-02 15:04:05")
	// https://stackoverflow.com/questions/27234861/correct-way-of-getting-clients-ip-addresses-from-http-request
	var remoteAddr string
	remoteAddr = req.Header.Get("x-forwarded-for")
	if remoteAddr == "" {
		remoteAddr = req.RemoteAddr
	}

	log.Printf(": Request from %s\n", remoteAddr)

	if err != nil {
		fmt.Fprintf(w, "%s : Error get hostname", tString)
	} else {
		fmt.Fprintf(w, "%s : Hello from %s\n", tString, hostname)
	}
}
