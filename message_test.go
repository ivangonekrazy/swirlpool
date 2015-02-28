package main

import "testing"

func TestMessage(t *testing.T) {
	m := NewMessage("Hello")
	m.AppendData("world")
	b := m.String()

	if b != "data: Hello\ndata: world\n\n\n" {
		t.Error("Didn't get what I expected")
	}
}

func TestOptionalEvent(t *testing.T) {
	m := NewMessage("Hello", "world")

	if m.event != "world" {
		t.Error("NewMessage contstructor didn't set event.")
	}
}

func TestSetEvent(t *testing.T) {
	m := NewMessage("one")
	m.SetEvent("hello")

	if m.String() != "event: hello\ndata: one\n\n\n" {
		t.Error("Should have an event set.")
	}
}

func TestMessageSetData(t *testing.T) {
	m := new(Message)
	m.SetData("one", "two")

	if m.String() != "data: one\ndata: two\n\n\n" {
		t.Error("Should then have one and two.")
	}
}

func TestMessageAppend(t *testing.T) {
	m := NewMessage("one")
	if m.String() != "data: one\n\n\n" {
		t.Error("Should start with one.")
	}

	m.AppendData("two")
	if m.String() != "data: one\ndata: two\n\n\n" {
		t.Error("Should then have one and two.")
	}
}

func TestMessageSame(t *testing.T) {
	m1 := NewMessage("one")
	m2 := NewMessage("one")
	if !m1.Same(m2) {
		t.Error("Message data should be the same.")
	}

	m1.AppendData("two")
	if m1.Same(m2) {
		t.Error("Message data should not be the same after one append.")
	}

	m2.AppendData("two")
	if !m1.Same(m2) {
		t.Error("Message data should be the same after another append.")
	}

	m1.SetEvent("newtopic")
	if m1.Same(m2) {
		t.Error("Message data should not be the same after one sets an event.")
	}
}
