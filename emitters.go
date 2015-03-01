package main

import (
	"bytes"
	"os/exec"
	"time"
)

/*
// Periodically generate some data with the 'date' command
*/
func broadcaster(broadcastChan chan Message) {
	var buf bytes.Buffer

	for {
		cmd := exec.Command("date")
		cmd.Stdout = &buf
		cmd.Stderr = &buf
		cmd.Run()

		m := NewMessage(buf.String())
		m.SetEvent("datetime")
		broadcastChan <- m
		buf.Reset()

		time.Sleep(1 * time.Second)
	}
}
