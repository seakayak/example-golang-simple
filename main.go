package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type app struct {
	m         sync.Mutex
	uniqueIPs map[string]struct{}
}

func (a *app) handleIndex(w http.ResponseWriter, r *http.Request) {
	log.Printf("serving %s request for %s", r.Method, r.URL.Path)

	// add the IP to the list
	remoteIP := r.Header.Get("X-Forwarded-For")
	if remoteIP == "" {
		http.Error(w, "missing X-Forwarded-For header", http.StatusBadRequest)
		return
	}
	a.uniqueIPs[remoteIP] = struct{}{}

	// save the IPs to disk
	err := saveIPs(a.uniqueIPs, "/data/ip.json")
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write to output
	fmt.Fprintln(w, "Hello from your Golang seakayak app!")
	fmt.Fprintf(w, "So far I have seen %d unique IPs\n", len(a.uniqueIPs))
}

func main() {
	// load the IPs
	ips, err := loadIPs("/data/ip.json")
	if err != nil {
		log.Fatal(err)
	}

	app := app{
		uniqueIPs: ips,
	}

	http.HandleFunc("/", app.handleIndex)
	http.ListenAndServe(":4000", nil)
}
