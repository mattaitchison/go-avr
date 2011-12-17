package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"go-avr.googlecode.com/git/avr"
)

// Flags
var (
	port = flag.Int("port", 8081, "port for webserver")
	addr = flag.String("addr", "10.0.0.177:23", "ip:port to AVR")
)

var amp *avr.Amp

func main() {
	flag.Parse()
	if *addr == "" {
		log.Fatalf("--addr required")
	}
	log.Printf("Connecting to %s ...", *addr)
	amp = avr.New(*addr)
	if err := amp.Ping(); err != nil {
		log.Fatalf("error connecting to amp at %s: %v", *addr, err)
	}
	log.Printf("Connected to AVR.")

	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
		log.Fatalf("http: %v", err)
	}
}
