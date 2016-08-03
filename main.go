package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"golang.org/x/net/http2"
)

func server() {

	certFile, _ := filepath.Abs("cert/server.crt")
	keyFile, _ := filepath.Abs("cert/server.key")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Protocol: %s\n", r.Proto)
	})

	err := http.ListenAndServeTLS(":3000", certFile, keyFile, nil)
	if err != nil {
		log.Printf("Error: %s", err)
	}

}

func client() {
	/*
		// If TLSClientConfig is not nil. Then http.Client does not work with HTTP/2.
		t := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	*/

	t := &http2.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: t}
	res, err := client.Get("https://localhost:3000")
	if err != nil {
		log.Print("Error:", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Print("Error:", err)
		return
	}
	log.Print("Res:", string(body))
}

func main() {
	if len(os.Args) == 1 {
		return
	}

	cmd := os.Args[1]
	switch cmd {
	case "server":
		server()
	case "client":
		client()
	}
}
