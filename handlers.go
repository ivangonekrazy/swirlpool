package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/jsonq"
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
		b.WriteString(m.String())
		b.WriteTo(w)
		flusher.Flush()
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	messageText := r.PostFormValue("message")
	h.broadcast <- NewMessage(messageText)

	fmt.Printf("PostHandler: received message: %s\n", messageText)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func GithubWebhookHandler(w http.ResponseWriter, r *http.Request) {
	// TODO Handle webhook delivery from Github
	// - X-Github-Event header will let us know what type of message we are
	//   dealing with (e.g. 'pull_request')
	// - The JSON payload will describe the pull request details

	//var ghEvent = r.Header.Get("X-Github-Event")
	var ghPayload = r.Body
	data := map[string]interface{}{}

	dec := json.NewDecoder(ghPayload)
	dec.Decode(&data) // decode into a generic map
	jq := jsonq.NewQuery(data)

	login, _ := jq.String("pull_request", "user", "login")
	m := NewMessage(login)
	m.SetEvent("pullrequest")
	h.broadcast <- m

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
