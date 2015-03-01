package main

import (
	"fmt"
	"log"
	"net/http"
)

// main message fan-out broker
var h = Hub{
	broadcast:   make(chan Message, 5),
	register:    make(chan *Connection),
	unregister:  make(chan *Connection),
	connections: make(map[*Connection]bool),
}

var p = EventsProcessor{
	messages: NewMessageBuffer(100),
	intake:   make(chan Message),
}

func main() {
	go h.Run()
	go p.Run(h.broadcast)
	//go broadcaster(h.broadcast)

	http.HandleFunc("/sse", ClientHandler)
	http.HandleFunc("/send", PostHandler)
	http.HandleFunc("/github", GithubWebhookHandler)
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./js/")))) // for index.html
	http.Handle("/test", http.StripPrefix("/test", http.FileServer(http.Dir("."))))   // for index.html

	port := ":8080"
	fmt.Printf("Starting swirlpool on port %v...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
