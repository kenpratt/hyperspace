package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var simulateLatency *time.Duration = nil

func main() {
	portPtr := flag.Int("port", 9393, "Port to run on.")
	debugPtr := flag.Bool("debug", false, "Whether or not to print debug messages.")
	flag.Parse()

	http.HandleFunc("/ws", serveWs)

	// TODO: Restrict static file serving to dev mode.
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	// Enable latency simulation if environment variable is set (value is round-trip latency in milliseconds)
	// For example: $ LATENCY=200 ./watch
	if s := os.Getenv("LATENCY"); len(s) > 0 {
		if v, err := strconv.Atoi(s); err == nil {
			d := time.Duration(v/2) * time.Millisecond
			simulateLatency = &d
			log.Printf("Enabled latency: %v", d*2)
		}
	}

	// Run the game simulation
	go game.run(*debugPtr)

	fmt.Println("Starting web server on http://localhost:9393/")
	err := http.ListenAndServe(fmt.Sprintf(":%d", *portPtr), nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
