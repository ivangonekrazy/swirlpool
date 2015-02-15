package main

import "testing"

func TestMessage(t *testing.T) {

	m := Message{data: []string{"Hello", "world"}}
	b := m.String()

	if b != "data: Hello\ndata: world\n\n\n" {
		t.Error("Didn't get what I expected")
	}

}
