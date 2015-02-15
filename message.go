package main

import (
	"bytes"
)

// encapsulates a SSE data frame
type Message struct {
	event string
	data  []string
	id    string
}

func (m *Message) Bytes() []byte {
	var buf bytes.Buffer

	if m.event != "" {
		buf.Write([]byte("event: "))
		buf.Write([]byte(m.event))
	}

	for _, d := range m.data {
		buf.Write([]byte("data: "))
		buf.Write([]byte(d))
		buf.Write([]byte("\n"))
	}

	if m.id != "" {
		buf.Write([]byte("id: "))
		buf.Write([]byte(m.id))
	}

	buf.Write([]byte("\n\n"))

	return buf.Bytes()
}
