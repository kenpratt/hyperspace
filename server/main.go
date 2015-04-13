package main

import (
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/websocket"
)

// Echo the data received on the WebSocket.
func EchoServer(ws *websocket.Conn) {
	io.Copy(ws, ws)
}

// This example demonstrates a trivial echo server.
func main() {
	http.Handle("/api", websocket.Handler(EchoServer))

	// TODO restrict static file serving to dev mode
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	fmt.Println("Starting web server on http://localhost:9393/")
	err := http.ListenAndServe(":9393", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
