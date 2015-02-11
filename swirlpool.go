package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"time"
)

// main message fan-out broker
var h = Hub{
	broadcast:   make(chan Message),
	register:    make(chan *Connection),
	unregister:  make(chan *Connection),
	connections: make(map[*Connection]bool),
}

// Periodically generate some data with the 'date' command
func broadcaster() {
	var buf bytes.Buffer

	for {
		cmd := exec.Command("date")
		cmd.Stdout = &buf
		cmd.Stderr = &buf
		cmd.Run()

		m := Message{data: buf.String()}
		h.broadcast <- m
		buf.Reset()

		time.Sleep(2 * time.Second)
	}
}

func main() {
	go h.Run()
	go broadcaster()

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(".")))) // for index.html
	http.HandleFunc("/sse", ClientHandler)
	http.HandleFunc("/send", PostHandler)

	fmt.Println("Starting swirlpool...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
