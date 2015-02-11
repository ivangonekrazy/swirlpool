package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"time"
)

type Message struct {
	event string
	data  string
	id    string
}

func (m *Message) Bytes() []byte {
	var buf bytes.Buffer

	if m.event != "" {
		buf.Write([]byte("event: "))
		buf.Write([]byte(m.event))
	}

	buf.Write([]byte("data: "))
	buf.Write([]byte(m.data))

	if m.id != "" {
		buf.Write([]byte("id: "))
		buf.Write([]byte(m.id))
	}

	buf.Write([]byte("\n\n"))

	return buf.Bytes()
}

type Hub struct {
	connections map[*Connection]bool // Registered connections.
	broadcast   chan Message         // Inbound messages from the connections.
	register    chan *Connection     // Register requests from the connections.
	unregister  chan *Connection     // Unregister requests from connections.
}

var h = Hub{
	broadcast:   make(chan Message),
	register:    make(chan *Connection),
	unregister:  make(chan *Connection),
	connections: make(map[*Connection]bool),
}

func (h *Hub) Run() {
	for {
		select {
		case c := <-h.register:
			fmt.Printf("Register connection: %s (%s)\n", c, len(h.connections))
			h.connections[c] = true
		case c := <-h.unregister:
			fmt.Printf("Unregister connection: %s (%s)\n", c, len(h.connections))
			if _, ok := h.connections[c]; ok {
				delete(h.connections, c)
				close(c.messageChan)
			}
		case m := <-h.broadcast:
			for c := range h.connections {
				select {
				case c.messageChan <- m:
				default:
					delete(h.connections, c)
					close(c.messageChan)
				}
			}
		}
	}
}

type Connection struct {
	messageChan chan Message
}

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

func ClientHandler(w http.ResponseWriter, r *http.Request) {
	var b bytes.Buffer
	flusher := w.(http.Flusher)

	conn := &Connection{messageChan: make(chan Message)}
	h.register <- conn
	defer func() { h.unregister <- conn }()

	closeNotifier := w.(http.CloseNotifier).CloseNotify()
	go func() {
		<-closeNotifier
		h.unregister <- conn
	}()

	// declare SSE MIME type
	w.Header().Set("Content-type", "text/event-stream")

	for {
		m := <-conn.messageChan
		b.Write(m.Bytes())
		b.WriteTo(w)
		flusher.Flush()
	}

}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	messageText := r.PostFormValue("message")
	h.broadcast <- Message{data: messageText}

	fmt.Printf("Received message: %s\n", messageText)

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	return
}

func main() {
	go h.Run()
	go broadcaster()

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(".")))) // for index.html
	http.HandleFunc("/sse", ClientHandler)
	http.HandleFunc("/send", PostHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
