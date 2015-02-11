package main

import (
	"bytes"
	"fmt"
	"net/http"
)

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
