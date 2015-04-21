package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	portPtr := flag.Int("port", 9393, "Port to run on.")
	debugPtr := flag.Bool("debug", false, "Whether or not to print debug messages.")
	flag.Parse()

	http.HandleFunc("/ws", serveWs)

	// TODO: Restrict static file serving to dev mode.
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	// Run the game simulation
	go game.run(*debugPtr)

	fmt.Println("Starting web server on http://localhost:9393/")
	err := http.ListenAndServe(fmt.Sprintf(":%d", *portPtr), nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
