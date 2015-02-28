package main

import (
	"fmt"
	"testing"
)

func TestMessageBuffer(t *testing.T) {
	CAPACITY := 5
	mb := NewMessageBuffer(CAPACITY)
	for i := 1; i <= CAPACITY*2; i++ {
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

func TestMessageStreak(t *testing.T) {
	CAPACITY := 10
	mb := NewMessageBuffer(CAPACITY)
	for i := 1; i <= CAPACITY-3; i++ {
		mb.PushBack(NewMessage(fmt.Sprintf("%d", i)))
	}

	s := LatestStreak(&mb)
	if s.length != 1 {
		t.Error("Didn't detect correct streak length %v", mb.List())
	}

	for j := 0; j < 3; j++ {
		mb.PushBack(NewMessage("streak"))
	}

	s = LatestStreak(&mb)
	if s.length != 3 {
		t.Error("Didn't detect correct streak length %v", mb.List())
	}
}
