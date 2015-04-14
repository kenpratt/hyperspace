package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ws", serveWs)

	// TODO restrict static file serving to dev mode
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	go h.run()

	fmt.Println("Starting web server on http://localhost:9393/")
	err := http.ListenAndServe(":9393", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
