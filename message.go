package main

import (
	"bytes"
	"fmt"
)

// encapsulates a SSE data frame
type Message struct {
	event string
	data  []string
	id    string
}

func NewMessage(data string, event_optional ...string) Message {
	m := Message{data: []string{data}}
	if len(event_optional) == 1 {
		m.SetEvent(event_optional[0])
	}
	return m
}

func (m *Message) Reset() {
	m.event = ""
	m.data = nil
	m.id = ""
}

func (m *Message) SetEvent(event string) {
	m.event = event
}

func (m *Message) SetId(id string) {
	m.id = id
}

func (m *Message) SetData(datas ...string) {
	m.data = datas
}

func (m *Message) AppendData(data string) {
	m.data = append(m.data, data)
}

func (m *Message) Len() int {
	return len(m.data)
}

func (m *Message) Same(o Message) bool {
	same := true

	if m.event != o.event {
		same = false
	}

	if m.Len() != o.Len() {
		same = false
	}

	// if we know things are not equal by now, we can
	// skip the following deep inspection
	if !same {
		return false
	}

	for i, _ := range m.data {
		if m.data[i] != o.data[i] {
			return false
		}
	}

	return same
}

func (m *Message) String() string {
	var buf bytes.Buffer

	if m.event != "" {
		buf.WriteString(fmt.Sprintf("event: %s\n", m.event))
	}

	for _, d := range m.data {
		buf.WriteString(fmt.Sprintf("data: %s\n", d))
	}

	if m.id != "" {
		buf.WriteString(fmt.Sprintf("id: %s\n", m.id))
	}

	buf.Write([]byte("\n\n"))

	return buf.String()
}
