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

	// we need to unregister our Connection at either the
	// close of the HTTP connection or at the end of this
	// function
	closeNotifier := w.(http.CloseNotifier).CloseNotify()
	go func() {
		<-closeNotifier
		h.unregister <- conn
	}()
	defer func() { h.unregister <- conn }()

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
	h.broadcast <- Message{data: []string{messageText}}

	fmt.Printf("Received message: %s\n", messageText)

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	return
}

func GithubWebhookHandler(w http.ResponseWriter, r *http.Request) {
	/* TODO Handle webhook delivery from Github
	 *
	 * - X-Github-Event header will let us know what type of message we are
	 *   dealing with (e.g. 'pull_request')
	 * - The JSON payload will describe the pull request details
	 */

	var event = r.Header.Get("X-Github-Event")
	var payload = r.Body

	return
}
