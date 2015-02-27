package main

import (
	"fmt"
	"testing"
)

func TestMessageBuffer(t *testing.T) {
	CAPACITY := 5
	mb := NewMessageBuffer(CAPACITY)
	for i := 1; i <= 10; i++ {
		mb.PushBack(NewMessage(fmt.Sprintf("%d", i)))
	}

	if len(mb.messages) > CAPACITY {
		t.Error("Too many messages in the buffer.")
	}

	if len(mb.messages) != mb.Len() {
		t.Error("Length incorrectly reported.")
	}

	if mb.messages[0].data[0] != "6" {
		t.Error("Wrong message at buffer head.")
	}
}
