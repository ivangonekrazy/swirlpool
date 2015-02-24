package main

import (
	"fmt"
	"log"
	"net/http"
)

// main message fan-out broker
var h = Hub{
	broadcast:   make(chan Message),
	register:    make(chan *Connection),
	unregister:  make(chan *Connection),
	connections: make(map[*Connection]bool),
}

func main() {
	go h.Run()
	go broadcaster()

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(".")))) // for index.html
	http.HandleFunc("/sse", ClientHandler)
	http.HandleFunc("/send", PostHandler)
	http.HandleFunc("/github", GithubWebhookHandler)

	fmt.Println("Starting swirlpool...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
